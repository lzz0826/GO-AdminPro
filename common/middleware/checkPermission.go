package middleware

import (
	"AdminPro/common/tool"
	"AdminPro/server/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

// 權限名 key = 方法名 db=admin_permit permit_key
func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求处理之前执行的逻辑

		// 获取处理当前请求的 handler 函数的信息
		methodName := getMethodName(c.Handler())
		adminId := admin.GetCurrentAdminId(c)

		bool := admin.CheckPermission(adminId, methodName)

		if bool == false {
			c.JSON(http.StatusOK, tool.RespFail(tool.NotPermissions.Code, tool.NotPermissions.Msg, nil))
			c.Abort() // 停止后续处理
			return
		}

		// 调用链中的下一个处理程序
		c.Next()

		// 在请求处理之后执行的逻辑
	}
}

// 获取 handler 函数的方法名
func getMethodName(handler interface{}) string {
	// 使用 runtime 获取函数的名称
	fullName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}
