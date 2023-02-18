/*上下文信息*/
package context

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"gin-admin/internal/pkg/conf"
	jwtToken "gin-admin/internal/pkg/jwt_token"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	// GinContextToken set/get Token 信息
	GinContextToken = "token"
	// GinContextToken set/get trace_id 信息
	GinContextTraceTd = "trace_id"
	// GinContextToken set/get span_id 信息
	GinContextSpanId = "span_id"
)

// GetUserId 获取用户 ID
func GetUserId(ctx *gin.Context) uint {
	v, ok := ctx.Get(GinContextToken)
	if !ok {
		return 0
	}
	token := v.(jwtToken.Token)
	return uint(token.UserId)
}

// GetTraceId 获取请求 TraceTd
func GetTraceId(ctx *gin.Context) string {
	var traceTd = ctx.Request.Header.Get(GinContextTraceTd)
	return traceTd
}

// GetSpanId 获取请求 SpanId
func GetSpanId(ctx *gin.Context) string {
	var spanId = ctx.Request.Header.Get(GinContextSpanId)
	if spanId == "" {
		spanId = GenerateTraceId(ctx)
	}
	return spanId
}

// GenerateTraceId 生成 traceId
func GenerateTraceId(ctx *gin.Context) string {
	data := fmt.Sprintf(`unix_nano: %v
					rand: %v,
					status: %v,
					method: %v,
					path: %v,
					query: %v,
					remote_addr: %v,
					user_agent: %v`,
		time.Now().UnixNano(),
		rand.Int63n(5000000),
		ctx.Writer.Status(),
		ctx.Request.URL.Path,
		ctx.Request.URL.Path,
		ctx.Request.URL.RawQuery,
		ctx.ClientIP(),
		ctx.Request.UserAgent(),
	)
	m := md5.New()
	m.Write([]byte(conf.Secret))
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}
