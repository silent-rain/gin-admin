// Package log 文件日志
package log

import (
	"github.com/silent-rain/gin-admin/pkg/conf"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志写入文件日志配置
func newFileSyncer(logger *conf.LoggerConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logger.Filename,   // 日志文件位置
		MaxSize:    logger.MaxSize,    // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: logger.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     logger.MaxAge,     // 保留旧文件的最大天数
		Compress:   false,             // 是否压缩/归档旧文件
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
