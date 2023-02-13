/*用户登录/登出
 */
package service

import (
	"bytes"

	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/conf"
	jwtToken "gin-admin/internal/pkg/jwt_token"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 用户登录/登出
type UserLoginService interface {
	Login(ctx *gin.Context, req systemDTO.UserLoginReq) *response.ResponseAPI
	Logout(ctx *gin.Context) *response.ResponseAPI
	Captcha(ctx *gin.Context) *response.ResponseAPI
	CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) *response.ResponseAPI
	Captcha2(ctx *gin.Context) *response.ResponseAPI
	Captcha2Verify(ctx *gin.Context, captchaId string, verifyValue string) *response.ResponseAPI
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
func (h *userLoginService) Login(ctx *gin.Context, req systemDTO.UserLoginReq) *response.ResponseAPI {
	if result := chechkCaptcha(ctx, req.CaptchaId, req.Captcha); result.Error() != nil {
		return result
	}

	// 查询用户
	user, ok, err := h.dao.GetUsername(req.Username, req.Password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.DBQueryEmptyError).Error("用户名或者密码不正确")
		return response.New().WithCode(statuscode.DBQueryEmptyError).WithMsg("用户名或者密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(statuscode.UserDisableError).Error("")
		return response.New().WithCode(statuscode.UserDisableError)
	}

	// 生成 Token
	token, err := jwtToken.GenerateToken(user.ID, user.Phone, user.Email, user.Password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.TokenGenerateError).Errorf("%v", err)
		return response.New().WithCode(statuscode.TokenGenerateError)
	}

	// 返回 Token
	result := systemDTO.UserLoginRsp{
		Token: token,
	}
	return response.New().WithMsg("登录成功").WithData(result)
}

// Logout 注销系统
func (h *userLoginService) Logout(ctx *gin.Context) *response.ResponseAPI {
	result := systemDTO.UserLoginRsp{}
	return response.New().WithMsg("注销成功").WithData(result)
}

// Captcha 验证码
func (h *userLoginService) Captcha(ctx *gin.Context) *response.ResponseAPI {
	captchaId, b64s, err := utils.NewCaptcha().MekeCaptcha(conf.CaptchaType)
	if err != nil {
		log.New(ctx).WithCode(statuscode.CaptchaGenerateError).Errorf("%v", err)
		return response.New().WithCode(statuscode.CaptchaGenerateError)
	}

	result := map[string]string{
		"captcha_id": captchaId,
		"b64s":       b64s,
	}
	return response.New().WithMsg("登录成功").WithData(result)
}

// CaptchaVerify 验证码验证
func (h *userLoginService) CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) *response.ResponseAPI {
	// 校验验证码
	// 注意 Verify(id, VerifyValue, true) 中的 true参数
	// 当为 true 时，校验 传入的id 的验证码，校验完 这个ID的验证码就要在内存中删除
	// 当为 false 时，校验 传入的id 的验证码，校验完 这个ID的验证码不删除
	if !utils.CaptchaStore.Verify(captchaId, verifyValue, true) {
		log.New(ctx).WithCode(statuscode.CaptchaVerifyError).Error("")
		return response.New().WithCode(statuscode.CaptchaVerifyError)
	}
	return response.New().WithMsg("验证成功")
}

// Captcha2 验证码
func (h *userLoginService) Captcha2(ctx *gin.Context) *response.ResponseAPI {
	captchaId := captcha.NewLen(5)

	var content bytes.Buffer
	ext := ".png"
	switch ext {
	case ".png":
		ctx.Header("Content-Type", "image/png")
		captcha.WriteImage(&content, captchaId, captcha.StdWidth, captcha.StdHeight)
	case ".wav":
		ctx.Header("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, captchaId, "zh")
	default:
		return response.New().WithCode(statuscode.CaptchaEtxNotFoundError)
	}

	download := false
	if download {
		ctx.Header("Content-Type", "application/octet-stream")
	}

	session := sessions.Default(ctx)
	session.Set("captcha_id", captchaId)
	_ = session.Save()

	return response.New().WithData(content.Bytes())
}

// Captcha2Verify 验证码验证
func (h *userLoginService) Captcha2Verify(ctx *gin.Context, captchaId string, verifyValue string) *response.ResponseAPI {
	if !captcha.VerifyString(captchaId, verifyValue) {
		log.New(ctx).WithCode(statuscode.CaptchaVerifyError).Error("")
		return response.New().WithCode(statuscode.CaptchaVerifyError)

	}
	return response.New().WithMsg("验证成功")
}

// 检查验证码
func chechkCaptcha(ctx *gin.Context, captchaId, captcha string) *response.ResponseAPI {
	if captcha == "" {
		log.New(ctx).WithCode(statuscode.SessionGetCaptchaEmptyError).Error("")
		return response.New().WithCode(statuscode.SessionGetCaptchaEmptyError)
	}
	if captchaId == "" {
		log.New(ctx).WithCode(statuscode.CaptchaNotFoundError).Error("")
		return response.New().WithCode(statuscode.CaptchaNotFoundError)
	}

	// 校验验证码
	if !utils.CaptchaStore.Verify(captchaId, captcha, true) {
		log.New(ctx).WithCode(statuscode.CaptchaVerifyError).Error("")
		return response.New().WithCode(statuscode.CaptchaVerifyError)
	}
	return response.New()
}
