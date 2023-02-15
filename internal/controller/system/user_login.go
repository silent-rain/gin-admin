/*用户登录/登出
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/code_errors"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	service "gin-admin/internal/service/system"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 用户登录/登出
type userLoginController struct {
	service service.UserLoginService
}

// NewUserLoginController 创建用户登录/登出 对象
func NewUserLoginController() *userLoginController {
	return &userLoginController{
		service: service.NewUserLoginService(),
	}
}

// Login 登录
func (c *userLoginController) Login(ctx *gin.Context) {
	req := systemDTO.UserLoginReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	result, err := c.service.Login(ctx, req)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}

// Logout 注销系统
func (c *userLoginController) Logout(ctx *gin.Context) {
	if _, err := c.service.Logout(ctx); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Captcha 验证码
func (c *userLoginController) Captcha(ctx *gin.Context) {
	result, err := c.service.Captcha(ctx)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}

// CaptchaVerify 验证码验证
func (c *userLoginController) CaptchaVerify(ctx *gin.Context) {
	verifyValue := ctx.DefaultQuery("captcha", "")
	captchaId := ctx.DefaultQuery("captcha_id", "")
	if verifyValue == "" {
		log.New(ctx).WithCode(code_errors.SessionGetCaptchaEmptyError).Error("")
		response.New(ctx).WithCode(code_errors.SessionGetCaptchaEmptyError).Json()
		return
	}
	if captchaId == "" {
		log.New(ctx).WithCode(code_errors.CaptchaNotFoundError).Error("")
		response.New(ctx).WithCode(code_errors.CaptchaNotFoundError).Json()
		return
	}

	if err := c.service.CaptchaVerify(ctx, captchaId, verifyValue); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Captcha2 验证码
func (c *userLoginController) Captcha2(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")

	result, err := c.service.Captcha2(ctx)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	ctx.Writer.Write(result)
}

// Captcha2Verify 验证码验证
func (c *userLoginController) Captcha2Verify(ctx *gin.Context) {
	value := ctx.DefaultQuery("captcha_id", "")
	if value == "" {
		log.New(ctx).WithCode(code_errors.SessionGetCaptchaEmptyError).Error("")
		response.New(ctx).WithCode(code_errors.SessionGetCaptchaEmptyError).Json()
		return
	}

	session := sessions.Default(ctx)
	captchaId := session.Get("captcha_id")
	if captchaId == nil {
		log.New(ctx).WithCode(code_errors.CaptchaNotFoundError).Error("")
		response.New(ctx).WithCode(code_errors.CaptchaNotFoundError).Json()
		return
	}
	session.Delete("captcha")
	_ = session.Save()

	if err := c.service.Captcha2Verify(ctx, captchaId.(string), value); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
}
