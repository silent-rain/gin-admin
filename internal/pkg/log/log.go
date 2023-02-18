/* 日志
 */
package log

import (
	"encoding/json"
	"fmt"
	"os"

	systemDAO "gin-admin/internal/dao/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/context"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
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
	consoleEncoder := getConsoleEncoder()
	jsonEncoder := getJsonEncoder()
	levelEnabler := getLevelEnabler()
	newCore := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), levelEnabler), // 写入控制台
		zapcore.NewCore(jsonEncoder, newWriteFileSyncer(), levelEnabler),       // 写入文件

	)
	logger := zap.New(newCore, zap.AddCaller())
	// 重新配置全局变量
	zap.ReplaceGlobals(logger)

	// 日志写入数据库
	dbLogger = zap.New(zapcore.NewCore(jsonEncoder, newDbLoggerAsyncer(), levelEnabler),
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
	config := conf.Instance().Logger
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
	switch conf.Instance().Logger.Level {
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
type dbLoggerAsyncer struct {
	*zap.Logger
}

// 创建数据库异步日志对象
func newDbLoggerAsyncer() *dbLoggerAsyncer {
	base := &dbLoggerAsyncer{
		new(zap.Logger),
	}
	return base
}

// Write 定义Write方法以实现Sink接口
func (d dbLoggerAsyncer) Write(p []byte) (int, error) {
	sysLog := systemModel.SystemLog{}
	if err := json.Unmarshal(p, &sysLog); err != nil {
		return len(p), err
	}
	go func() {
		systemDAO.NewSystemLogDao().Add(sysLog)
	}()
	// 返回写入日志的长度,以及错误
	return len(p), nil
}

// Close 定义Close方法以实现Sink接口
func (d dbLoggerAsyncer) Close() error {
	// 涉及不到关闭对象的问题, 所以return就可以
	return nil
}

// Sync 定义Sync方法以实现Sink接口
func (d *dbLoggerAsyncer) Sync() error {
	// 涉及不到缓存同步问题, 所以return就可以
	return nil
}

// 日志结构
type logger struct {
	ctx     *gin.Context
	zapLog  *zap.Logger
	fields  []zapcore.Field
	extends map[string]interface{} // 消息扩展字段
}

// New 创建日志对象
func New(ctx *gin.Context) *logger {
	traceId := context.GetTraceId(ctx)
	userId := context.GetUserId(ctx)
	fields := []zapcore.Field{
		zap.String("trace_id", traceId),
		zap.Uint("user_id", userId),
	}
	return &logger{
		ctx:     ctx,
		zapLog:  zap.L().WithOptions(zap.AddCallerSkip(1)),
		fields:  fields,
		extends: make(map[string]interface{}, 0),
	}
}

// WithCode 添加错误码
func (l *logger) WithCode(code errcode.ErrorCode) *logger {
	l.fields = append(l.fields, zap.Uint("error_code", uint(code)), zap.String("error_msg", code.Msg()))
	return l
}

// WithCodeError 添加响应状态码及状态码对应的信息
func (l *logger) WithCodeError(err error) *logger {
	code, ok := err.(*errcode.Error)
	if !ok {
		l.fields = append(l.fields, zap.Uint("error_code", uint(errcode.UnknownError)),
			zap.String("error_msg", errcode.UnknownError.Msg()))
		return l
	}
	l.fields = append(l.fields, zap.Uint("error_code", uint(code.Code)), zap.String("error_msg", code.Msg))
	return l
}

// WithField 添加扩展字段
func (l *logger) WithField(key string, value interface{}) *logger {
	l.extends[key] = value
	return l
}

// 获取日志字段
func (l *logger) getFields() []zap.Field {
	buf, err := json.Marshal(l.extends)
	if err != nil {
		return l.fields
	}
	l.fields = append(l.fields, zap.String("extend", string(buf)))
	return l.fields
}

func (l *logger) Debug(msg string) {
	dbLogger.Debug(msg, l.getFields()...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	dbLogger.Debug(fmt.Sprintf(template, args...), l.getFields()...)
}

func (l *logger) Info(msg string) {
	dbLogger.Info(msg, l.getFields()...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	dbLogger.Info(fmt.Sprintf(template, args...), l.getFields()...)
}

func (l *logger) Warn(msg string) {
	dbLogger.Warn(msg, l.getFields()...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	dbLogger.Warn(fmt.Sprintf(template, args...), l.getFields()...)
}

func (l *logger) Error(msg string) {
	dbLogger.Error(msg, l.getFields()...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	dbLogger.Error(fmt.Sprintf(template, args...), l.getFields()...)
}

func (l *logger) Panic(msg string) {
	dbLogger.Panic(msg, l.getFields()...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	dbLogger.Panic(fmt.Sprintf(template, args...), l.getFields()...)
}
