/*指标记录*/
package middleware

import (
	"bytes"
	"encoding/json"
	"time"

	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/constant"
	"gin-admin/internal/pkg/core"
	"gin-admin/internal/pkg/metrics"
	"gin-admin/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// Metrics 指标记录
func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !conf.Instance().Server.Base.EnableRecordMetrics {
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

		metrics.RecordHandler(systemModel.MetricsMessage{
			ProjectName: constant.ProjectName,
			Env:         conf.Instance().Environment.Env,
			TraceID:     core.GetContext(ctx).TraceId,
			HOST:        ctx.Request.Host,
			Path:        ctx.Request.URL.Path,
			Method:      ctx.Request.Method,
			HttpStatus:  ctx.Writer.Status(),
			ErrorCode:   uint(result.Code),
			CostSeconds: float64(time.Since(start).Nanoseconds()),
		})
	}
}
