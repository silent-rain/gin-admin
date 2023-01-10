/*
 * @Author: silent-rain
 * @Date: 2023-01-10 21:26:07
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-10 22:24:09
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/middleware/trace_log.go
 * @Descripttion: trace 日志链路跟踪中间件
 */
package middleware

import (
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// TraceLogger 日志链路跟踪中间件
func TraceLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 验证 API 的 Content-Type 是否为 json
		if err := utils.VerifyContentTypeJson(ctx); err != nil {
			return
		}

		// 依赖与外部入 traceTd 参存在风险
		// var traceId = ctx.Request.Header.Get(utils.GinContextTraceTd)

		var traceId = utils.GenerateTraceId(ctx)
		// 请求
		ctx.Request.Header.Set(utils.GinContextTraceTd, traceId)
		// 响应
		ctx.Header(utils.GinContextTraceTd, traceId)
		ctx.Next()
	}
}
