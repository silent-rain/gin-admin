/*跨域
 */
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cros 处理跨域请求, 支持options访问
func Cros() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")
		if len(origin) == 0 {
			ctx.Next()
			return
		}

		// 同源直接过
		host := ctx.GetHeader("Host")
		if origin == "http://"+host || origin == "https://"+host {
			ctx.Next()
			return
		}

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// OPTIONS 过
		method := ctx.Request.Method
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			ctx.Abort()
		}
		ctx.Next()
	}
}
