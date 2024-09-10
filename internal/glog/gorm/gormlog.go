package gorm

import (
	"AdminPro/internal/glog"
)

// 紀錄GORM 沒有 上下文的日誌
type gormZapLogger struct {
}

func NewGormZapLogger() *gormZapLogger {
	return &gormZapLogger{}
}

func (gormZapLogger) Printf(prefix string, v ...interface{}) {
	if len(v) == 4 {
		glog.Infof("[gorm-sql]:time=%v | row=%v | pos=%v |sql=%v", v[1], v[2], v[0], v[3])
		return
	}
	if len(v) == 5 {
		glog.Infof("[gorm-sql]:time=%v | row=%v | pos=%v | sql=%v | err=%v", v[2], v[3], v[0], v[4], v[1])
		return
	}
	glog.Infof("[gorm-sql]:%v", v)
}
