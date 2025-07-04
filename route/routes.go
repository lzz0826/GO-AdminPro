package routes

import (
	"AdminPro/common/jwt"
	"AdminPro/common/middleware"
	admin2 "AdminPro/controller/admin"
	"AdminPro/controller/index"
	"AdminPro/controller/mq"
	"AdminPro/controller/user"
	"AdminPro/server/task/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	unprotected := router.Group("/")
	{
		unprotected.GET("/", index.IndexHome)
		unprotected.POST("/login", admin2.Login)
	}

	protected := router.Group("/")
	protected.Use(jwt.JwtAuthMiddleware())

	protected.GET("/Logout", admin2.Logout)

	//TODO 待加到登入後的權限
	protected.POST("/TaskHttpHandler", handler.TaskHttpHandler)
	protected.POST("/RefreshToken", admin2.RefreshToken)
	//測試 Rabbitmq 發消息
	protected.GET("/SetupRabbitmqController", mq.SetupRabbitmqController)

	//測試 KafKa 收消息
	protected.POST("/SendMessageToKafka", mq.SendMessageToKafka)

	protected.Use(middleware.CheckPermission())
	{
		protected.GET("/user/:id", user.GetById)
		protected.GET("/GetAllRoleList", admin2.GetAllRoleList)
		protected.GET("/GetAllPermitList", admin2.GetAllPermitList)
		protected.GET("/GetAllAdminList", admin2.GetAllAdminList)

		//CommonResponse
		protected.GET("/GetAllAdminCommonResponse", admin2.GetAllAdminCommonResponse)

		protected.GET("/GetAdminRole", admin2.GetAdminRole)
		protected.GET("/GetAdminExtraPermits", admin2.GetAdminExtraPermits)
		protected.GET("/GetAdminAllPermits", admin2.GetAdminAllPermits)
		protected.GET("/GetRolePermits", admin2.GetRolePermits)

		protected.POST("/AddAdmin", admin2.AddAdmin)
		protected.POST("/AddRole", admin2.AddRole)
		protected.POST("/AddRolePermits", admin2.AddRolePermits)
		protected.POST("/AddAdminRoles", admin2.AddAdminRoles)
		protected.POST("/AddAdminPermits", admin2.AddAdminPermits)
		protected.POST("/RemoveRolePermits", admin2.RemoveRolePermits)
		protected.POST("/RemoveAdminPermits", admin2.RemoveAdminPermits)
		protected.POST("/RemoveAdminRoles", admin2.RemoveAdminRoles)

		//任务
		//protected.POST("/TaskHttpHandler", handler.TaskHttpHandler)

	}

}
