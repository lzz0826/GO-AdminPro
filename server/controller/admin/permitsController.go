package admin

import (
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/dao/service/admin"
	"AdminPro/vo/model/adminVo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllPermitList(ctx *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(ctx)
	permits, err := admin.GetAllPermitList(&pagination)
	if err != nil {
		ctx.JSON(http.StatusOK, tool.RespFail(tool.SelectFail.Code, tool.SelectFail.Msg, nil))
		return
	}
	vo := adminVo.PermitListVO{
		PermitList: permits,
	}
	ctx.JSON(http.StatusOK, tool.RespOk(vo))

}
