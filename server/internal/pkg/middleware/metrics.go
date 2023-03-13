/*指标记录*/
package middleware

import (
	"bytes"
	"encoding/json"
	"time"

	systemModel "github.com/silent-rain/gin-admin/internal/model/system"
	"github.com/silent-rain/gin-admin/internal/pkg/conf"
	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	"github.com/silent-rain/gin-admin/pkg/metrics"

	"github.com/gin-gonic/gin"
)

// Metrics 指标记录
func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !conf.Instance().Server.Plugin.EnableRecordMetrics {
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
		metrics.RecordMetrics(systemModel.MetricsMessage{
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
