package tool

type Status struct {
	Code int
	Msg  string
}

var statusList = []Status{
	Success,
	SystemError,
	LoginError,
	UserError,
	PasswordError,
	NotPermissions,
	SelectFail,
	AddRolePermitsFail,
	NotFindRole,
	NotFinPermit,
}

// 添加狀態 上面statusList也需要添加 才能收尋到
var (
	//系統 0
	Success     = Status{Code: 0, Msg: "成功"}
	SystemError = Status{Code: -1, Msg: "失敗"}

	//登入 1000
	LoginError    = Status{Code: 1000, Msg: "登入失敗"}
	UserError     = Status{Code: 1001, Msg: "帳號錯誤"}
	PasswordError = Status{Code: 1002, Msg: "密碼錯誤"}

	//權限 2000
	NotPermissions = Status{Code: 2000, Msg: "沒有權限"}
	NotFindRole    = Status{Code: 2001, Msg: "找不到指定角色"}
	NotFinPermit   = Status{Code: 2002, Msg: "找不到指定權限"}

	//API 3000
	SelectFail         = Status{Code: 3000, Msg: "查詢失敗"}
	AddRolePermitsFail = Status{Code: 3001, Msg: "角色添加權限失敗"}
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
