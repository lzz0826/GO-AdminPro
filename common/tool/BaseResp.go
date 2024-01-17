package tool

// BaseResp 通用的 JSON 返回結構
type BaseResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RespOk 方法創建成功的 BaseResp
func RespOk(data interface{}) BaseResp {
	return BaseResp{
		Code:    Success.Code,
		Message: Success.Msg,
		Data:    data,
	}
}

// RespFail 方法創建失敗的 BaseResp
func RespFail(statusCode int, message string, data interface{}) BaseResp {
	return BaseResp{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
}

// 根據拋出的狀態返回 BaseResp
func GetResponseForError(err error) BaseResp {
	status := GetStatusByMsg(err.Error())
	return RespFail(status.Code, status.Msg, nil)
}
