// mysql db drives
package mysql

import (
	"AdminPro/config"
	myGorm "AdminPro/internal/glog/gorm"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
	"time"
)

//可以在這裡封裝所有DB 在下面分別init
//type AllGormDB struct {
//	GormDb01 *gorm.DB //主
//	GormDb02 *gorm.DB //從
//
//}

var GormDb *gorm.DB
var GormDbErr error

func init() {
	//  get db config
	config.Init()
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pwd"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.name"),
		viper.GetString("mysql.charset"),
	)

	// 使用你自定義的 GormWriter（不需要 logFile）
	myWriter := &myGorm.GormWriter{}

	// 初始化 GORM logger
	myLogger := myGorm.NewGormLogger(myWriter, logger.Config{
		SlowThreshold:             time.Millisecond * 200, // 超過這時間標示 SLOW SQL
		LogLevel:                  logger.Info,            // 你想記錄的層級
		IgnoreRecordNotFoundError: true,                   // 忽略 ErrRecordNotFound
		Colorful:                  false,                  // 關閉顏色，寫檔建議 false
	})

	// open GORM connection
	GormDb, GormDbErr = gorm.Open(mysql.Open(dbDSN),
		//GORM配置
		&gorm.Config{
			Logger: myLogger, //GORM 自訂日誌
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //表名保持与模型名一致而不进行复数化
			},
		})
	if GormDbErr != nil {
		panic("failed to connect to database: " + GormDbErr.Error())
	}

	// set connection pool settings
	sqlDB, err := GormDb.DB()
	if err != nil {
		panic("failed to get DB instance: " + err.Error())
	}

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(viper.GetString("mysql.pool.max_open_conns"))
	sqlDB.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(viper.GetString("mysql.pool.max_idle_conns"))
	sqlDB.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <= 0 will be forever
	dbMaxLifetimeConns, _ := strconv.Atoi(viper.GetString("max_lifetime_conns"))
	sqlDB.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	// check db connection at once to avoid connect failed
	// else errors will be reported until the first SQL operation
	if err := sqlDB.Ping(); err != nil {
		panic("database connect failed: " + err.Error())
	}
}
