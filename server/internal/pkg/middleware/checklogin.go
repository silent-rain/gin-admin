/*登录验证中间件
 */
package middleware

import (
	"strings"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/core"
	"gin-admin/internal/pkg/jwt_token"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// CheckLogin 登录验证中间件
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// OPTIONS 过滤
		if ctx.Request.Method == "OPTIONS" {
			ctx.Next()
			return
		}
		// 请求是否禁用登录检查
		if core.GetContext(ctx).DisableCheckLogin {
			ctx.Next()
			return
		}

		// 从请求头中获取Token
		token := ctx.GetHeader(conf.TokenHeader)
		if token == "" {
			log.New(ctx).WithCode(errcode.TokenNotFound).Errorf("")
			response.New(ctx).WithCode(errcode.TokenNotFound).Json()
			ctx.Abort()
			return
		}
		// 字符串替换
		token = strings.Replace(token, "Bearer ", "", 1)
		// Token 解析
		claim, err := jwt_token.ParseToken(token)
		if err != nil {
			log.New(ctx).WithCodeError(err).Errorf("")
			response.New(ctx).WithCodeError(err).Json()
			ctx.Abort()
			return
		}
		core.GetContext(ctx).UserId = claim.UserId
		ctx.Next()
	}
}