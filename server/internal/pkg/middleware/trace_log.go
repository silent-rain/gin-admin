/*trace 日志链路跟踪中间件
 */
package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"gin-admin/internal/pkg/constant"
	"gin-admin/internal/pkg/core"

	"github.com/gin-gonic/gin"
)

// TraceLogger 日志链路跟踪中间件
func TraceLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := ctx.Request.Header.Get(constant.HeaderTraceTd)
		if traceId == "" {
			traceId = generateTraceId(ctx)
		}
		// 设置上下文
		core.GetContext(ctx).TraceId = traceId
		ctx.Next()
	}
}

// 生成 traceId
func generateTraceId(ctx *gin.Context) string {
	rand.Seed(time.Now().UnixNano())
	data := time.Now().UTC().GoString() + ctx.Request.URL.Path + ctx.ClientIP() + ctx.Request.UserAgent()
	m := md5.New()
	m.Write([]byte(constant.Secret))
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}
