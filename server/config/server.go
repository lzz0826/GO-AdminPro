package config

import "github.com/spf13/viper"

// GetServerConfig 服务配置 防止变量污染故用函数组织
func GetServerConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)

	serverConfig["HOST"] = viper.GetString("server.HOST") //监听地址
	serverConfig["PORT"] = viper.GetString("server.PORT") //监听端口
	serverConfig["ENV"] = viper.GetString("server.ENV")   // 环境模式 release/debug/test
	return
}
