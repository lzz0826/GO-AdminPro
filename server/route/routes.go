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
	protected.Use(middleware.CheckPermission())
	{
		protected.GET("/user/:id", user.GetById)
		protected.GET("/GetAllRoleList", admin.GetAllRoleList)
		protected.GET("/GetAllPermitList", admin.GetAllPermitList)

		protected.POST("/AddRole", admin.AddRole)
		protected.POST("/AddRolePermits", admin.AddRolePermits)

	}

}
