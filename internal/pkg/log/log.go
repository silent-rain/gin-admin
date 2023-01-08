/*
 * @Author: silent-rain
 * @Date: 2023-01-07 22:02:42
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:31:35
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/log/log.go
 * @Descripttion: 日志
 */
package log

import (
	"os"

	"gin-admin/internal/pkg/conf"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logTmFmt = "2006-01-02 15:04:05.000"
)

// Init 初始化日志输出配置
func Init() {
	getLogger()
}

// 获取日志配置
func getLogger() {
	consoleEncoder := getConsoleEncoder()
	fileEncoder := getFileEncoder()
	writeSyncer := getWriteSyncer()
	levelEnabler := getLevelEnabler()
	newCore := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), levelEnabler), // 写入控制台
		zapcore.NewCore(fileEncoder, writeSyncer, levelEnabler),                // 写入文件
	)
	logger := zap.New(newCore, zap.AddCaller())
	// 重新配置全局变量
	zap.ReplaceGlobals(logger)
}

// 输出日志到文件
func getFileEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:       "ts",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller_line",
			FunctionKey:   zapcore.OmitKey,
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			// 默认换行符 \n
			LineEnding: zapcore.DefaultLineEnding,
			// 日志等级
			EncodeLevel: zapcore.CapitalLevelEncoder,
			// 时间序列化成浮点秒数
			EncodeTime:     zapcore.TimeEncoderOfLayout(logTmFmt),
			EncodeDuration: zapcore.SecondsDurationEncoder,
			// 路径编码器, 以 包名/文件名:行数 格式序列化
			EncodeCaller: zapcore.ShortCallerEncoder,
		})
}

// 输出日志到控制台
func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(logTmFmt)
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志写入文件日志配置
func getWriteSyncer() zapcore.WriteSyncer {
	config := conf.Instance().LoggerConfig
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 自定义日志级别
func getLevelEnabler() zapcore.Level {
	var level = zapcore.DebugLevel
	switch conf.Instance().LoggerConfig.Level {
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.DebugLevel
	}
	return level
}

// 自定义日志级别显示
// func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString("[" + level.CapitalString() + "]")
// }

// 自定义时间格式显示
// func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString("[" + t.Format(logTmFmt) + "]")
// }

// 自定义行号显示
// func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString("[" + caller.TrimmedPath() + "]")
// }
