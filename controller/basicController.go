package controller

import (
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"AdminPro/common/utils"
	"github.com/gin-gonic/gin"
)

// GetCurrentAdminId 从上下文中获取管理员ID
func GetCurrentAdminId(c *gin.Context) string {
	if adminId, ok := utils.GetGinContextKey(c, "adminId").(string); ok {
		return adminId
	}
	return ""
}

func WebResp(c *gin.Context, errCode enum.ResponseCodeEnum, data interface{}, Msg string) {
	if data == nil {
		data = struct{}{}
	}
	respMap := map[string]interface{}{"code": errCode, "data": data, "message": Msg}
	c.JSON(200, respMap)
}

func WebRespFromCommonResp[T any](c *gin.Context, data model.CommonResponse[T]) {
	c.JSON(200, data)
}

// 健康狀態
func Health(c *gin.Context) {
	WebResp(c, enum.HEALTH_STATUS_OK, nil, enum.GetResponseMsg(enum.HEALTH_STATUS_OK))
}

// 驗證參數
func CheckParams[T any](c *gin.Context, params *T) bool {
	if err := c.ShouldBind(params); err != nil {
		response := new(model.CommonResponse[any])
		fromError := response.Failure(enum.PARAM_ERROR)
		WebRespFromCommonResp(c, *fromError)
		return false
	}
	return true
}

// TODO 白名單驗證
//func IsWhiteIp(ip string) bool {
//	//判斷是否白名單
//	isEnable := sqldb.IpIsEnable(ip)
//	return isEnable
//}
