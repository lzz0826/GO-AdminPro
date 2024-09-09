package context

import (
	"AdminPro/internal/glog/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

type Handler func(ctx *Context)

type Context struct {
	*gin.Context
	Trace string //可添加自訂 Header
}

func Background(c *gin.Context) *Context {
	return &Context{c, c.GetHeader(viper.GetString("http.headerTrace"))}
}

func (c *Context) Info(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Info(args...)
}

func (c *Context) Infof(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Infof(template, args...)
}

func (c *Context) Warn(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Warn(args...)
}

func (c *Context) Warnf(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Warnf(template, args...)
}

func (c *Context) Error(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Error(args...)
}

func (c *Context) Errorf(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Errorf(template, args...)
}

func funcName(trace string) string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName) + " " + trace
}
