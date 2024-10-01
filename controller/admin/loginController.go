package admin

import (
	"AdminPro/common/tool"
	"AdminPro/controller"
	"AdminPro/server/admin"
	"AdminPro/server/tonke"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Gin 的上下文（gin.Context）是为每个请求单独创建的，在请求完成后就会销毁。
func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	vo, err := admin.CheckUserAndPassword(ctx, username, password)
	if err != nil {
		ctx.JSON(http.StatusOK, tool.RespFail(err.Code(), err.Msg(), nil))
		return
	}
	ctx.JSON(http.StatusOK, tool.RespOk(vo))
}

func Logout(c *gin.Context) {
	adminId := controller.GetCurrentAdminId(c)

	fmt.Printf("%+v\n", adminId)

	// 1. 使 JWT Token 失效，可以將 token 加入失效列表
	// ...

	// 2. 清理用戶相關的會話信息
	admin.RemovePermissionByAdminId(adminId)
	// ...

	// 3. 清理客戶端存儲的 Token，比如清除 cookie 或者 localStorage
	tonke.RemoveTokenToRides(c, adminId)

	// 返回登出成功的消息
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
