package main

import (
	"AdminPro/common/mysql"
	"AdminPro/common/utils"
	"AdminPro/config"
	"AdminPro/internal/glog"
	routes "AdminPro/route"
	"AdminPro/server/task/tcp/hadnler"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	go RunHttp()

	//TCP
	listenUrl := fmt.Sprintf("0.0.0.0:%s", "8040")
	listener, err := net.Listen("tcp", listenUrl)
	if err != nil {
		glog.Infof("Error starting TCP config: %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("tcp %s", listenUrl)
	glog.Infof("tcp %s", listenUrl)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	ctx, cancel := context.WithCancel(context.Background()) //cancel
	go handleSignal(ctx, cancel, quit)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				glog.Errorf("Connection accept error:%s", err)
				continue
			}
			utils.GoSafe(func() {
				hadnler.TcpHandleConnection(ctx, conn)
			})
		}
	}

}

// RunHttp 配置并启动服务
func RunHttp() {
	// 服务配置
	serverConfig := config.GetServerConfig()

	// gin 运行时 release debug test
	gin.SetMode(serverConfig["ENV"])

	HttpServer = gin.Default()

	// 注册路由
	routes.RegisterRoutes(HttpServer)

	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	// 启动服务
	err := HttpServer.Run(serverAddr)

	if nil != err {
		panic("config run errors: " + err.Error())
	}
}

// handleSignal 函数用于处理操作系统信号，实现优雅关闭应用程序的功能。
func handleSignal(ctx context.Context, cancel context.CancelFunc, c chan os.Signal) {
	// 等待接收信号并根据接收到的信号进行处理
	switch a := <-c; a {
	// 处理特定的系统信号
	case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGKILL:
		// 打印关闭消息到控制台
		fmt.Println("Shutdown quickly, bye...", a)
		// 记录关闭消息到日志
		glog.Info("Shutdown quickly, bye...", a)
		// 关闭数据库连接
		//sqldb.Close()

		// 启动新的 goroutine 进行异步清理操作
		go func() {
			// 销毁房间节点管理器实例
			//room.RoomNodeManagerInstance.Destroy()
			// 关闭 Zookeeper 会话
			//xzookeeper.ZkSessionClose()
			// 关闭令牌会话
			//xzookeeper.TokenSessionClose()
		}()
	}

	// 强制退出程序，状态码 0 表示正常退出
	os.Exit(0)
}
