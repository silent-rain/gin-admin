// Package log 日志
package log

import (
	"github.com/silent-rain/gin-admin/global"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var dbLogger *zap.Logger

const (
	logTmFmt = "2006-01-02 15:04:05.000"
)

// Init 初始化日志输出配置
func Init() {
	getLogger()
}

// 获取日志配置
func getLogger() {
	levelEnabler := getLevelEnabler()
	jsonEncoder := getJsonEncoder()
	consoleEncoder := newConsoleEncoder()

	cfg := global.Instance().Config().Logger
	consoleSyncer := newConsoleSyncer()
	fileSyncer := newFileSyncer(cfg)
	dbSyncer := newDbAsyncer()

	newCore := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleSyncer, levelEnabler), // 写入控制台
		zapcore.NewCore(jsonEncoder, fileSyncer, levelEnabler),       // 写入文件
	)
	logger := zap.New(newCore, zap.AddCaller())
	// 重新配置全局变量
	zap.ReplaceGlobals(logger)

	// 日志写入数据库
	dbLogger = zap.New(zapcore.NewCore(jsonEncoder, dbSyncer, levelEnabler),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
}

// zapcore 输出配置
func getJsonEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(
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

// 自定义日志级别
func getLevelEnabler() zapcore.Level {
	var level = zapcore.DebugLevel
	switch global.Instance().Config().Logger.Level {
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
