package base

import (
	"AdminPro/common/enum"
	"AdminPro/common/jwt"
	"AdminPro/common/model"
	"AdminPro/common/utils"
	"AdminPro/internal/glog"
	"AdminPro/internal/myContext"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var headerTrace = viper.GetString("http.headerTrace")

func WebResp(c *myContext.MyContext, errCode enum.ResponseCodeEnum, data interface{}, Msg string) {
	if data == nil {
		data = struct{}{}
	}
	respMap := map[string]interface{}{"code": errCode, "data": data, "message": Msg}
	c.Header(headerTrace, c.Trace)
	c.JSON(200, respMap)
}

func WebRespFromCommonResp[T any](c *myContext.MyContext, data model.CommonResponse[T]) {
	// respMap := map[string]interface{}{"code": data.Code, "data": data.Data, "message": data.Msg}
	c.Header(headerTrace, c.Trace)
	c.JSON(200, data)
}

// 健康狀態
func Health(c *myContext.MyContext) {
	WebResp(c, enum.HEALTH_STATUS_OK, nil, enum.GetResponseMsg(enum.HEALTH_STATUS_OK))
}

func TraceLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先從header獲取trace,如果不在
		var trace = c.GetHeader(headerTrace)
		if trace == "" {
			trace = utils.GenerateId()
		}
		c.Request.Header.Set(headerTrace, trace)
		// sql語句trace
		glog.GetUUid(trace)
		c.Next()
	}
}

func GetAdminIdByGinContext(c *myContext.MyContext) (string, error) {
	tokenData, err := GetTokenDataByGinContext(c, "Authorization")
	if err != nil {
		return "", err
	}
	return tokenData.AdminId, nil
}

func GetTokenDataByGinContext(c *myContext.MyContext, tokenKey string) (*jwt.Claims, error) {
	tokenData, err := ParseToken(c.GetHeader(tokenKey))
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}

// 解析Token
func ParseToken(tokenStr string) (*jwt.Claims, error) {
	tokendata, err := jwt.GetTokenData(tokenStr)
	if err != nil {
		return nil, err
	}
	return tokendata, nil
}
