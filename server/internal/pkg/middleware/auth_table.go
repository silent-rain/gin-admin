/*鉴权表*/
package middleware

import (
	"regexp"

	"gin-admin/internal/pkg/core"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	// 请求放行白名单
	apiReleaseWhiteList = []string{
		// 注册/验证码/登录
		"^/api/v1/register",
		"^/api/v1/captcha",
		"^/api/v1/captcha/verify",
		"^/api/v1/login",

		// WEB
		"^/favicon.ico$",
		"^/$",
		"^/static/.*",

		// pprof 性能剖析工具
		"^/debug/pprof",
	}
	// 接口禁止限流白名单
	rateLimiterList = []string{
		"^/api/ping",
		"^/api/health",
	}
)

// AuthTable 鉴权表
func AuthTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		// API 登录检查放行
		for _, item := range apiReleaseWhiteList {
			reg := regexp.MustCompile(item)
			result := reg.FindAllString(path, -1)
			if len(result) > 0 {
				zap.S().Error(result)
				core.GetContext(ctx).DisableCheckLogin = true
			}
		}

		// 接口禁止限流
		for _, item := range rateLimiterList {
			reg := regexp.MustCompile(item)
			result := reg.FindAllString(path, -1)
			if len(result) > 0 {
				zap.S().Error(result)
				core.GetContext(ctx).DisableCheckLogin = true
			}
		}
		ctx.Next()
	}
}
