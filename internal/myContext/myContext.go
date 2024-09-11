package myContext

import (
	"AdminPro/internal/glog/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

type Handler func(ctx *MyContext)

type MyContext struct {
	*gin.Context
	Trace string //可添加自訂 Header
}

func Background(c *gin.Context) *MyContext {
	return &MyContext{c, c.GetHeader(viper.GetString("http.headerTrace"))}
}

func (c *MyContext) Info(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Info(args...)
}

func (c *MyContext) Infof(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Infof(template, args...)
}

func (c *MyContext) Warn(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Warn(args...)
}

func (c *MyContext) Warnf(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Warnf(template, args...)
}

func (c *MyContext) Error(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Error(args...)
}

func (c *MyContext) Errorf(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Errorf(template, args...)
}

func funcName(trace string) string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName) + " " + trace
}
