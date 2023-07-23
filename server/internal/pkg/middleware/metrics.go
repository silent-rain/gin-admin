// Package middleware 指标记录
package middleware

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/silent-rain/gin-admin/global"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/pkg/constant"
	"github.com/silent-rain/gin-admin/pkg/metrics"
	"github.com/silent-rain/gin-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// Metrics 指标记录
func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !global.Instance().Config().Server.Plugin.EnableRecordMetrics {
			ctx.Next()
			return
		}

		start := time.Now()
		// record response info
		blw := &ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		ctx.Next()

		// 判断是否为接口, 当为接口时记录返回信息
		result := &response.ResponseAPI{}
		if err := json.Unmarshal(blw.Body.Bytes(), result); err != nil {
			return
		}

		// 记录指标
		metrics.RecordMetrics(model.MetricsMessage{
			ProjectName: constant.ProjectName,
			Env:         global.Instance().Config().Environment.Env,
			TraceID:     core.Context(ctx).TraceId,
			HOST:        ctx.Request.Host,
			Path:        ctx.Request.URL.Path,
			Method:      ctx.Request.Method,
			HttpStatus:  ctx.Writer.Status(),
			ErrorCode:   uint(result.Code),
			CostSeconds: float64(time.Since(start).Nanoseconds()),
		})
	}
}
