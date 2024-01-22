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

// GetAllPermitList 查詢所有權限
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

// 查詢所指定管理員所有權限
func GetAdminAllPermits(ctx *gin.Context) {
	adminId := ctx.PostForm("adminId")
	if adminId == "" {
		ctx.JSON(http.StatusOK, tool.RespFail(tool.MissingParameters.Code, tool.MissingParameters.Msg, nil))
		return
	}
	permits, err := service.GetAllPermitByAdminId(adminId)

	if err != nil {
		ctx.JSON(http.StatusOK, tool.GetResponseForError(err))
		return
	}

	newPermits := service.RemoveDuplicatesPermits(permits)

	vo := adminVo.PermitListVO{
		PermitList: newPermits,
	}
	ctx.JSON(http.StatusOK, tool.RespOk(vo))

}

// GetAdminExtraPermits 查詢管理員 額外設置 的權限
func GetAdminExtraPermits(ctx *gin.Context) {
	adminId := ctx.PostForm("adminId")
	if adminId == "" {
		ctx.JSON(http.StatusOK, tool.RespFail(tool.MissingParameters.Code, tool.MissingParameters.Msg, nil))
		return
	}
	permits, err := service.GetPermitsByAdminId(adminId)

	if err != nil {
		ctx.JSON(http.StatusOK, tool.GetResponseForError(err))
		return
	}
	vo := adminVo.PermitListVO{
		PermitList: permits,
	}
	ctx.JSON(http.StatusOK, tool.RespOk(vo))

}

// AddAdminPermits 為管理員添加權限
func AddAdminPermits(ctx *gin.Context) {
	var request struct {
		AdminId    string   `json:"adminId" binding:"required"`
		PermitsIds []string `json:"permitsIds" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adminId := request.AdminId
	permitsIds := request.PermitsIds

	currentAdminId := service.GetCurrentAdminId(ctx)

	err := service.AddAdminPermits(adminId, permitsIds, currentAdminId)

	if err != nil {
		ctx.JSON(http.StatusOK, tool.GetResponseForError(err))
		return
	}
	ctx.JSON(http.StatusOK, tool.RespOk(tool.Success.Msg))

}
