package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func init() {
	Init()
}

// mtsql.go init()會優先在 viper init()執行 在此給mysql調用
func Init() {

	viper.SetConfigName("config") // 读取yaml配置文件
	//viper.SetConfigName("config") // 读取json配置文件
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath("server/.") // 设置配置文件和可执行二进制文件在用一个目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore errors if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another errors was produced
			log.Println("read config errors")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}

func Test() {
	fmt.Println("获取配置文件的mysql.url", viper.GetString(`mysql.host`))
	fmt.Println("获取配置文件的mysql.username", viper.GetString(`mysql.port`))
	fmt.Println("获取配置文件的mysql.password", viper.GetString(`mysql.name`))
	fmt.Println("获取配置文件的mysql.max_open_conns", viper.GetString(`mysql.pool.max_open_conns`))

}
