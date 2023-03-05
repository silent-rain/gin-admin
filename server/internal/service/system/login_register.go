/*用户登录/登出/注册
 */
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	permissionDTO "gin-admin/internal/dto/permission"
	systemDTO "gin-admin/internal/dto/system"
	jwt "gin-admin/internal/pkg/jwt"
	"gin-admin/internal/pkg/log"
	permissionService "gin-admin/internal/service/permission"
	systemVO "gin-admin/internal/vo/system"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// UserLoginRegisterService 用户登录/登出/注册接口
type UserLoginRegisterService interface {
	Login(ctx *gin.Context, req systemDTO.UserLoginReq) (systemVO.UserLogin, error)
	Logout(ctx *gin.Context) (systemVO.UserLogin, error)
	Register(ctx *gin.Context, req permissionDTO.AddUserReq) error
}

// 用户登录/登出/注册
type userLoginRegisterService struct {
	dao systemDAO.Login
}

// NewUserLoginRegisterService 创建用户登录/登出/注册 对象
func NewUserLoginRegisterService() *userLoginRegisterService {
	return &userLoginRegisterService{
		dao: systemDAO.NewLoginDao(),
	}
}

// Login 登录
func (h *userLoginRegisterService) Login(ctx *gin.Context, req systemDTO.UserLoginReq) (systemVO.UserLogin, error) {
	// 返回 Token
	result := systemVO.UserLogin{}

	// 检查验证码
	if err := chechkCaptcha(ctx, req.CaptchaId, req.Captcha); err != nil {
		return result, err
	}

	// 查询登录用户信息
	user, ok, err := h.dao.Login(req.Username, req.Password)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return result, errcode.New(errcode.DBQueryError)
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Error("用户名或者密码不正确")
		return result, errcode.New(errcode.DBQueryEmptyError).WithMsg("用户名或者密码不正确")
	}

	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(errcode.UserDisableError).Error("")
		return result, errcode.New(errcode.UserDisableError)
	}

	// 生成 Token
	token, err := jwt.GenerateToken(user.ID, user.Nickname)
	if err != nil {
		log.New(ctx).WithCode(errcode.TokenGenerateError).Errorf("%v", err)
		return result, errcode.New(errcode.TokenGenerateError)
	}
	result.Token = token
	return result, nil
}

// Logout 注销系统
func (h *userLoginRegisterService) Logout(ctx *gin.Context) (systemVO.UserLogin, error) {
	result := systemVO.UserLogin{}
	return result, nil
}

// Register 注册
func (h *userLoginRegisterService) Register(ctx *gin.Context, req permissionDTO.AddUserReq) error {
	// 注册入口检查验证码
	if ctx.Request.URL.Path == "/api/v1/register" {
		if err := chechkCaptcha(ctx, req.CaptchaId, req.Captcha); err != nil {
			return err
		}
	}
	// 添加用户
	return permissionService.NewUserService().Add(ctx, req)
}
