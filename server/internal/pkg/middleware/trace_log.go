/*trace 日志链路跟踪中间件
 */
package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/constant"
	"gin-admin/internal/pkg/core"

	"github.com/gin-gonic/gin"
)

// TraceLogger 日志链路跟踪中间件
func TraceLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var traceId = ctx.Request.Header.Get(constant.HeaderTraceTd)
		var spanId = generateSpanId(ctx)
		// 设置请求头
		ctx.Header(constant.HeaderTraceTd, spanId)
		// 设置上下文
		core.GetContext(ctx).TraceId = traceId
		core.GetContext(ctx).SpanId = spanId
		ctx.Next()
	}
}

// 生成 spanId
func generateSpanId(ctx *gin.Context) string {
	rand.Seed(time.Now().UnixNano())
	data := time.Now().UTC().GoString() + ctx.Request.URL.Path + ctx.ClientIP() + ctx.Request.UserAgent()
	m := md5.New()
	m.Write([]byte(conf.Secret))
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}
