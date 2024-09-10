package admin

import (
	"AdminPro/admin/service"
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/dao/service/admin"
	"AdminPro/internal/context"
	"AdminPro/internal/glog"
	"AdminPro/server/controller/base"
	"AdminPro/vo/model/adminVo"
	"fmt"
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

func GetAllAdminCommonResponse(c *gin.Context) {
	ctx := context.Background(c)
	var pagination model.Pagination
	if err := ctx.ShouldBind(&pagination); err != nil {
		fmt.Printf("参数绑定错误:%s\n", err)
	}
	//測試解析Token GetTokenDataByGinContext
	//response := new(model.CommonResponse[jwt.Claims])
	//
	//ginContext, err := base.GetTokenDataByGinContext(ctx)
	//if err != nil {
	//	//ctx.JSON(http.StatusOK, tool.RespFail(tool.SelectFail.Code, tool.SelectFail.Msg, nil))
	//	fromError := response.FailureFromError(err.Error())
	//	base.WebRespFromCommonResp(ctx, *fromError)
	//}
	//from := response.SuccessFrom(enum.GetResponseMsg(enum.SUCCESS), *ginContext)
	response := new(model.CommonResponse[adminVo.AdminListVO])
	admins, err := admin.GetAllAdminList(&pagination)
	if err != nil {
		//ctx.JSON(http.StatusOK, tool.RespFail(tool.SelectFail.Code, tool.SelectFail.Msg, nil))
		fromError := response.FailureFromError(err.Error())
		base.WebRespFromCommonResp(ctx, *fromError)
	}
	vo := adminVo.AdminListVO{
		AdminList: admins,
	}
	glog.Infof("xXX ", vo)

	from := response.SuccessFrom(enum.GetResponseMsg(enum.SUCCESS), vo)
	base.WebRespFromCommonResp(ctx, *from)
}
