package glog

type gormZapLogger struct {
}

func NewGormZapLogger() *gormZapLogger {
	return &gormZapLogger{}
}

func (gormZapLogger) Printf(prefix string, v ...interface{}) {
	if len(v) == 4 {
		Infof("[gorm-sql]:time=%v | row=%v | pos=%v |sql=%v", v[1], v[2], v[0], v[3])
		return
	}
	if len(v) == 5 {
		Infof("[gorm-sql]:time=%v | row=%v | pos=%v | sql=%v | err=%v", v[2], v[3], v[0], v[4], v[1])
		return
	}
	Infof("[gorm-sql]:%v", v)
}
