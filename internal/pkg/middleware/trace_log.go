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
		// 依赖与外部入 traceTd 参存在风险
		// var traceId = ctx.Request.Header.Get(context.GinContextTraceTd)

		var traceId = context.GenerateTraceId(ctx)
		// 请求
		ctx.Request.Header.Set(context.GinContextTraceTd, traceId)
		// 响应
		ctx.Header(context.GinContextTraceTd, traceId)
		ctx.Next()
	}
}
