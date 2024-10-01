package tool

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var statusList = []Status{
	Success,
	SystemError,
	SetRedisKeyError,
	LoginError,
	UserError,
	PasswordError,
	TokenError,
	CreateTokenError,
	NotPermissions,
	SelectFail,
	AddRolePermitsFail,
	NotFindRole,
	NotFinPermit,
	AddAdminRolesFail,
	NotFindAdmin,
	AddAdminPermitsFail,
	MissingParameters,
	RemoveRolePermitsFail,
	RemoveAdminPermitsFail,
	RegisterAdminFail,
	RemoveAdminRolesFail,
	AdminIsExits,
}

// 添加狀態 上面statusList也需要添加 才能收尋到
var (
	//系統 0
	Success           = Status{Code: 0, Msg: "成功"}
	SystemError       = Status{Code: -1, Msg: "失敗"}
	MissingParameters = Status{Code: 1, Msg: "缺少必要參數"}

	//Redis 100
	SetRedisKeyError = Status{Code: 100, Msg: "Redis Set錯誤"}

	//登入 1000
	LoginError       = Status{Code: 1000, Msg: "登入失敗"}
	UserError        = Status{Code: 1001, Msg: "帳號錯誤"}
	PasswordError    = Status{Code: 1002, Msg: "密碼錯誤"}
	TokenError       = Status{Code: 1003, Msg: "Token無效"}
	CreateTokenError = Status{Code: 1004, Msg: "生產Token失敗"}

	//權限 2000
	NotPermissions = Status{Code: 2000, Msg: "沒有權限"}
	NotFindRole    = Status{Code: 2001, Msg: "找不到指定角色"}
	NotFinPermit   = Status{Code: 2002, Msg: "找不到指定權限"}

	//API 3000
	SelectFail             = Status{Code: 3000, Msg: "查詢失敗"}
	AddRolePermitsFail     = Status{Code: 3001, Msg: "角色添添加權限失敗"}
	AddAdminRolesFail      = Status{Code: 3002, Msg: "管理員添加角色失敗"}
	AddAdminPermitsFail    = Status{Code: 3003, Msg: "管理員添加權限失敗"}
	RemoveRolePermitsFail  = Status{Code: 3004, Msg: "移除角色所屬的權限失敗"}
	RemoveAdminPermitsFail = Status{Code: 3005, Msg: "移除管理員額外的權限失敗"}
	RemoveAdminRolesFail   = Status{Code: 3006, Msg: "移除管理員角色失敗"}
	RegisterAdminFail      = Status{Code: 3007, Msg: "註冊管理員失敗"}

	//管理員 4000
	NotFindAdmin = Status{Code: 4000, Msg: "找不到管理員"}
	AdminIsExits = Status{Code: 4001, Msg: "管理員已存在"}
)

func GetStatusByCode(code int) Status {
	for _, status := range statusList {
		if status.Code == code {
			return status
		}
	}
	return Status{Code: -1, Msg: "未知狀態"}
}

func GetStatusByMsg(msg string) Status {
	for _, status := range statusList {
		if status.Msg == msg {
			return status
		}
	}
	return Status{Code: -1, Msg: "未知狀態"}
}

func GetStatusCodeFromError(err error) int {
	status := GetStatusByMsg(err.Error())
	return status.Code
}

func GetStatusMsgFromError(err error) string {
	status := GetStatusByMsg(err.Error())
	return status.Msg
}
