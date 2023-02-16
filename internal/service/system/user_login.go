/*用户登录/登出
 */
package system

import (
	"bytes"

	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/code_errors"
	"gin-admin/internal/pkg/conf"
	jwtToken "gin-admin/internal/pkg/jwt_token"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/utils"
	systemVO "gin-admin/internal/vo/system"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 用户登录/登出
type UserLoginService interface {
	Login(ctx *gin.Context, req systemDTO.UserLoginReq) (systemDTO.UserLoginRsp, error)
	Logout(ctx *gin.Context) (systemDTO.UserLoginRsp, error)
	Captcha(ctx *gin.Context) (systemVO.Captcha, error)
	CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) error
	Captcha2(ctx *gin.Context) ([]byte, error)
	Captcha2Verify(ctx *gin.Context, captchaId string, verifyValue string) error
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
func (h *userLoginService) Login(ctx *gin.Context, req systemDTO.UserLoginReq) (systemDTO.UserLoginRsp, error) {
	// 返回 Token
	result := systemDTO.UserLoginRsp{
		Token: "",
	}

	if err := chechkCaptcha(ctx, req.CaptchaId, req.Captcha); err != nil {
		return result, err
	}

	// 查询用户
	user, ok, err := h.dao.GetUsername(req.Username, req.Password)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBQueryError).Errorf("%v", err)
		return result, code_errors.New(code_errors.DBQueryError)
	}
	if !ok {
		log.New(ctx).WithCode(code_errors.DBQueryEmptyError).Error("用户名或者密码不正确")
		return result, code_errors.New(code_errors.DBQueryEmptyError).WithMsg("用户名或者密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(code_errors.UserDisableError).Error("")
		return result, code_errors.New(code_errors.UserDisableError)
	}

	// 生成 Token
	token, err := jwtToken.GenerateToken(user.ID, user.Phone, user.Email, user.Password)
	if err != nil {
		log.New(ctx).WithCode(code_errors.TokenGenerateError).Errorf("%v", err)
		return result, code_errors.New(code_errors.TokenGenerateError)
	}
	result.Token = token
	return result, nil
}

// Logout 注销系统
func (h *userLoginService) Logout(ctx *gin.Context) (systemDTO.UserLoginRsp, error) {
	result := systemDTO.UserLoginRsp{}
	return result, nil
}

// Captcha 验证码
func (h *userLoginService) Captcha(ctx *gin.Context) (systemVO.Captcha, error) {
	result := systemVO.Captcha{
		CaptchaId: "",
		B64s:      "",
	}

	captchaId, b64s, err := utils.NewCaptcha().MekeCaptcha(conf.CaptchaType)
	if err != nil {
		log.New(ctx).WithCode(code_errors.CaptchaGenerateError).Errorf("%v", err)
		return result, code_errors.New(code_errors.CaptchaGenerateError)
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
		log.New(ctx).WithCode(code_errors.CaptchaVerifyError).Error("")
		return code_errors.New(code_errors.CaptchaVerifyError)
	}
	return nil
}

// Captcha2 验证码
func (h *userLoginService) Captcha2(ctx *gin.Context) ([]byte, error) {
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
		log.New(ctx).WithCode(code_errors.CaptchaEtxNotFoundError).Error("")
		return nil, code_errors.New(code_errors.CaptchaEtxNotFoundError)
	}

	download := false
	if download {
		ctx.Header("Content-Type", "application/octet-stream")
	}

	session := sessions.Default(ctx)
	session.Set("captcha_id", captchaId)
	_ = session.Save()

	return content.Bytes(), nil
}

// Captcha2Verify 验证码验证
func (h *userLoginService) Captcha2Verify(ctx *gin.Context, captchaId string, verifyValue string) error {
	if !captcha.VerifyString(captchaId, verifyValue) {
		log.New(ctx).WithCode(code_errors.CaptchaVerifyError).Error("")
		return code_errors.New(code_errors.CaptchaVerifyError)

	}
	return nil
}

// 检查验证码
func chechkCaptcha(ctx *gin.Context, captchaId, captcha string) error {
	if captcha == "" {
		log.New(ctx).WithCode(code_errors.SessionGetCaptchaEmptyError).Error("")
		return code_errors.New(code_errors.SessionGetCaptchaEmptyError)
	}
	if captchaId == "" {
		log.New(ctx).WithCode(code_errors.CaptchaNotFoundError).Error("")
		return code_errors.New(code_errors.CaptchaNotFoundError)
	}

	// 校验验证码
	if !utils.CaptchaStore.Verify(captchaId, captcha, true) {
		log.New(ctx).WithCode(code_errors.CaptchaVerifyError).Error("")
		return code_errors.New(code_errors.CaptchaVerifyError)
	}
	return nil
}