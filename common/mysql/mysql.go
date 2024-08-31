// mysql db drives
package mysql

import (
	"AdminPro/server/config"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

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

	// open GORM connection
	GormDb, GormDbErr = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
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
	// else error will be reported until the first SQL operation
	if err := sqlDB.Ping(); err != nil {
		panic("database connect failed: " + err.Error())
	}
}
