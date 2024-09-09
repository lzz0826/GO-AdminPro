package enum

type ResponseCodeEnum int

func GetResponseMsg(r ResponseCodeEnum) string {
	if e, ok := ResponseCode[r]; ok {
		return e
	}
	return ""
}

// 系統 0
const (
	SUCCESS ResponseCodeEnum = iota
	ERROR
	PARAM_ERROR
	RECORD_NOT_EXIST
	HEALTH_STATUS_OK //健康狀態
)

// 登入 1000
const (
	PARAM_VALID_FAILED            ResponseCodeEnum = iota + 1001 // 参数较验失败
	TOKEN_OUT                                                    // 登入凭证无效
	USERNAME_OR_PASSWORD_ERROR                                   // 用户名或密码错误
	PERMISSION_ACCESS_DENIED                                     // 权限不足请联系管理员
	LOGIN_FAILURE                                                // 登入失败
	OP_NOT_EXIST                                                 // 指令不存在
	USERNAME_RULE_NOT_MATCH                                      // 用户帐号必须为6-16数字+字母
	TWO_PASSWORD_NOT_EQUALS                                      // 两次密码不一致
	PASSWORD_RULE_NOT_MATCH                                      // 密码由数字 字母和!@#$%^&*长度为6-16
	EXIST_SAME_USERNAME                                          // 以存在相同的用户帐号
	OUT_OF_ALLOW_PERMISSION_RANGE                                // 超出可配置范围
	AUTH_USER_DISABLE                                            // 帐号以停用 请联系相关人员
	NEED_ENABLE_ANCESTOR_USER                                    // 父帳戶未起用
)

// 權限 2000
const (
	NotPermissions ResponseCodeEnum = iota + 2000 // 沒有權限
	NotFindRole                                   // 找不到指定角色
	NotFinPermit                                  // 找不到指定權限
)

// API 3000
const (
	SelectFail             ResponseCodeEnum = iota + 3000 // 查詢失敗
	AddRolePermitsFail                                    // 角色添添加權限失敗
	AddAdminRolesFail                                     // 管理員添加角色失敗
	AddAdminPermitsFail                                   // 管理員添加權限失敗
	RemoveRolePermitsFail                                 // 移除角色所屬的權限失敗
	RemoveAdminPermitsFail                                // 移除管理員額外的權限失敗
	RemoveAdminRolesFail                                  // 移除管理員角色失敗
	RegisterAdminFail                                     // 註冊管理員失敗
)

// 管理員 4000
const (
	NotFindAdmin ResponseCodeEnum = iota + 4000 // 找不到管理員
	AdminIsExits                                // 管理員已存在
)

// EATTypeNames
var ResponseCode = map[ResponseCodeEnum]string{
	// 系統 0
	SUCCESS:          "成功",
	ERROR:            "失敗",
	PARAM_ERROR:      "参数错误",
	RECORD_NOT_EXIST: "纪录不存在",
	HEALTH_STATUS_OK: "服務健康狀態正常",

	// 登入 1000
	PARAM_VALID_FAILED:            "参数较验失败",
	TOKEN_OUT:                     "登入凭证无效",
	USERNAME_OR_PASSWORD_ERROR:    "用户名或密码错误",
	PERMISSION_ACCESS_DENIED:      "权限不足请联系管理员",
	LOGIN_FAILURE:                 "登入失败",
	OP_NOT_EXIST:                  "指令不存在",
	USERNAME_RULE_NOT_MATCH:       "用户帐号必须为6-16数字+字母",
	TWO_PASSWORD_NOT_EQUALS:       "两次密码不一致",
	PASSWORD_RULE_NOT_MATCH:       "密码由数字 字母和!@#$%^&*长度为6-16",
	EXIST_SAME_USERNAME:           "以存在相同的用户帐号",
	OUT_OF_ALLOW_PERMISSION_RANGE: "超出可配置范围",
	AUTH_USER_DISABLE:             "帐号以停用 请联系相关人员",
	NEED_ENABLE_ANCESTOR_USER:     "父帳戶未起用",

	// 權限 2000
	NotPermissions: "沒有權限",
	NotFindRole:    "找不到指定角色",
	NotFinPermit:   "找不到指定權限",

	// API 3000
	SelectFail:             "查詢失敗",
	AddRolePermitsFail:     "角色添添加權限失敗",
	AddAdminRolesFail:      "管理員添加角色失敗",
	AddAdminPermitsFail:    "管理員添加權限失敗",
	RemoveRolePermitsFail:  "移除角色所屬的權限失敗",
	RemoveAdminPermitsFail: "移除管理員額外的權限失敗",
	RemoveAdminRolesFail:   "移除管理員角色失敗",
	RegisterAdminFail:      "註冊管理員失敗",

	// 管理員 4000
	NotFindAdmin: "找不到管理員",
	AdminIsExits: "管理員已存在",
}
