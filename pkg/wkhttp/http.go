package wkhttp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/cache"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/log"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/gocraft/dbr/v2"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// UserRole 用户角色
type UserRole string

const (
	// Admin 管理员
	Admin UserRole = "admin"
	// SuperAdmin 超级管理员
	SuperAdmin UserRole = "superAdmin"
)

// WKHttp WKHttp
type WKHttp struct {
	r    *gin.Engine
	pool sync.Pool
}

// New New
func New() *WKHttp {
	l := &WKHttp{
		r:    gin.New(),
		pool: sync.Pool{},
	}
	l.r.Use(gin.Recovery())
	l.pool.New = func() interface{} {
		return allocateContext()
	}
	return l
}

func allocateContext() *Context {
	return &Context{Context: nil, lg: log.NewTLog("context")}
}

// Context Context
type Context struct {
	*gin.Context
	lg log.Log
}

func (c *Context) reset() {
	c.Context = nil
}

// ResponseError ResponseError
func (c *Context) ResponseError(err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":    err.Error(),
		"status": http.StatusBadRequest,
	})
}

// ResponseErrorf ResponseErrorf
func (c *Context) ResponseErrorf(msg string, err error) {
	if err != nil {
		c.lg.Error(msg, zap.Error(err), zap.String("path", c.FullPath()))
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":    msg,
		"status": http.StatusBadRequest,
	})
}

// ResponseErrorWithStatus ResponseErrorWithStatus
func (c *Context) ResponseErrorWithStatus(err error, status int) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":    err.Error(),
		"status": status,
	})
}

// GetPage 获取页参数
func (c *Context) GetPage() (pageIndex int64, pageSize int64) {
	pageIndex, _ = strconv.ParseInt(c.Query("page_index"), 10, 64)
	pageSize, _ = strconv.ParseInt(c.Query("page_size"), 10, 64)
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 15
	}
	return
}

// ResponseOK 返回成功
func (c *Context) ResponseOK() {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

// Response Response
func (c *Context) Response(data interface{}) {
	c.JSON(http.StatusOK, data)
}

// ResponseWithStatus ResponseWithStatus
func (c *Context) ResponseWithStatus(status int, data interface{}) {
	c.JSON(status, data)
}

// GetLoginUID 获取当前登录的用户uid
func (c *Context) GetLoginUID() string {
	return c.MustGet("uid").(string)
}

// GetAppID appID
func (c *Context) GetAppID() string {
	return c.GetHeader("appid")
}

// GetLoginName 获取当前登录的用户名字
func (c *Context) GetLoginName() string {
	return c.MustGet("name").(string)
}

// GetLoginRole 获取当前登录用户的角色
func (c *Context) GetLoginRole() string {
	return c.GetString("role")
}

// GetSpanContext 获取当前请求的span context
func (c *Context) GetSpanContext() opentracing.SpanContext {
	return c.MustGet("spanContext").(opentracing.SpanContext)
}

// CheckLoginRole 检查登录角色权限
func (c *Context) CheckLoginRole() error {
	role := c.GetLoginRole()
	if role == "" {
		return errors.New("登录用户角色错误")
	}
	if role != string(Admin) && role != string(SuperAdmin) {
		return errors.New("该用户无权执行此操作")
	}
	return nil
}

// CheckLoginRoleIsSuperAdmin 检查登录用户为超级管理员
func (c *Context) CheckLoginRoleIsSuperAdmin() error {
	role := c.GetLoginRole()
	if role == "" {
		return errors.New("登录用户角色错误")
	}
	if role != string(SuperAdmin) {
		return errors.New("该用户无权执行此操作")
	}
	return nil
}

// HandlerFunc HandlerFunc
type HandlerFunc func(c *Context)

// WKHttpHandler WKHttpHandler
func (l *WKHttp) WKHttpHandler(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		hc := l.pool.Get().(*Context)
		hc.reset()
		hc.Context = c
		defer l.pool.Put(hc)

		handlerFunc(hc)

		//handlerFunc(&Context{Context: c})
	}
}

// Run Run
func (l *WKHttp) Run(addr ...string) error {
	return l.r.Run(addr...)
}

func (l *WKHttp) RunTLS(addr, certFile, keyFile string) error {
	return l.r.RunTLS(addr, certFile, keyFile)
}

// POST POST
func (l *WKHttp) POST(relativePath string, handlers ...HandlerFunc) {
	l.r.POST(relativePath, l.handlersToGinHandleFuncs(handlers)...)
}

// GET GET
func (l *WKHttp) GET(relativePath string, handlers ...HandlerFunc) {
	l.r.GET(relativePath, l.handlersToGinHandleFuncs(handlers)...)
}

// Any Any
func (l *WKHttp) Any(relativePath string, handlers ...HandlerFunc) {
	l.r.Any(relativePath, l.handlersToGinHandleFuncs(handlers)...)
}

// Static Static
func (l *WKHttp) Static(relativePath string, root string) {
	l.r.Static(relativePath, root)
}

// LoadHTMLGlob LoadHTMLGlob
func (l *WKHttp) LoadHTMLGlob(pattern string) {
	l.r.LoadHTMLGlob(pattern)
}

// UseGin UseGin
func (l *WKHttp) UseGin(handlers ...gin.HandlerFunc) {
	l.r.Use(handlers...)
}

// Use Use
func (l *WKHttp) Use(handlers ...HandlerFunc) {
	l.r.Use(l.handlersToGinHandleFuncs(handlers)...)
}

// ServeHTTP ServeHTTP
func (l *WKHttp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	l.r.ServeHTTP(w, req)
}

// Group Group
func (l *WKHttp) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return newRouterGroup(l.r.Group(relativePath, l.handlersToGinHandleFuncs(handlers)...), l)
}

// HandleContext HandleContext
func (l *WKHttp) HandleContext(c *Context) {
	l.r.HandleContext(c.Context)
}

func (l *WKHttp) handlersToGinHandleFuncs(handlers []HandlerFunc) []gin.HandlerFunc {
	newHandlers := make([]gin.HandlerFunc, 0, len(handlers))
	if handlers != nil {
		for _, handler := range handlers {
			newHandlers = append(newHandlers, l.WKHttpHandler(handler))
		}
	}
	return newHandlers
}

// 基本认证
func (l *WKHttp) BasicAuthMiddleware(username, password string) HandlerFunc {
	return func(c *Context) {
		_username, _password, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "请先登录！",
			})
			return
		}

		if _username != username || _password != password {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "用户名或密码错误！",
			})
			return
		}

		c.Request.SetBasicAuth(username, password)
		c.Next()
	}
}

// AuthMiddleware 认证中间件
func (l *WKHttp) AuthMiddleware(cache cache.Cache, tokenPrefix string) HandlerFunc {

	return func(c *Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "token不能为空，请先登录！",
			})
			return
		}
		uidAndName := GetLoginUID(token, tokenPrefix, cache)
		if strings.TrimSpace(uidAndName) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "请先登录！",
			})
			return
		}
		uidAndNames := strings.Split(uidAndName, "@")
		if len(uidAndNames) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "token有误！",
			})
			return
		}
		c.Set("uid", uidAndNames[0])
		c.Set("name", uidAndNames[1])
		if len(uidAndNames) > 2 {
			c.Set("role", uidAndNames[2])
		}
		c.Next()
	}
}

type OperationLog struct {
	LogID    int64  `db:"log_id"`
	UID      string `db:"uid"`
	Username string `db:"username"`
	Method   string `db:"method"`
	Path     string `db:"path"`
	IP       string `db:"ip"`
	Payload  string `db:"payload"`
	Response string `db:"response"`
	ErrorMsg string `db:"error_msg"`
	Status   int    `db:"status"`
	Duration int64  `db:"duration"`
}

// 管理员操作记录中间件
func (l *WKHttp) AdminOperateRecordMiddleware(session *dbr.Session) HandlerFunc {
	var bufferSize = 1024
	return func(c *Context) {
		if c.Request.Method != http.MethodGet {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.lg.Error("读取body失败", zap.Error(err))
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			record := &OperationLog{
				IP:       c.ClientIP(),
				Method:   c.Request.Method,
				Path:     c.Request.URL.Path,
				UID:      c.GetLoginUID(),
				Username: c.GetLoginName(),
				Payload:  "",
				Response: "",
				Status:   0,
				Duration: 0,
			}
			// 上传文件时候 中间件日志进行裁断操作
			if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
				record.Payload = "[文件]"
			} else {
				if len(body) > bufferSize {
					record.Payload = "[超出记录长度]"
				} else {
					record.Payload = string(body)
				}
			}

			writer := responseBodyWriter{
				ResponseWriter: c.Writer,
				body:           &bytes.Buffer{},
			}
			c.Writer = writer
			now := time.Now()

			c.Next()

			record.ErrorMsg = c.Errors.ByType(gin.ErrorTypePrivate).String()
			record.Status = c.Writer.Status()
			record.Duration = time.Since(now).Milliseconds()
			record.Response = writer.body.String()

			if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
				strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
				strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
				strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
				strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
				strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
				strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
				strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
				strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
				if len(record.Response) > bufferSize {
					// 截断
					record.Response = "超出记录长度"
				}
			}

			// 存储 record 到mysql
			_, err = session.InsertInto("operation_log").Columns(util.AttrToUnderscore(record)...).Record(record).Exec()
			if err != nil {
				c.lg.Error("存储操作日志失败", zap.Error(err))
			}
		}
		c.Next()
	}
}

// GetLoginUID GetLoginUID
func GetLoginUID(token string, tokenPrefix string, cache cache.Cache) string {
	uid, err := cache.Get(tokenPrefix + token)
	if err != nil {
		return ""
	}
	return uid
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// RouterGroup RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
	L *WKHttp
}

func newRouterGroup(g *gin.RouterGroup, l *WKHttp) *RouterGroup {
	return &RouterGroup{RouterGroup: g, L: l}
}

// POST POST
func (r *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.POST(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// GET GET
func (r *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.GET(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// DELETE DELETE
func (r *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.DELETE(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// PUT PUT
func (r *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.PUT(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// CORSMiddleware 跨域
func CORSMiddleware() HandlerFunc {

	return func(c *Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, token, accept, origin, Cache-Control, X-Requested-With, appid, noncestr, sign, timestamp")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE,PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
