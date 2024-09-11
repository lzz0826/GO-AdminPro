package glog

import (
	"AdminPro/internal/glog/log"
	context2 "AdminPro/internal/myContext"
	"context"
	"fmt"
	"path"
	"runtime"
)

var uuidRequest string

func GetUUid(uid string) {
	uuidRequest = uid
}

func Info(args ...interface{}) {
	args = append(args, fmt.Sprintf(" %s ", uuidRequest))
	log.ZapLog.Named(funcName()).Info(args...)
}

func Infof(template string, args ...interface{}) {
	template += " %s "
	args = append(args, uuidRequest)
	log.ZapLog.Named(funcName()).Infof(template, args...)
}

func Warn(args ...interface{}) {
	args = append(args, fmt.Sprintf(" commitId:%s ", uuidRequest))
	log.ZapLog.Named(funcName()).Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	template += " %s "
	args = append(args, uuidRequest)
	log.ZapLog.Named(funcName()).Warnf(template, args...)
}

func Error(args ...interface{}) {
	args = append(args, fmt.Sprintf(" %s", uuidRequest))
	log.ZapLog.Named(funcName()).Error(args...)
}

func Errorf(template string, args ...interface{}) {
	template += " %s "
	args = append(args, uuidRequest)
	log.ZapLog.Named(funcName()).Errorf(template, args...)
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName)
}

func InfofWithContext(ctx context.Context, template string, args ...interface{}) {
	template += " %s "
	traceID := getTraceID(ctx)
	args = append(args, traceID)
	log.ZapLog.Named(funcNameWithTrace(traceID)).Infof(template, args...)
}

func getTraceID(ctx context.Context) string {
	var traceID string
	if c, ok := ctx.(context2.MyContext); ok {
		traceID = c.Trace
	} else if c, ok := ctx.(*context2.MyContext); ok {
		traceID = c.Trace
	} else {
		if traceID, ok = ctx.Value("Trace").(string); !ok {
			traceID = ""
		}
	}

	if traceID == "" {
		traceID = uuidRequest
	}
	return traceID
}

func funcNameWithTrace(trace string) string {
	//callStack := make([]uintptr, 10)
	//frames := runtime.Callers(0, callStack)
	//callStack = callStack[:frames]
	//for index, pc := range callStack {
	//	n := runtime.FuncForPC(pc).Name()
	//	file, line := runtime.FuncForPC(pc).FileLine(pc)
	//	fmt.Println(n, index, file, line)
	//}

	pc, _, _, _ := runtime.Caller(6)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName) + " " + trace
}
