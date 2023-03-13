/*验证码*/
package system

import (
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// 验证码
type captchaController struct {
	service systemService.CaptchaService
}

// NewCaptchaController 创建验证码对象
func NewCaptchaController() *captchaController {
	return &captchaController{
		service: systemService.NewCaptchaService(),
	}
}

// Captcha 验证码
func (c *captchaController) Captcha(ctx *gin.Context) {
	result, err := c.service.Captcha(ctx)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}

// CaptchaVerify 验证码验证
func (c *captchaController) CaptchaVerify(ctx *gin.Context) {
	verifyValue := ctx.DefaultQuery("captcha", "")
	captchaId := ctx.DefaultQuery("captcha_id", "")
	if verifyValue == "" {
		log.New(ctx).WithCode(errcode.SessionGetCaptchaEmptyError).Error("")
		response.New(ctx).WithCode(errcode.SessionGetCaptchaEmptyError).Json()
		return
	}
	if captchaId == "" {
		log.New(ctx).WithCode(errcode.CaptchaNotFoundError).Error("")
		response.New(ctx).WithCode(errcode.CaptchaNotFoundError).Json()
		return
	}

	if err := c.service.CaptchaVerify(ctx, captchaId, verifyValue); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
