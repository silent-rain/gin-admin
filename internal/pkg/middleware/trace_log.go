/*trace 日志链路跟踪中间件
 */
package middleware

import (
	"gin-admin/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

// TraceLogger 日志链路跟踪中间件
func TraceLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var traceId = ctx.Request.Header.Get(context.GinContextTraceTd)

		var spanId = context.GenerateTraceId(ctx)
		// 请求
		ctx.Request.Header.Set(context.GinContextTraceTd, traceId)
		ctx.Request.Header.Set(context.GinContextSpanId, spanId)
		// 响应
		ctx.Header(context.GinContextTraceTd, spanId)
		ctx.Next()
	}
}
