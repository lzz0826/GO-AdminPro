package main

import (
	"AdminPro/common/mysql"
	"AdminPro/server/server/server"
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
	server.Run(HttpServer)

}
