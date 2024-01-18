package admin

import (
	"AdminPro/admin/service"
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	"AdminPro/vo/model/adminVo"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetAllRoleList(ctx *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(ctx)
	roles, err := admin.GetAllRoleList(&pagination)
	if err != nil {
		ctx.JSON(http.StatusOK, tool.RespFail(tool.SelectFail.Code, tool.SelectFail.Msg, nil))
		return
	}
	vo := adminVo.RoleListVO{
		RoleList: roles,
	}
	ctx.JSON(http.StatusOK, tool.RespOk(vo))
}

// AddRole 添加角色
func AddRole(ctx *gin.Context) {
	var role adminDao.RoleDAO
	// 綁定JSON參數到結構體
	if err := ctx.BindJSON(&role); err != nil {
		// 如果解析JSON失敗，返回錯誤
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.CreateTime = time.Now()
	role.UpdateTime = time.Now()
	role.CreatorID = service.GetCurrentAdminId(ctx)
	role.UpdaterID = service.GetCurrentAdminId(ctx)

	// 调用 InsertRole 函数插入 Role 数据
	err := admin.InsertRole(role)
	if err != nil {
		ctx.JSON(http.StatusOK, tool.RespFail(tool.SystemError.Code, tool.SystemError.Msg, nil))
		return
	}
	ctx.JSON(http.StatusOK, tool.RespOk(tool.Success.Msg))
}

// AddRolePermits 為角色添加新權限
func AddRolePermits(ctx *gin.Context) {
	// 創建結構體，用於存放 JSON 資料
	var request struct {
		RoleId    string   `json:"roleId" binding:"required"`
		PermitIds []string `json:"permitIds" binding:"required"`
	}

	// 使用 ShouldBindJSON 方法綁定 JSON 資料到結構體
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// 如果綁定失敗，回應錯誤信息
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 獲取結構體中的值
	roleId := request.RoleId
	permitIds := request.PermitIds
	adminId := service.GetCurrentAdminId(ctx)

	err := service.AddRolePermits(roleId, permitIds, adminId)

	if err != nil {

		ctx.JSON(http.StatusOK, tool.GetResponseForError(err))
		return
	}
	ctx.JSON(http.StatusOK, tool.RespOk(tool.Success.Msg))

}

// AddAdminRoles 為管理員添加腳色
func AddAdminRoles(ctx *gin.Context) {
	var request struct {
		AdminId string   `json:"adminId" binding:"required"`
		RoleIds []string `json:"roleIds" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adminId := request.AdminId
	roleIds := request.RoleIds

	currentAdminId := service.GetCurrentAdminId(ctx)

	err := service.AddAdminRoles(adminId, roleIds, currentAdminId)

	if err != nil {
		ctx.JSON(http.StatusOK, tool.GetResponseForError(err))
		return
	}
	ctx.JSON(http.StatusOK, tool.RespOk(tool.Success.Msg))

}
