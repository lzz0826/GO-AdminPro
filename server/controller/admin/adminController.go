package admin

import (
	"AdminPro/admin/service"
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/dao/service/admin"
	"AdminPro/vo/model/adminVo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddAdmin 添加管理員
func AddAdmin(ctx *gin.Context) {
	var request struct {
		Username  string `json:"username" binding:"required"`
		Password  string `json:"password" binding:"required"`
		AdminName string `json:"adminName" binding:"required"`
		NickName  string `json:"nickName" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := request.Username
	password := request.Password
	adminName := request.AdminName
	nickName := request.NickName

	currentAdminId := service.GetCurrentAdminId(ctx)

	err := service.AddAdmin(username, password, adminName, nickName, currentAdminId)
	if err != nil {
		ctx.JSON(http.StatusOK, tool.GetResponseForError(err))
		return
	}
	ctx.JSON(http.StatusOK, tool.RespOk(tool.Success.Msg))

}

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
