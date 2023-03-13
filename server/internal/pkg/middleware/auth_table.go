/*鉴权表*/
package middleware

import (
	"regexp"

	"github.com/silent-rain/gin-admin/internal/pkg/core"

	"github.com/gin-gonic/gin"
)

var (
	// 请求放行白名单
	apiReleaseWhiteList = []string{
		// 注册/验证码/登录
		"^/api/v1/register",
		"^/api/v1/captcha",
		"^/api/v1/captcha/verify",
		"^/api/v1/login",
		"^/api/v1/config/webSiteConfigList",

		// WEB
		"^/favicon.ico$",
		"^/$",
		"^/static/.*",
		"^/upload/.*",

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
// 需要登录验证中间件之前注册，防止过白名单失败
func AuthTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		// API 登录检查放行
		for _, item := range apiReleaseWhiteList {
			reg := regexp.MustCompile(item)
			result := reg.FindAllString(path, -1)
			if len(result) > 0 {
				core.GetContext(ctx).DisableCheckLogin = true
				break
			}
		}

		// 接口禁止限流
		for _, item := range rateLimiterList {
			reg := regexp.MustCompile(item)
			result := reg.FindAllString(path, -1)
			if len(result) > 0 {
				core.GetContext(ctx).DisableCheckLogin = true
				break
			}
		}
		ctx.Next()
	}
}
