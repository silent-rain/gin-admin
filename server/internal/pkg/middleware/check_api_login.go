// Package middleware 检查 API 访问鉴权中间件
package middleware

import (
	"strings"

	apiAuthDAO "github.com/silent-rain/gin-admin/internal/app/api_auth/dao"
	permissionDAO "github.com/silent-rain/gin-admin/internal/app/permission/dao"
	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// CheckApiLogin 检查 API 令牌鉴权中间件
// 需要在检查登录中间件之前注册，防止二次验证
func CheckApiLogin() gin.HandlerFunc {
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

		// 从请求头中获取 API Token
		token := ctx.GetHeader(constant.ApiHttpToken)
		if token == "" {
			ctx.Next()
			return
		}
		// 从请求头中获取 API Token 口令
		passphrase := ctx.GetHeader(constant.ApiHttpTokenPassphrase)
		if passphrase == "" {
			log.New(ctx).WithCode(errcode.ApiHttpTokenPassphraseEmptyError).Errorf("")
			response.New(ctx).WithCode(errcode.ApiHttpTokenPassphraseEmptyError).Json()
			ctx.Abort()
			return
		}

		// 获取缓存信息
		tokenUri := token + ctx.Request.URL.Path
		user, err := apiAuthDAO.NewApiTokenLoginCacheDao().Get(tokenUri)
		if err == nil {
			core.Context(ctx).UserId = user.UserId
			core.Context(ctx).Nickname = user.Nickname
			// 存在 API 令牌的情况下，不再验证用户密码
			core.Context(ctx).DisableCheckLogin = true
			ctx.Next()
			return
		}

		ct := chechkApiToken{}
		// 令牌口令信息验证
		if err := ct.checkApiToken(ctx, token, passphrase); err != nil {
			response.New(ctx).WithError(err).Json()
			ctx.Abort()
			return
		}
		// 设置用户信息到上下文
		if err := ct.setUserInfo(ctx, token); err != nil {
			response.New(ctx).WithError(err).Json()
			ctx.Abort()
			return
		}
		// 访问权限验证
		if err := ct.checkApiUri(ctx, token); err != nil {
			response.New(ctx).WithError(err).Json()
			ctx.Abort()
			return
		}
		// 设置 API Token 访问权限缓存
		if err := ct.SetCache(ctx, token); err != nil {
			response.New(ctx).WithError(err).Json()
			ctx.Abort()
			return
		}

		// 存在 API 令牌的情况下，不再验证用户密码
		core.Context(ctx).DisableCheckLogin = true

		ctx.Next()
	}
}

// API 令牌校验
type chechkApiToken struct{}

// 密匙口令验证
func (c chechkApiToken) checkApiToken(ctx *gin.Context, token, passphrase string) error {
	tokenObj, ok, err := permissionDAO.NewUserApiTokenDao().Info(token)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("")
		return errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Errorf("API 令牌不存在")
		return errcode.DBQueryEmptyError.WithMsg("API 令牌不存在")
	}
	if tokenObj.Passphrase != passphrase {
		log.New(ctx).WithCode(errcode.ApiHttpTokenPassphraseError).Errorf("")
		return errcode.ApiHttpTokenPassphraseError
	}
	if !strings.Contains(tokenObj.Permission, ctx.Request.Method) {
		log.New(ctx).WithCode(errcode.ApiHttpTokenMethodPermissionError).Errorf("")
		return errcode.ApiHttpTokenMethodPermissionError
	}
	return nil
}

// 设置用户信息到上下文
func (c chechkApiToken) setUserInfo(ctx *gin.Context, token string) error {
	user, ok, err := permissionDAO.NewUserDao().InfoByApiToken(token)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("")
		return errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.ApiHttpTokenInvalidError).Errorf("")
		return errcode.ApiHttpTokenInvalidError
	}
	core.Context(ctx).UserId = user.ID
	core.Context(ctx).Nickname = user.Nickname
	return nil
}

// 访问 URI 资源权限验证
func (c chechkApiToken) checkApiUri(ctx *gin.Context, token string) error {
	apiInfo, ok, err := apiAuthDAO.NewApiHttpDao().GetUriListByToken(token, ctx.Request.RequestURI)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("")
		return errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Errorf("没有该资源访问权限")
		return errcode.DBQueryEmptyError.WithMsg("没有该资源访问权限")
	}
	if apiInfo.Method != ctx.Request.Method {
		log.New(ctx).WithCode(errcode.ApiHttpTokenMethodPermissionError).Errorf("")
		return errcode.ApiHttpTokenMethodPermissionError
	}
	return nil
}

// 设置 API Token 访问权限缓存
func (c chechkApiToken) SetCache(ctx *gin.Context, token string) error {
	tokenUri := token + ctx.Request.URL.Path
	userId := core.Context(ctx).UserId
	Nickname := core.Context(ctx).Nickname
	if err := apiAuthDAO.NewApiTokenLoginCacheDao().Set(tokenUri, userId, Nickname); err == nil {
		log.New(ctx).WithCode(errcode.RedisSetKeyError).Errorf("%#v", err)
		return nil
	}
	return nil
}
