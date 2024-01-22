package admin

import (
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/dao/service/admin"
	"AdminPro/vo/model/adminVo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAdminList(ctx *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(ctx)
	admins, err := admin.GetAllAdminList(&pagination)
	if err != nil {
		ctx.JSON(http.StatusOK, tool.RespFail(tool.SelectFail.Code, tool.SelectFail.Msg, nil))
		return
	}
	vo := adminVo.AdminListVO{
		AdminList: admins,
	}
	ctx.JSON(http.StatusOK, tool.RespOk(vo))
}
