/*扩展 gin context*/
package core

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var contextKey = "__context__"

// 扩展 Context
type etxContext struct {
	// 用户信息
	UserId   uint   // 用户 ID
	Nickname string // 用户昵称

	// 日志
	TraceId string
	Span    *logSpan

	// 中间件处理
	DisableCheckLogin  bool // 禁用登录检查
	DisableRateLimiter bool // 禁用接口限流
}

// GetContext 获取扩展 Context
func GetContext(ctx *gin.Context) *etxContext {
	newC := &etxContext{
		DisableCheckLogin:  false,
		DisableRateLimiter: false,
		Span:               &logSpan{},
	}
	if ctx == nil {
		return newC
	}
	c, ok := ctx.Get(contextKey)
	if !ok {
		ctx.Set(contextKey, newC)
		return newC
	}
	return c.(*etxContext)
}

// span 日志
type logSpan struct {
	startTime time.Time // span 记录开始时间
	endTime   time.Time // span 记录结束时间
	spanId    string    // spanId 日志 ID
	cost      int64     // 耗时,纳秒
}

// Start 开始记录 span
func (s *logSpan) Start() *logSpan {
	s.startTime = time.Now()
	return s
}

// 生成 SpanId
func (s *logSpan) generateSpanId() string {
	rand.Seed(time.Now().UnixNano())
	data := time.Now().UTC().GoString() + s.startTime.GoString() + s.endTime.GoString() + strconv.Itoa(int(s.cost))
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

// Finish span 完成
func (s *logSpan) Finish() *logSpan {
	s.endTime = time.Now()
	s.spanId = s.generateSpanId()
	return s
}

// SpanId 获取 SpanId
func (s *logSpan) SpanId() string {
	return s.spanId
}
