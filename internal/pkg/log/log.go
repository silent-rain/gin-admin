/*
 * @Author: silent-rain
 * @Date: 2023-01-07 22:02:42
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-11 00:55:04
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/log/log.go
 * @Descripttion: 日志
 */
package log

import (
	"fmt"
	"os"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
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
	jsonEncoder := getJsonEncoder()
	levelEnabler := getLevelEnabler()
	newCore := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), levelEnabler),     // 写入控制台
		zapcore.NewCore(jsonEncoder, newWriteFileSyncer(), levelEnabler),           // 写入文件
		zapcore.NewCore(getDbJsonEncoder(), newDbZapLoggerAsyncer(), levelEnabler), // 写入数据库
	)
	logger := zap.New(newCore, zap.AddCaller())
	// 重新配置全局变量
	zap.ReplaceGlobals(logger)
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

// 输出日志到控制台
func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(logTmFmt)
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志写入文件日志配置
func newWriteFileSyncer() zapcore.WriteSyncer {
	config := conf.Instance().LoggerConfig
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Filename,   // 日志文件位置
		MaxSize:    config.MaxSize,    // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: config.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     config.MaxAge,     // 保留旧文件的最大天数
		Compress:   false,             // 是否压缩/归档旧文件
		LocalTime:  true,
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

// 数据库异步日志
type dbZapLoggerAsyncer struct {
	*zap.Logger
}

// zapcore 输出配置
func getDbJsonEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(
		zapcore.EncoderConfig{
			TimeKey:       "ts",
			LevelKey:      "level",
			NameKey:       "dblogger",
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

// 创建数据库异步日志对象
func newDbZapLoggerAsyncer() *dbZapLoggerAsyncer {
	base := &dbZapLoggerAsyncer{
		new(zap.Logger),
	}
	// base.WithOptions(zap.AddCaller(), zap.AddCallerSkip(2))
	return base
}

// 定义Write方法以实现Sink接口
func (d dbZapLoggerAsyncer) Write(p []byte) (n int, err error) {
	fmt.Printf("dbZapLoggerAsyncer ================ %#v", string(p))
	// 返回写入日志的长度,以及错误
	return len(p), nil
}

// 定义Close方法以实现Sink接口
func (d dbZapLoggerAsyncer) Close() error {
	// 涉及不到关闭对象的问题, 所以return就可以
	return nil
}

// 定义Sync方法以实现Sink接口
func (d *dbZapLoggerAsyncer) Sync() error {
	// 涉及不到缓存同步问题, 所以return就可以
	return nil
}

func Debug(ctx *gin.Context, msg string, fields ...zap.Field) {
	traceId := utils.GetTraceId(ctx)
	userId := utils.GetUserId(ctx)
	zap.L().Debug(msg,
		append(fields,
			zap.String("trace_id", traceId),
			zap.Uint("user_id", userId))...,
	)
}

func Debugf(ctx *gin.Context, template string, args ...interface{}) {
	// traceId := utils.GetTraceId(ctx)
	// userId := utils.GetUserId(ctx)
	zap.S().Debugf(template, args)
}
