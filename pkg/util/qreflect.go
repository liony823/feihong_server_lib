package util

import (
	"reflect"
)

// AttrToUnderscore 获取struct的所有属性并转为下划线模式
func AttrToUnderscore(st interface{}) []string {
	t := reflect.ValueOf(st)
	vType := t.Elem().Type()
	names := make([]string, 0)
	for i := 0; i < vType.NumField(); i++ {
		if vType.Field(i).Type.Kind() == reflect.Struct {
			continue
		}
		name := vType.Field(i).Name
		if name != "" {
			names = append(names, UnderscoreName(name))
		}
	}
	return names
}

// IsZeroValue 检查一个值是否为其类型的零值
func IsZeroValue(v interface{}) bool {
	if v == nil {
		return true
	}

	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Slice, reflect.Map, reflect.Array:
		return val.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	default:
		return reflect.DeepEqual(v, reflect.Zero(val.Type()).Interface())
	}
}

// CopyStructFields 将结构体字段复制到map中
func CopyStructFields(dest *map[string]interface{}, src interface{}) error {
	if dest == nil {
		return nil
	}

	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue // 跳过非导出字段
		}

		fieldValue := v.Field(i).Interface()
		(*dest)[field.Name] = fieldValue
	}

	return nil
}
