// Package middleware 登录验证中间件
package middleware

import (
	"strings"

	"github.com/silent-rain/gin-admin/global"
	"github.com/silent-rain/gin-admin/internal/app/system/cache"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"
	"github.com/silent-rain/gin-admin/pkg/jwt"
	"github.com/silent-rain/gin-admin/pkg/response"

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
		if core.Context(ctx).DisableCheckLogin {
			ctx.Next()
			return
		}
		cfgJWT := global.Instance().Config().JWT
		// 从请求头中获取Token
		token := ctx.GetHeader(cfgJWT.Header)
		if token == "" {
			log.New(ctx).WithCode(errcode.TokenNotFound).Errorf("")
			response.New(ctx).WithCode(errcode.TokenNotFound).Json()
			ctx.Abort()
			return
		}
		// 字符串替换
		token = strings.Replace(token, cfgJWT.Prefix, "", 1)
		// Token 解析
		claim, err := jwt.ParseToken(token)
		if err != nil {
			log.New(ctx).WithError(err).Errorf("")
			response.New(ctx).WithError(err).Json()
			ctx.Abort()
			return
		}
		core.Context(ctx).UserId = claim.UserId
		core.Context(ctx).Nickname = claim.Nickname

		// 检查单点登录
		if err := checkSingleLogin(claim.UserId, token); err != nil {
			log.New(ctx).WithError(err).Errorf("")
			response.New(ctx).WithError(err).Json()
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// 检查单点登录
func checkSingleLogin(userId uint, token string) error {
	if !global.Instance().Config().Server.Plugin.EnableSingleLogin {
		return nil
	}
	tk, err := cache.NewUserLoginCache().Get(userId)
	if err != nil {
		return err
	}
	if tk == "" {
		return errcode.TokenDisableCurrentLoginError
	}
	if tk != token {
		return errcode.TokenUnconformityError
	}
	return nil
}
