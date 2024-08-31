package mongodb

import (
	"fmt"
	"path"
	"runtime"
)

type MongoLogger struct {
}

func (m *MongoLogger) Info(level int, message string, keysAndValues ...interface{}) {
	//log.ZapLog.Named(funcName()).Error(level, message, keysAndValues)
	fmt.Println(level, message, keysAndValues)
}

func (m *MongoLogger) Error(err error, message string, keysAndValues ...interface{}) {
	//log.ZapLog.Named(funcName()).Error(err, message, keysAndValues)
	fmt.Println(err, message, keysAndValues)
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName)
}
