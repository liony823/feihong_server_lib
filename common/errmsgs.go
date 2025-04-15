package common

const (
	// 通用错误码
	ErrDataFormatError           = "errDataFormatError"           // 数据格式有误
	ErrDBTransactionOpenFailed   = "errDBTransactionOpenFailed"   // 开启事务失败
	ErrCreateTransactionFailed   = "errCreateTransactionFailed"   // 创建事务失败
	ErrDBTransactionCommitFailed = "errDBTransactionCommitFailed" // 数据库事物提交失败
	ErrFileUploadFailed          = "errFileUploadFailed"          // 文件上传失败

	// 消息相关错误码
	ErrSenderEmpty              = "errSenderEmpty"              // 发送者不能为空
	ErrContentEmpty             = "errContentEmpty"             // 发送内容不能为空
	ErrSendMessageFailed        = "errSendMessageFailed"        // 发送消息失败
	ErrCannotProcessSendRequest = "errCannotProcessSendRequest" // 无法处理发送消息请求
	ErrSaveMessageFailed        = "errSaveMessageFailed"        // 保存消息失败
	ErrMessageSyncReceiptFailed = "errMessageSyncReceiptFailed" // 消息同步回执失败

	// 频道相关错误码
	ErrGetChannelNodeFailed            = "errGetChannelNodeFailed"            // 获取频道所在节点失败
	ErrCreateChannelFailed             = "errCreateChannelFailed"             // 创建频道失败
	ErrAddOrUpdateChannelFailed        = "errAddOrUpdateChannelFailed"        // 添加或更新频道信息失败
	ErrQueryChannelFailed              = "errQueryChannelFailed"              // 查询频道失败
	ErrAddChannelFailed                = "errAddChannelFailed"                // 添加频道失败
	ErrChannelIDEmpty                  = "errChannelIDEmpty"                  // 频道ID不能为空
	ErrPersonalChannelNotSupported     = "errPersonalChannelNotSupported"     // 暂不支持个人频道
	ErrPersonalChannelNotSupportAddSub = "errPersonalChannelNotSupportAddSub" // 个人频道不支持添加订阅者
	ErrRemoveAllBlacklistFailed        = "errRemoveAllBlacklistFailed"        // 移除所有黑名单失败
	ErrUpdateChannelInfoFailed         = "errUpdateChannelInfoFailed"         // 更新频道信息失败
	ErrRemoveAllWhitelistFailed        = "errRemoveAllWhitelistFailed"        // 移除所有白明单失败
	ErrSendChannelUpdateMessageFailed  = "errSendChannelUpdateMessageFailed"  // 发送频道更新消息失败

	// 会话相关错误码
	ErrGetConversationFailed   = "errGetConversationFailed"   // 获取conversation失败
	ErrGetRecentMessagesFailed = "errGetRecentMessagesFailed" // 获取最近消息失败

	// 用户相关错误码
	ErrSystemAccountUpdateTokenNotAllowed = "errSystemAccountUpdateTokenNotAllowed" // 系统账号不允许更新token
	ErrSlotLeaderNodeNotExist             = "errSlotLeaderNodeNotExist"             // 槽的领导节点不存在
	ErrAddSystemAccountFailed             = "errAddSystemAccountFailed"             // 添加系统账号失败
	ErrAddSystemAccountToCacheFailed      = "errAddSystemAccountToCacheFailed"      // 添加系统账号到缓存失败
	ErrGetSlotNodeFailed                  = "errGetSlotNodeFailed"                  // 获取slot所在节点失败
	ErrRemoveSystemAccountFailed          = "errRemoveSystemAccountFailed"          // 移除系统账号失败
	ErrRemoveSystemAccountFromCacheFailed = "errRemoveSystemAccountFromCacheFailed" // 移除系统账号从缓存失败
	ErrGetSystemAccountFailed             = "errGetSystemAccountFailed"             // 获取系统账号失败
	ErrUsernameOrPasswordError            = "errUsernameOrPasswordError"            // 用户名或密码错误
	ErrJwtSecretNotConfigured             = "errJwtSecretNotConfigured"             // 没有配置jwt.secret

	// 注册登录相关错误码
	ErrUsernameRegisterOff          = "errUsernameRegisterOff"          // 用户名已注册功能已关闭
	ErrRequestDataError             = "errRequestDataError"             // 请求数据格式有误
	ErrNicknameEmpty                = "errNicknameEmpty"                // 昵称不能为空
	ErrNicknameLengthInvalid        = "errNicknameLengthInvalid"        // 昵称长度无效
	ErrUsernameNotInvalid           = "errUsernameNotInvalid"           // 用户名须以字母开头，仅支持使用6～18个字母、数字、下划线自由组合
	ErrInviteCodeEmpty              = "errInviteCodeEmpty"              // 邀请码不能为空
	ErrInviteCodeNotExist           = "errInviteCodeNotExist"           // 邀请码不存在
	ErrPasswordEmpty                = "errPasswordEmpty"                // 密码不能为空
	ErrUsernameExist                = "errUsernameExist"                // 该用户名已存在
	ErrUsernameNotExist             = "errUsernameNotExist"             // 该用户名不存在
	ErrUserNotExist                 = "errUserNotExist"                 // 该用户不存在
	ErrPasswordIncorrect            = "errPasswordIncorrect"            // 密码不正确
	ErrRegisterFailed               = "errRegisterFailed"               // 注册失败
	ErrVerifyCharEmpty              = "errVerifyCharEmpty"              // 校验字符不能为空
	ErrSignatureEmpty               = "errSignatureEmpty"               // 签名字符不能为空
	ErrUserNoPublicKey              = "errUserNoPublicKey"              // 该用户未上传公钥
	ErrSignatureInfoNotExist        = "errSignatureInfoNotExist"        // 签名信息不存在
	ErrVerifySignatureError         = "errVerifySignatureError"         // 校验签名错误
	ErrSignatureError               = "errSignatureError"               // 签名错误
	ErrPublicKeyEmpty               = "errPublicKeyEmpty"               // 公钥不能为空
	ErrUserDisabledOrBanned         = "errUserDisabledOrBanned"         // 该用户不存在或被封禁/禁用
	ErrUserAlreadyUploadPublicKey   = "errUserAlreadyUploadPublicKey"   // 该用户已上传过公钥信息
	ErrVerifyTypeNotMatch           = "errVerifyTypeNotMatch"           // 验证类型不匹配
	ErrNewPasswordSameAsOld         = "errNewPasswordSameAsOld"         // 新密码不能和旧密码相同
	ErrOldPasswordIncorrect         = "errOldPasswordIncorrect"         // 旧密码错误
	ErrUpdateLoginPasswordFailed    = "errUpdateLoginPasswordFailed"    // 修改登录密码错误
	ErrGetShortNoFailed             = "errGetShortNoFailed"             // 获取短编号失败
	ErrSetTokenCacheFailed          = "errSetTokenCacheFailed"          // 设置token缓存失败
	ErrUpdateIMTokenFailed          = "errUpdateIMTokenFailed"          // 更新IM token失败
	ErrZoneNotExist                 = "errZoneNotExist"                 // 缺少手机区号
	ErrPhoneNumberNotExist          = "errPhoneNumberNotExist"          // 缺少手机号
	ErrZoneNotSupport               = "errZoneNotSupport"               // 不支持的手机区号
	ErrSendVerifyCodeFailed         = "errSendVerifyCodeFailed"         // 发送验证码失败
	ErrVerificationCodeIncorrect    = "errVerificationCodeIncorrect"    // 验证码错误
	ErrVerificationCodeExpired      = "errVerificationCodeExpired"      // 验证码已过期
	ErrVerificationCodeTypeNotMatch = "errVerificationCodeTypeNotMatch" // 验证码类型不匹配
	ErrVerificationCodeSendTooMany  = "errVerificationCodeSendTooMany"  // 验证码发送次数过多
	ErrVerificationCodeSendFailed   = "errVerificationCodeSendFailed"   // 验证码发送失败
	ErrEditShortNoNotAllowed        = "errEditShortNoNotAllowed"        // 用户名编辑功能已关闭
	ErrEditShortNoOnce              = "errEditShortNoOnce"              // 用户名只能修改一次
	ErrUpdateUserInfoFailed         = "errUpdateUserInfoFailed"         // 更新用户信息失败

	// OAuth登录相关错误码
	ErrRegistrationNotSupported = "errRegistrationNotSupported" // 不支持注册
	ErrGetLoginStatusFailed     = "errGetLoginStatusFailed"     // 获取登录状态失败
	ErrLoginStatusExpired       = "errLoginStatusExpired"       // 登录状态已过期
	ErrGetGiteeUserInfoFailed   = "errGetGiteeUserInfoFailed"   // 获取gitee用户信息失败
	ErrQueryGiteeUserInfoFailed = "errQueryGiteeUserInfoFailed" // 查询gitee用户信息失败
	ErrInsertGiteeUserFailed    = "errInsertGiteeUserFailed"    // 插入gitee user失败
	ErrQRCodeVerificationFailed = "errQRCodeVerificationFailed" // 二维码验证失败

	// 设备相关错误码
	ErrGetDeviceInfoFailed         = "errGetDeviceInfoFailed"         // 获取设备信息失败
	ErrDeviceNotFound              = "errDeviceNotFound"              // 未查询到该设备
	ErrDeleteDeviceFailed          = "errDeleteDeviceFailed"          // 删除设备失败
	ErrQueryDeviceListFailed       = "errQueryDeviceListFailed"       // 查询设备列表失败
	ErrWebLogoutFailed             = "errWebLogoutFailed"             // 退出web设备失败
	ErrPCLogoutFailed              = "errPCLogoutFailed"              // 退出PC设备失败
	ErrQueryUserOnlineStatusFailed = "errQueryUserOnlineStatusFailed" // 查询用户在线状态失败
	ErrQueryUserFriendsFailed      = "errQueryUserFriendsFailed"      // 查询用户好友失败

	// 群组相关错误码
	ErrQueryGroupListFailed        = "errQueryGroupListFailed"        // 查询群列表错误
	ErrQueryGroupCountFailed       = "errQueryGroupCountFailed"       // 查询群数量错误
	ErrGroupIDEmpty                = "errGroupIDEmpty"                // 操作群ID不能为空
	ErrOperationStatusEmpty        = "errOperationStatusEmpty"        // 操作状态不能为空
	ErrQueryGroupInfoFailed        = "errQueryGroupInfoFailed"        // 查询群信息错误
	ErrGroupNotExist               = "errGroupNotExist"               // 操作的群不存在
	ErrUnknownOperationType        = "errUnknownOperationType"        // 未知操作类型
	ErrCallIMUpdateChannelFailed   = "errCallIMUpdateChannelFailed"   // 调用IM修改channel信息服务失败
	ErrUpdateGroupInfoFailed       = "errUpdateGroupInfoFailed"       // 更新群信息失败
	ErrGroupMemberEmpty            = "errGroupMemberEmpty"            // 群成员不能为空
	ErrQueryGroupMemberFailed      = "errQueryGroupMemberFailed"      // 查询群成员错误
	ErrQueryGroupMemberCountFailed = "errQueryGroupMemberCountFailed" // 查询群成员总数错误

	// 好友相关错误码
	ErrFriendUIDEmpty                   = "errFriendUIDEmpty"                   // 好友ID不能为空
	ErrQueryFriendApplicationFailed     = "errQueryFriendApplicationFailed"     // 查询申请记录错误
	ErrApplicationNotExist              = "errApplicationNotExist"              // 申请记录不存在
	ErrUpdateApplicationFailed          = "errUpdateApplicationFailed"          // 修改申请记录错误
	ErrDeleteApplicationFailed          = "errDeleteApplicationFailed"          // 删除申请记录错误
	ErrQueryFriendAppListFailed         = "errQueryFriendAppListFailed"         // 查询好友申请列表错误
	ErrQueryApplicantInfoFailed         = "errQueryApplicantInfoFailed"         // 查询申请用户信息错误
	ErrApplicantNotExist                = "errApplicantNotExist"                // 申请者不存在
	ErrDeleteFriendFailed               = "errDeleteFriendFailed"               // 删除好友错误
	ErrUpdateFriendRelationFailed       = "errUpdateFriendRelationFailed"       // 修改好友单项关系错误
	ErrSendDeleteFriendEventFailed      = "errSendDeleteFriendEventFailed"      // 发送删除好友事件失败
	ErrQueryUserFriendSettingsFailed    = "errQueryUserFriendSettingsFailed"    // 查询用户好友设置错误
	ErrResetFriendSettingsFailed        = "errResetFriendSettingsFailed"        // 重置好友设置错误
	ErrCommitTransactionFailed          = "errCommitTransactionFailed"          // 提交事务失败
	ErrCannotAddSelfAsFriend            = "errCannotAddSelfAsFriend"            // 不能添加自己为好友
	ErrLoggedInUserNotExist             = "errLoggedInUserNotExist"             // 登录用户不存在
	ErrCheckIsFriendFailed              = "errCheckIsFriendFailed"              // 查询是否是好友失败
	ErrAlreadyFriends                   = "errAlreadyFriends"                   // 已经是好友，不能再申请
	ErrFriendRequestReceiverNotExist    = "errFriendRequestReceiverNotExist"    // 接收好友请求的用户不存在
	ErrQueryFriendInfoFailed            = "errQueryFriendInfoFailed"            // 查询好友信息错误
	ErrFriendInfoNotExist               = "errFriendInfoNotExist"               // 好友信息不存在
	ErrVerificationCodeEmpty            = "errVerificationCodeEmpty"            // 验证码不能为空
	ErrSetApplicationTokenFailed        = "errSetApplicationTokenFailed"        // 设置申请token失败
	ErrQueryFriendRequestFailed         = "errQueryFriendRequestFailed"         // 查询好友申请记录错误
	ErrQueryUserAddressBookBadgeFailed  = "errQueryUserAddressBookBadgeFailed"  // 查询用户通讯录红点信息错误
	ErrAddFriendRequestFailed           = "errAddFriendRequestFailed"           // 新增好友申请记录错误
	ErrUpdateFriendRequestFailed        = "errUpdateFriendRequestFailed"        // 修改好友申请记录错误
	ErrAddUserAddressBookBadgeFailed    = "errAddUserAddressBookBadgeFailed"    // 新增用户通讯录红点信息错误
	ErrUpdateUserAddressBookBadgeFailed = "errUpdateUserAddressBookBadgeFailed" // 修改用户通讯录红点信息错误

	// 邀请相关错误码
	ErrQueryCreatorOrAdminUidFailed = "errQueryCreatorOrAdminUidFailed" // 查询创建者或管理员的uid失败
	ErrStartEventFailed             = "errStartEventFailed"             // 开启事件失败
	ErrAddInviteDataFailed          = "errAddInviteDataFailed"          // 添加邀请数据失败
	ErrAddInviteItemFailed          = "errAddInviteItemFailed"          // 添加邀请项失败
	ErrQueryIsAdminOrCreatorFailed  = "errQueryIsAdminOrCreatorFailed"  // 查询是否管理者或创建者失败
	ErrNotGroupAdminOrCreator       = "errNotGroupAdminOrCreator"       // 你不是群主或管理员
	ErrGetAuthInfoFailed            = "errGetAuthInfoFailed"            // 获取授权信息失败
	ErrDecodeAuthJsonFailed         = "errDecodeAuthJsonFailed"         // 解码认证信息的JSON数据失败
	ErrAuthCodeNotConfirmInvite     = "errAuthCodeNotConfirmInvite"     // 授权码不是确认邀请
	ErrQueryInviteDetailFailed      = "errQueryInviteDetailFailed"      // 查询邀请详情失败
	ErrInviteInfoNotFound           = "errInviteInfoNotFound"           // 没有查询到邀请信息
	ErrInviteStatusNotPending       = "errInviteStatusNotPending"       // 邀请信息不是待邀请状态
	ErrQueryInviterUserInfoFailed   = "errQueryInviterUserInfoFailed"   // 查询邀请者的用户信息失败
	ErrInviterUserInfoNotFound      = "errInviterUserInfoNotFound"      // 没有查到邀请者的用户信息
	ErrUpdateInviteStatusFailed     = "errUpdateInviteStatusFailed"     // 更新邀请信息状态失败
	ErrUpdateInviteItemStatusFailed = "errUpdateInviteItemStatusFailed" // 更新邀请信息项状态失败
	ErrAddMemberFailed              = "errAddMemberFailed"              // 添加成员失败
	ErrGetInviteItemFailed          = "errGetInviteItemFailed"          // 获取邀请项失败
	ErrQueryMemberInfoFailed        = "errQueryMemberInfoFailed"        // 查询成员信息错误

	// 举报相关错误码
	ErrQueryChannelTypeEmpty     = "errQueryChannelTypeEmpty"     // 查询频道类型不能为空
	ErrQueryReportListFailed     = "errQueryReportListFailed"     // 查询举报列表错误
	ErrQueryReportCountFailed    = "errQueryReportCountFailed"    // 查询举报总数量错误
	ErrQueryUserInfoFailed       = "errQueryUserInfoFailed"       // 查询用户信息错误
	ErrQueryReportUserSetFailed  = "errQueryReportUserSetFailed"  // 查询举报用户集合错误
	ErrQueryReportGroupSetFailed = "errQueryReportGroupSetFailed" // 查询举报群集合错误

	// 机器人相关错误码
	ErrRobotIDEmpty                    = "errRobotIDEmpty"                    // 机器人ID不能为空
	ErrQueryRobotMenuFailed            = "errQueryRobotMenuFailed"            // 查询机器人菜单错误
	ErrQueryOperatedRobotFailed        = "errQueryOperatedRobotFailed"        // 查询操作的机器人错误
	ErrOperatedRobotNotExist           = "errOperatedRobotNotExist"           // 操作的机器人不存在
	ErrDeleteRobotMenuFailed           = "errDeleteRobotMenuFailed"           // 删除机器人菜单失败
	ErrUpdateRobotVersionFailed        = "errUpdateRobotVersionFailed"        // 修改机器人版本号错误
	ErrUpdateRobotStatusFailed         = "errUpdateRobotStatusFailed"         // 修改机器人状态信息错误
	ErrSendStreamStartFailed           = "errSendStreamStartFailed"           // 发送stream start消息失败
	ErrSendStreamEndFailed             = "errSendStreamEndFailed"             // 发送stream end消息失败
	ErrNotAllowSendToChannel           = "errNotAllowSendToChannel"           // 不允许发送消息到此频道
	ErrSendTypingFailed                = "errSendTypingFailed"                // 发送typing消息失败
	ErrUnsupportedType                 = "errUnsupportedType"                 // 不支持的type
	ErrInvalidPayload                  = "errInvalidPayload"                  // 无效的payload
	ErrGetRobotInfoFailed              = "errGetRobotInfoFailed"              // 获取机器人信息失败
	ErrRobotNotExist                   = "errRobotNotExist"                   // 机器人不存在
	ErrRobotNoAppID                    = "errRobotNoAppID"                    // 机器人没有app_id
	ErrBatchQueryRobotFailed           = "errBatchQueryRobotFailed"           // 批量查询机器人数据错误
	ErrBatchQueryRobotByUsernameFailed = "errBatchQueryRobotByUsernameFailed" // 批量通过username查询机器人数据错误
	ErrBatchQueryRobotMenuFailed       = "errBatchQueryRobotMenuFailed"       // 批量查询机器人菜单数据错误

	// 工作台相关错误码
	ErrAppIDEmpty                      = "errAppIDEmpty"                      // 应用ID不能为空
	ErrDeleteUseRecordFailed           = "errDeleteUseRecordFailed"           // 删除使用记录错误
	ErrQueryUseRecordFailed            = "errQueryUseRecordFailed"            // 查询使用记录错误
	ErrQueryAppBatchFailed             = "errQueryAppBatchFailed"             // 查询一批应用错误
	ErrQueryUserSavedAppFailed         = "errQueryUserSavedAppFailed"         // 查询用户已保存的应用错误
	ErrQueryAppFailed                  = "errQueryAppFailed"                  // 查询应用错误
	ErrAppDeletedOrUnavailable         = "errAppDeletedOrUnavailable"         // 该应用已删除或不可用
	ErrAddUseRecordFailed              = "errAddUseRecordFailed"              // 新增使用记录错误
	ErrUpdateUseRecordFailed           = "errUpdateUseRecordFailed"           // 修改使用记录错误
	ErrCategoryIDEmpty                 = "errCategoryIDEmpty"                 // 分类编号不能为空
	ErrQueryAppByCategoryFailed        = "errQueryAppByCategoryFailed"        // 通过分类查询应用错误
	ErrUpdateUserAppOrderFailed        = "errUpdateUserAppOrderFailed"        // 修改用户app顺序错误
	ErrDeleteUserAppFailed             = "errDeleteUserAppFailed"             // 删除用户app错误
	ErrQueryUserAppFailed              = "errQueryUserAppFailed"              // 查询用户某个应用错误
	ErrQueryUserMaxOrderAppFailed      = "errQueryUserMaxOrderAppFailed"      // 查询用户最大序号app错误
	ErrAddUserAppFailed                = "errAddUserAppFailed"                // 添加用户app错误
	ErrQueryCategoryFailed             = "errQueryCategoryFailed"             // 查询分类错误
	ErrQueryBannerDataFailed           = "errQueryBannerDataFailed"           // 查询横幅数据错误
	ErrCategoryNotExist                = "errCategoryNotExist"                // 该分类不存在
	ErrUpdateCategoryFailed            = "errUpdateCategoryFailed"            // 修改分类错误
	ErrDeleteCategoryFailed            = "errDeleteCategoryFailed"            // 删除分类错误
	ErrQueryAllAppFailed               = "errQueryAllAppFailed"               // 查询所有应用错误
	ErrQueryTotalCountFailed           = "errQueryTotalCountFailed"           // 查询总数量错误
	ErrSearchAppFailed                 = "errSearchAppFailed"                 // 搜索应用错误
	ErrDeleteCategoryAppFailed         = "errDeleteCategoryAppFailed"         // 删除分类下app错误
	ErrAddedAppNotExist                = "errAddedAppNotExist"                // 添加的应用不存在
	ErrQueryCategoryAppFailed          = "errQueryCategoryAppFailed"          // 查询该分类下应用错误
	ErrQueryCategoryAppMaxOrderFailed  = "errQueryCategoryAppMaxOrderFailed"  // 查询分类应用最大序号错误
	ErrAddCategoryAppFailed            = "errAddCategoryAppFailed"            // 添加分类下app错误
	ErrUpdateCategoryAppOrderFailed    = "errUpdateCategoryAppOrderFailed"    // 修改分类下app排序错误
	ErrGetCategoryAppFailed            = "errGetCategoryAppFailed"            // 获取分类下的app错误
	ErrBannerIDEmpty                   = "errBannerIDEmpty"                   // 横幅编号不能为空
	ErrBannerURLEmpty                  = "errBannerURLEmpty"                  // 横幅跳转地址不能为空
	ErrBannerCoverEmpty                = "errBannerCoverEmpty"                // 横幅封面不能为空
	ErrUpdateBannerFailed              = "errUpdateBannerFailed"              // 修改横幅错误
	ErrQueryBannerFailed               = "errQueryBannerFailed"               // 查询横幅错误
	ErrDeleteBannerFailed              = "errDeleteBannerFailed"              // 删除横幅错误
	ErrQueryAllCategoryFailed          = "errQueryAllCategoryFailed"          // 查询所有分类错误
	ErrAddBannerFailed                 = "errAddBannerFailed"                 // 新增横幅信息错误
	ErrUpdateAppInfoFailed             = "errUpdateAppInfoFailed"             // 修改应用信息错误
	ErrCategoryIDAndAppIDEmpty         = "errCategoryIDAndAppIDEmpty"         // 分类ID和应用ID均不能为空
	ErrDeleteAppFailed                 = "errDeleteAppFailed"                 // 删除应用错误
	ErrDeleteUserAppUsageRecordFailed  = "errDeleteUserAppUsageRecordFailed"  // 删除用户app使用记录错误
	ErrQueryCategoryAppNameExistFailed = "errQueryCategoryAppNameExistFailed" // 查询此分类下app是否存在此名称错误
	ErrAppNameExist                    = "errAppNameExist"                    // 此应用名已存在
	ErrAddAppFailed                    = "errAddAppFailed"                    // 新增应用错误
	ErrQueryCategoryByNameFailed       = "errQueryCategoryByNameFailed"       // 通过分类名查询分类错误
	ErrCategoryNameExist               = "errCategoryNameExist"               // 该分类名称已存在
	ErrQueryMaxOrderCategoryFailed     = "errQueryMaxOrderCategoryFailed"     // 查询最大序号分类错误
	ErrAddCategoryFailed               = "errAddCategoryFailed"               // 新增分类错误

	// 通用模块相关错误码
	ErrQueryAppModuleFailed = "errQueryAppModuleFailed" // 查询app模块错误
	ErrDeleteModuleNotExist = "errDeleteModuleNotExist" // 删除的模块不存在

	// 节点相关错误码
	ErrNodeNotExist = "errNodeNotExist" // 节点不存在

	// 压测相关错误码
	ErrStressTestNotEnabled       = "errStressTestNotEnabled"       // 没有开启压测配置，不支持压测
	ErrSlotLeaderNodeInfoNotExist = "errSlotLeaderNodeInfoNotExist" // 槽领导对应的节点信息不存在
	ErrStressMachineNotExist      = "errStressMachineNotExist"      // 压测机器不存在
)
