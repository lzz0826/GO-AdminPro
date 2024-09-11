package controller

import (
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"github.com/gin-gonic/gin"
)

func WebRespFromCommonResp[T any](c *gin.Context, data model.CommonResponse[T]) {
	c.JSON(200, data)
}

func CheckParams[T any](c *gin.Context, params *T) bool {
	if err := c.ShouldBind(params); err != nil {
		response := new(model.CommonResponse[any])
		fromError := response.Failure(enum.PARAM_ERROR)
		WebRespFromCommonResp(c, *fromError)
		return false
	}
	return true
}
