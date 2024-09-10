package gorm

import (
	"AdminPro/internal/glog"
	"context"
)

// 紀錄GORM 有 上下文的日誌
type Writer interface {
	Printf(context.Context, string, ...interface{})
}

type GormWriter struct {
}

func (w *GormWriter) Printf(ctx context.Context, format string, v ...interface{}) {
	if len(v) == 4 {
		glog.InfofWithContext(ctx, "[gorm-sql]:time=%v | row=%v | pos=%v |sql=%v", v[1], v[2], v[0], v[3])
		return
	}
	if len(v) == 5 {
		glog.InfofWithContext(ctx, "[gorm-sql]:time=%v | row=%v | pos=%v | sql=%v | err=%v", v[2], v[3], v[0], v[4], v[1])
		return
	}
	glog.InfofWithContext(ctx, "[gorm-sql]:%v", v)
}
