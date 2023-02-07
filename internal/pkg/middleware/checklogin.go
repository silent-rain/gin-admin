/*
 * @Author: silent-rain
 * @Date: 2023-01-08 21:43:52
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-15 01:26:29
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/middleware/checklogin.go
 * @Descripttion: 登录验证中间件
 */
package middleware

import (
	"errors"
	"strings"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 放行白名单
var apiWhiteList = []string{
	// API
	"/api/ping",
	"/api/v1/sayHello",
	"/api/v1/register",
	"/api/v1/captcha",
	"/api/v1/captcha/verify",
	"/api/v1/login",
}

// CheckLogin 登录验证中间件
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// OPTIONS 过滤
		if ctx.Request.Method == "OPTIONS" {
			ctx.Next()
			return
		}
		// 请求白名单过滤
		if utils.IndexOfArray(apiWhiteList, ctx.Request.URL.Path) != -1 {
			ctx.Next()
			return
		}
		// 非 API 请求过滤
		if !strings.HasPrefix(ctx.Request.URL.Path, "/api") {
			ctx.Next()
			return
		}

		// 从请求头中获取Token
		token := ctx.GetHeader(conf.TokenHeader)
		if token == "" {
			log.New(ctx).WithCode(statuscode.TokenNotFound).Errorf("")
			response.New(ctx).WithCode(statuscode.TokenNotFound).Json()
			ctx.Abort()
			return
		}
		// 字符串替换
		token = strings.Replace(token, "Bearer ", "", 1)
		// Token 解析
		claim, err := utils.ParseToken(token)
		if err != nil {
			parseTokenErr(ctx, err)
			ctx.Abort()
			return
		}
		ctx.Set(utils.GinContextToken, *claim)
		ctx.Next()
	}
}

// Token 解析异常处理
func parseTokenErr(ctx *gin.Context, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, statuscode.TokenParsingError.Error()) {
		log.New(ctx).WithCode(statuscode.TokenParsingError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.TokenParsingError).Json()
	} else if errors.Is(err, statuscode.TokeConvertError.Error()) {
		log.New(ctx).WithCode(statuscode.TokeConvertError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.TokeConvertError).Json()
	} else if errors.Is(err, statuscode.TokenInvalidError.Error()) {
		log.New(ctx).WithCode(statuscode.TokenInvalidError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.TokenInvalidError).Json()
	} else if errors.Is(err, statuscode.TokenExpiredError.Error()) {
		log.New(ctx).WithCode(statuscode.TokenExpiredError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.TokenExpiredError).Json()
	}
	return err
}
