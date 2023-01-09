/*
 * @Author: silent-rain
 * @Date: 2023-01-08 21:43:52
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-09 23:03:31
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/middleware/checklogin.go
 * @Descripttion: 登录验证中间件
 */
package middleware

import (
	"errors"
	"strings"

	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 放行白名单
var whiteList = []string{
	"/api/v1/register",
	"/api/v1/captcha",
	"/api/v1/captcha/verify",
	"/api/v1/login",
}

// CheckLogin 登录验证中间件
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if utils.IndexOfArray(whiteList, ctx.Request.URL.Path) != -1 {
			ctx.Next()
			return
		}

		// 从请求头中获取Token
		token := ctx.GetHeader("Authorization")
		// 字符串替换
		token = strings.Replace(token, "Bearer ", "", 1)
		// Token 解析
		claim, err := utils.ParseToken(token)
		if err != nil {
			parseTokenErr(ctx, err)
			zap.S().Errorf("token 解析失败, err: %v", err)
			ctx.Abort()
			return
		}
		ctx.Set("token", *claim)
		ctx.Next()
	}
}

// Token 解析异常处理
func parseTokenErr(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	if errors.Is(err, statuscode.TokenParsingError.Error()) {
		response.New(ctx).WithCode(statuscode.TokenParsingError).Json()
	} else if errors.Is(err, statuscode.TokeConvertError.Error()) {
		response.New(ctx).WithCode(statuscode.TokeConvertError).Json()
	} else if errors.Is(err, statuscode.TokenInvalidError.Error()) {
		response.New(ctx).WithCode(statuscode.TokenInvalidError).Json()
	} else if errors.Is(err, statuscode.TokenExpiredError.Error()) {
		response.New(ctx).WithCode(statuscode.TokenExpiredError).Json()
	}
}
