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
	Login(ctx *gin.Context, req systemDTO.UserLoginReq)
	Logout(ctx *gin.Context)
	Captcha(ctx *gin.Context)
	CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string)
	Captcha2(ctx *gin.Context)
	Captcha2Verify(ctx *gin.Context, captchaId string, verifyValue string)
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
func (h *userLoginService) Login(ctx *gin.Context, req systemDTO.UserLoginReq) {
	if !chechkCaptcha(ctx, req.CaptchaId, req.Captcha) {
		return
	}

	// 查询用户
	user, ok, err := h.dao.GetUsername(req.Username, req.Password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.DBQueryEmptyError).Error("用户名或者密码不正确")
		response.New(ctx).WithCode(statuscode.DBQueryEmptyError).WithMsg("用户名或者密码不正确").Json()
		return
	}
	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(statuscode.UserDisableError).Error("")
		response.New(ctx).WithCode(statuscode.UserDisableError).Json()
		return
	}

	// 生成 Token
	token, err := jwtToken.GenerateToken(user.ID, user.Phone, user.Email, user.Password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.TokenGenerateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.TokenGenerateError).Json()
		return
	}

	// 返回 Token
	result := systemDTO.UserLoginRsp{
		Token: token,
	}
	response.New(ctx).WithMsg("登录成功").WithData(result).Json()
}

// Logout 注销系统
func (h *userLoginService) Logout(ctx *gin.Context) {
	result := systemDTO.UserLoginRsp{}
	response.New(ctx).WithMsg("注销成功").WithData(result).Json()
}

// Captcha 验证码
func (h *userLoginService) Captcha(ctx *gin.Context) {
	captchaId, b64s, err := utils.NewCaptcha().MekeCaptcha(conf.CaptchaType)
	if err != nil {
		log.New(ctx).WithCode(statuscode.CaptchaGenerateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.CaptchaGenerateError).Json()
		return
	}

	result := map[string]string{
		"captcha_id": captchaId,
		"b64s":       b64s,
	}
	response.New(ctx).WithMsg("登录成功").WithData(result).Json()
}

// CaptchaVerify 验证码验证
func (h *userLoginService) CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) {
	// 校验验证码
	// 注意 Verify(id, VerifyValue, true) 中的 true参数
	// 当为 true 时，校验 传入的id 的验证码，校验完 这个ID的验证码就要在内存中删除
	// 当为 false 时，校验 传入的id 的验证码，校验完 这个ID的验证码不删除
	if !utils.CaptchaStore.Verify(captchaId, verifyValue, true) {
		log.New(ctx).WithCode(statuscode.CaptchaVerifyError).Error("")
		response.New(ctx).WithCode(statuscode.CaptchaVerifyError).Json()
		return
	}
	response.New(ctx).WithMsg("验证成功").Json()
}

// Captcha2 验证码
func (h *userLoginService) Captcha2(ctx *gin.Context) {
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
		response.New(ctx).WithCode(statuscode.CaptchaEtxNotFoundError).Json()
		return
	}

	download := false
	if download {
		ctx.Header("Content-Type", "application/octet-stream")
	}

	session := sessions.Default(ctx)
	session.Set("captcha_id", captchaId)
	_ = session.Save()

	ctx.Writer.Write(content.Bytes())
}

// Captcha2Verify 验证码验证
func (h *userLoginService) Captcha2Verify(ctx *gin.Context, captchaId string, verifyValue string) {
	if !captcha.VerifyString(captchaId, verifyValue) {
		log.New(ctx).WithCode(statuscode.CaptchaVerifyError).Error("")
		response.New(ctx).WithCode(statuscode.CaptchaVerifyError).Json()
		return
	}
	response.New(ctx).WithMsg("验证成功").Json()
}

// 检查验证码
func chechkCaptcha(ctx *gin.Context, captchaId, captcha string) bool {
	if captcha == "" {
		log.New(ctx).WithCode(statuscode.SessionGetCaptchaEmptyError).Error("")
		response.New(ctx).WithCode(statuscode.SessionGetCaptchaEmptyError).Json()
		return false
	}
	if captchaId == "" {
		log.New(ctx).WithCode(statuscode.CaptchaNotFoundError).Error("")
		response.New(ctx).WithCode(statuscode.CaptchaNotFoundError).Json()
		return false
	}

	// 校验验证码
	if !utils.CaptchaStore.Verify(captchaId, captcha, true) {
		log.New(ctx).WithCode(statuscode.CaptchaVerifyError).Error("")
		response.New(ctx).WithCode(statuscode.CaptchaVerifyError).Json()
		return false
	}
	return true
}
