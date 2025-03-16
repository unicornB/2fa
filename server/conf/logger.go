package conf

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	// 配置 lumberjack 进行日志分割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./logs/app.log", // 日志文件路径
		MaxSize:    10,               // 每个日志文件最大大小（MB）
		MaxBackups: 3,                // 最多保留的旧日志文件数量
		MaxAge:     28,               // 日志文件最多保留的天数
		Compress:   true,             // 是否压缩旧日志文件
	}

	// 配置 logrus 使用 lumberjack 作为输出
	logrus.SetOutput(lumberjackLogger)
	logrus.SetLevel(logrus.InfoLevel) // 设置日志级别
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}
