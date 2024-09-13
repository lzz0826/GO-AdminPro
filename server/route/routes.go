package routes

import (
	"AdminPro/common/jwt"
	"AdminPro/common/middleware"
	"AdminPro/server/controller/admin"
	"AdminPro/server/controller/index"
	"AdminPro/server/controller/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	unprotected := router.Group("/")
	{
		unprotected.GET("/", index.IndexHome)
		unprotected.POST("/login", admin.Login)
	}

	protected := router.Group("/")
	protected.Use(jwt.JwtAuthMiddleware())

	protected.GET("/Logout", admin.Logout)
	
	protected.Use(middleware.CheckPermission())
	{
		protected.GET("/user/:id", user.GetById)
		protected.GET("/GetAllRoleList", admin.GetAllRoleList)
		protected.GET("/GetAllPermitList", admin.GetAllPermitList)
		protected.GET("/GetAllAdminList", admin.GetAllAdminList)

		//CommonResponse
		protected.GET("/GetAllAdminCommonResponse", admin.GetAllAdminCommonResponse)

		protected.GET("/GetAdminRole", admin.GetAdminRole)
		protected.GET("/GetAdminExtraPermits", admin.GetAdminExtraPermits)
		protected.GET("/GetAdminAllPermits", admin.GetAdminAllPermits)
		protected.GET("/GetRolePermits", admin.GetRolePermits)

		protected.POST("/AddAdmin", admin.AddAdmin)
		protected.POST("/AddRole", admin.AddRole)
		protected.POST("/AddRolePermits", admin.AddRolePermits)
		protected.POST("/AddAdminRoles", admin.AddAdminRoles)
		protected.POST("/AddAdminPermits", admin.AddAdminPermits)
		protected.POST("/RemoveRolePermits", admin.RemoveRolePermits)
		protected.POST("/RemoveAdminPermits", admin.RemoveAdminPermits)
		protected.POST("/RemoveAdminRoles", admin.RemoveAdminRoles)

	}

}
