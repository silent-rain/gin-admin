/*限速器*/
package middleware

import (
	"gin-admin/internal/pkg/conf"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter 限速器中间件
func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
