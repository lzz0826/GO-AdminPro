package main

import (
	"AdminPro/common/mysql"
	"AdminPro/config"
	routes "AdminPro/route"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	// 服务停止时清理数据库链接
	defer func() {
		if mysql.GormDb != nil {
			sqlDB, err := mysql.GormDb.DB()
			if err != nil {
				panic("failed to get DB instance: " + err.Error())
			}
			_ = sqlDB.Close()
		}
	}()
	//HTTP 启动服务
	Run(HttpServer)

	//listenUrl := fmt.Sprintf("0.0.0.0:%s", config.GetConfig().Room.SocketPort)
	//listener, err := net.Listen("tcp", listenUrl)
	//if err != nil {
	//	glog.Infof("Error starting TCP server: %v", err)
	//	os.Exit(1)
	//}
	//defer listener.Close()
	//glog.Infof("tcp %s", listenUrl)
	//
	//ctx, cancel := context.WithCancel(context.Background()) //cancel
	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		return
	//	default:
	//		conn, err := listener.Accept()
	//		if err != nil {
	//			glog.Errorf("Connection accept error:%s", err)
	//			continue
	//		}
	//		utils.GoSafe(func() {
	//			processor.TcpHandleConnection(ctx, conn)
	//		})
	//	}
	//}

}

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
