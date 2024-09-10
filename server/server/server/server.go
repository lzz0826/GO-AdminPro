package server

import (
	"AdminPro/server/config"
	"AdminPro/server/route"
	"github.com/gin-gonic/gin"
)

// Run 配置并启动服务
func Run(httpServer *gin.Engine) {
	// 服务配置
	serverConfig := config.GetServerConfig()

	// gin 运行时 release debug test
	gin.SetMode(serverConfig["ENV"])

	httpServer = gin.Default()

	// 注册路由
	routes.RegisterRoutes(httpServer)

	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	// 启动服务
	err := httpServer.Run(serverAddr)

	if nil != err {
		panic("server run errors: " + err.Error())
	}
}
