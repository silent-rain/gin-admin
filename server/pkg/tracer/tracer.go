// Package tracer 日志链路跟踪
package tracer

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// GenerateTraceId 日志链路 ID
func GenerateTraceId(ctx *gin.Context) string {
	rand.Seed(time.Now().UnixNano())
	data := time.Now().UTC().GoString() + ctx.Request.URL.Path + ctx.ClientIP() + ctx.Request.UserAgent()
	m := md5.New()
	m.Write([]byte(constant.Secret))
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

// span 日志
type logSpan struct {
	StartTime time.Time    `json:"start_time"` // span 记录开始时间
	EndTime   time.Time    `json:"end_time"`   // span 记录结束时间
	SpanId    string       `json:"span_id"`    // spanId 日志 ID
	Cost      int          `json:"cost"`       // 耗时,纳秒
	ctx       *gin.Context `json:"-"`          // 上下文
}

// SpanStart 开始记录 span
func SpanStart(ctx *gin.Context) *logSpan {
	return &logSpan{
		StartTime: time.Now().Local(),
		ctx:       ctx,
	}
}

// 生成 SpanId
func (s *logSpan) generateSpanId() string {
	rand.Seed(time.Now().UnixNano())
	data := time.Now().UTC().GoString() + s.StartTime.GoString() + s.EndTime.GoString() + strconv.Itoa(int(s.Cost))
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

// Finish span 完成
func (s *logSpan) Finish() *logSpan {
	s.EndTime = time.Now().Local()
	s.Cost = s.EndTime.Nanosecond() - s.StartTime.Nanosecond()
	s.SpanId = s.generateSpanId()

	log.New(s.ctx).WithCode(errcode.Ok).
		WithCallerSkip(1).
		WithSpanId(s.SpanId).
		WithField("data", *s).
		Info("span 日志埋点")
	return s
}

// GetSpanId 获取 SpanId
func (s *logSpan) GetSpanId() string {
	return s.SpanId
}
