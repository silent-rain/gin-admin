// Package log 数据库日志
package log

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志结构
type logger struct {
	ctx        *gin.Context
	zapLog     *zap.Logger
	fields     []zapcore.Field
	extends    map[string]interface{} // 消息扩展字段
	callerSkip int
}

// New 创建日志对象
func New(ctx *gin.Context) *logger {
	extCtx := core.Context(ctx)

	traceId := extCtx.TraceId
	userId := extCtx.UserId
	nickname := extCtx.Nickname

	fields := []zapcore.Field{
		zap.String("trace_id", traceId),
		zap.Uint("user_id", userId),
		zap.String("nickname", nickname),
	}
	return &logger{
		ctx:        ctx,
		zapLog:     zap.L().WithOptions(zap.AddCallerSkip(1)),
		fields:     fields,
		extends:    make(map[string]interface{}, 0),
		callerSkip: 1,
	}
}

// WithCode 添加错误码
func (l *logger) WithCode(code errcode.ErrorCode) *logger {
	l.fields = append(
		l.fields,
		zap.Uint("error_code", uint(code)),
		zap.String("error_msg", code.Error()),
	)
	return l
}

// WithCodeError 添加响应状态码及状态码对应的信息
func (l *logger) WithError(err error) *logger {
	// 业务错误码 error code
	if code, ok := err.(errcode.ErrorCode); ok {
		l.fields = append(l.fields, zap.Uint("error_code", uint(code)), zap.String("error_msg", code.Error()))
		return l
	}

	// 业务错误码附加信息 erro code
	if msg, ok := err.(*errcode.ErrorMsg); ok {
		l.fields = append(l.fields, zap.Uint("error_code", uint(msg.Code)), zap.String("error_msg", msg.Error()))
		return l
	}

	// 原始错误
	l.fields = append(l.fields, zap.Uint("error_code", uint(errcode.UnknownError)),
		zap.String("error_msg", err.Error()))
	return l
}

// WithField 添加扩展字段
func (l *logger) WithField(key string, value interface{}) *logger {
	l.extends[key] = value
	return l
}

// WithStack 添加堆栈信息
func (l *logger) WithStack() *logger {
	l.fields = append(l.fields, zap.String("stack", string(debug.Stack())))
	return l
}

// WithSpanId 添加日志链路 spanID
func (l *logger) WithSpanId(value string) *logger {
	l.fields = append(l.fields, zap.String("span_id", value))
	return l
}

// WithCallerSkip 调整日志位置
func (l *logger) WithCallerSkip(skip int) *logger {
	l.callerSkip = skip
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
	dbLogger.WithOptions(zap.AddCallerSkip(l.callerSkip)).Debug(msg, l.getFields()...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	dbLogger.Debug(fmt.Sprintf(template, args...), l.getFields()...)
}

func (l *logger) Info(msg string) {
	dbLogger.WithOptions(zap.AddCallerSkip(l.callerSkip)).Info(msg, l.getFields()...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	dbLogger.WithOptions(zap.AddCallerSkip(l.callerSkip)).Info(fmt.Sprintf(template, args...), l.getFields()...)
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
