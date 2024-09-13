package admin

import (
	"AdminPro/admin/service"
	"AdminPro/common/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	vo, err := service.CheckUserAndPassword(username, password)
	if err != nil {
		ctx.JSON(http.StatusOK, tool.RespFail(err.Code(), err.Msg(), nil))
		return
	}

	ctx.JSON(http.StatusOK, tool.RespOk(vo))
}

func Logout(c *gin.Context) {
	// 1. 使 JWT Token 失效，可以將 token 加入失效列表
	// ...

	// 2. 清理用戶相關的會話信息
	// ...

	// 3. 清理客戶端存儲的 Token，比如清除 cookie 或者 localStorage
	// ...

	// 返回登出成功的消息
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
