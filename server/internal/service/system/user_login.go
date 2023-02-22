/*用户登录/登出
 */
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/constant"
	jwtToken "gin-admin/internal/pkg/jwt_token"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/utils"
	systemVO "gin-admin/internal/vo/system"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// UserLoginService 用户登录/登出
type UserLoginService interface {
	Login(ctx *gin.Context, req systemDTO.UserLoginReq) (systemVO.UserLogin, error)
	Logout(ctx *gin.Context) (systemVO.UserLogin, error)
	Captcha(ctx *gin.Context) (systemVO.Captcha, error)
	CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) error
}

// 用户登录/登出
type userLoginService struct {
	dao systemDAO.User
}

// NewUserLoginService 创建用户登录/登出 对象
func NewUserLoginService() *userLoginService {
	return &userLoginService{
		dao: systemDAO.NewUserDao(),
	}
}

// Login 登录
func (h *userLoginService) Login(ctx *gin.Context, req systemDTO.UserLoginReq) (systemVO.UserLogin, error) {
	// 返回 Token
	result := systemVO.UserLogin{
		Token: "",
	}

	if err := chechkCaptcha(ctx, req.CaptchaId, req.Captcha); err != nil {
		return result, err
	}

	// 查询用户
	user, ok, err := h.dao.GetUsername(req.Username, req.Password)
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
	token, err := jwtToken.GenerateToken(user.ID, user.Nickname, user.Phone, user.Email, user.Password)
	if err != nil {
		log.New(ctx).WithCode(errcode.TokenGenerateError).Errorf("%v", err)
		return result, errcode.New(errcode.TokenGenerateError)
	}
	result.Token = token
	return result, nil
}

// Logout 注销系统
func (h *userLoginService) Logout(ctx *gin.Context) (systemVO.UserLogin, error) {
	result := systemVO.UserLogin{}
	return result, nil
}

// Captcha 验证码
func (h *userLoginService) Captcha(ctx *gin.Context) (systemVO.Captcha, error) {
	result := systemVO.Captcha{
		CaptchaId: "",
		B64s:      "",
	}

	captchaId, b64s, err := utils.NewCaptcha().MekeCaptcha(constant.CaptchaType)
	if err != nil {
		log.New(ctx).WithCode(errcode.CaptchaGenerateError).Errorf("%v", err)
		return result, errcode.New(errcode.CaptchaGenerateError)
	}

	result.CaptchaId = captchaId
	result.B64s = b64s
	return result, nil
}

// CaptchaVerify 验证码验证
func (h *userLoginService) CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) error {
	// 校验验证码
	// 注意 Verify(id, VerifyValue, true) 中的 true参数
	// 当为 true 时，校验 传入的id 的验证码，校验完 这个ID的验证码就要在内存中删除
	// 当为 false 时，校验 传入的id 的验证码，校验完 这个ID的验证码不删除
	if !utils.CaptchaStore.Verify(captchaId, verifyValue, true) {
		log.New(ctx).WithCode(errcode.CaptchaVerifyError).Error("")
		return errcode.New(errcode.CaptchaVerifyError)
	}
	return nil
}

// 检查验证码
func chechkCaptcha(ctx *gin.Context, captchaId, captcha string) error {
	if captcha == "" {
		log.New(ctx).WithCode(errcode.SessionGetCaptchaEmptyError).Error("")
		return errcode.New(errcode.SessionGetCaptchaEmptyError)
	}
	if captchaId == "" {
		log.New(ctx).WithCode(errcode.CaptchaNotFoundError).Error("")
		return errcode.New(errcode.CaptchaNotFoundError)
	}

	// 校验验证码
	if !utils.CaptchaStore.Verify(captchaId, captcha, true) {
		log.New(ctx).WithCode(errcode.CaptchaVerifyError).Error("")
		return errcode.New(errcode.CaptchaVerifyError)
	}
	return nil
}
