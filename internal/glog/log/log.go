package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	ZapLog *zap.SugaredLogger // 簡易版日誌文件
	// Logger *zap.Logger // 功能強大些的日誌
	logLevel = zap.NewAtomicLevel()
)

func init() {
	logConfLevel := viper.GetString("glog.logConfLevel")
	logPath := viper.GetString("glog.logPath")
	logType := viper.GetString("glog.logType")
	err := InitLog(logConfLevel, logPath, logType)
	if err != nil {
		panic("failed to init glog: " + err.Error())
	}
}

// InitLog 初始日誌
func InitLog(logConfLevel, logPath, logType string) error {
	loglevel := zapcore.InfoLevel
	switch logConfLevel {
	case "INFO":
		loglevel = zapcore.InfoLevel
	case "ERROR":
		loglevel = zapcore.ErrorLevel
	}
	setLevel(loglevel)

	var core zapcore.Core

	logPath = logPath + "dz-admin.log"

	// 打印至文件中
	if logType == "file" {
		c := zap.NewProductionEncoderConfig()
		c.EncodeTime = zapcore.ISO8601TimeEncoder
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    128, // MB
			LocalTime:  true,
			Compress:   true,
			MaxBackups: 8, // 最多保留 n 個備份
		})

		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(c),
			w,
			logLevel,
		)
	} else {
		// 打印在控制台
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), logLevel)
	}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	ZapLog = logger.Sugar()

	return nil
}

func setLevel(level zapcore.Level) {
	logLevel.SetLevel(level)
}
