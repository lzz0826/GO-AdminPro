package glog

import "context"

type Writer interface {
	Printf(context.Context, string, ...interface{})
}

type GormWriter struct {
}

func (w *GormWriter) Printf(ctx context.Context, format string, v ...interface{}) {
	if len(v) == 4 {
		InfofWithContext(ctx, "[gorm-sql]:time=%v | row=%v | pos=%v |sql=%v", v[1], v[2], v[0], v[3])
		return
	}
	if len(v) == 5 {
		InfofWithContext(ctx, "[gorm-sql]:time=%v | row=%v | pos=%v | sql=%v | err=%v", v[2], v[3], v[0], v[4], v[1])
		return
	}
	InfofWithContext(ctx, "[gorm-sql]:%v", v)
}
