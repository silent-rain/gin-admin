/*限速器*/
package middleware

import (
	"time"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/core"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter 接口限流中间件
func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 请求禁用接口限流
		if core.GetContext(ctx).DisableCheckLogin {
			ctx.Next()
			return
		}

		cfg := conf.Instance().Server.Base
		if !cfg.EnableRateLimiter {
			ctx.Next()
			return
		}
		limiter := rate.NewLimiter(rate.Every(time.Second*1), cfg.MaxRequestsPerSecond)
		if !limiter.Allow() {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
