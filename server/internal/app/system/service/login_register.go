// Package service 用户登录/登出/注册
package service

import (
	permissionDTO "github.com/silent-rain/gin-admin/internal/app/permission/dto"
	permissionService "github.com/silent-rain/gin-admin/internal/app/permission/service"
	"github.com/silent-rain/gin-admin/internal/app/system/cache"
	"github.com/silent-rain/gin-admin/internal/app/system/dao"
	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"
	"github.com/silent-rain/gin-admin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// UserLoginRegisterService 用户登录/登出/注册
type UserLoginRegisterService struct {
	dao   *dao.Login
	cache cache.UserLoginCache
}

// NewUserLoginRegisterService 创建用户登录/登出/注册 对象
func NewUserLoginRegisterService() *UserLoginRegisterService {
	return &UserLoginRegisterService{
		dao:   dao.NewLoginDao(),
		cache: cache.NewUserLoginCache(),
	}
}

// Login 登录
func (h *UserLoginRegisterService) Login(ctx *gin.Context, req dto.UserLoginReq) (dto.UserLogin, error) {
	// 返回 Token
	result := dto.UserLogin{}

	// 检查验证码
	if err := chechkCaptcha(ctx, req.CaptchaId, req.Captcha); err != nil {
		return result, err
	}

	// 查询登录用户信息
	user, ok, err := h.dao.Login(req.Username, req.Password)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return result, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Error("用户名或者密码不正确")
		return result, errcode.DBQueryEmptyError.WithMsg("用户名或者密码不正确")
	}

	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(errcode.UserDisableError).Error("")
		return result, errcode.UserDisableError
	}

	// 生成 Token
	token, err := jwt.GenerateToken(user.ID, user.Nickname)
	if err != nil {
		log.New(ctx).WithCode(errcode.TokenGenerateError).Errorf("%v", err)
		return result, errcode.TokenGenerateError
	}

	// 存储登录日志
	_, err = NewUserLoginService().Add(ctx, model.UserLogin{
		UserId:     user.ID,
		Nickname:   user.Nickname,
		RemoteAddr: ctx.ClientIP(),
		UserAgent:  ctx.Request.UserAgent(),
		Status:     1,
	})
	if err != nil {
		log.New(ctx).WithError(err).Errorf("%v", err)
		return result, err
	}

	// 存储缓存
	if err := h.cache.Set(user.ID, token); err != nil {
		return result, err
	}

	result.Token = token
	return result, nil
}

// Logout 注销系统
func (h *UserLoginRegisterService) Logout(ctx *gin.Context) (dto.UserLogin, error) {
	result := dto.UserLogin{}
	return result, nil
}

// Register 注册
func (h *UserLoginRegisterService) Register(ctx *gin.Context, req permissionDTO.AddUserReq) error {
	// 注册入口检查验证码
	if ctx.Request.URL.Path == "/api/v1/register" {
		if err := chechkCaptcha(ctx, req.CaptchaId, req.Captcha); err != nil {
			return err
		}
	}
	// 添加用户
	return permissionService.NewUserService().Add(ctx, req)
}
