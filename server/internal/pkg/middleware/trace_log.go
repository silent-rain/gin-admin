/*trace 日志链路跟踪中间件
 */
package middleware

import (
	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/pkg/tracer"

	"github.com/gin-gonic/gin"
)

// TraceLogger 日志链路跟踪中间件
func TraceLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := ctx.Request.Header.Get(constant.HeaderTraceTd)
		if traceId == "" {
			traceId = tracer.GenerateTraceId(ctx)
		}
		// 设置上下文
		core.GetContext(ctx).TraceId = traceId
		ctx.Next()
	}
}
