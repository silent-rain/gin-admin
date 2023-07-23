// Package service 验证码
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/captcha"
	"github.com/silent-rain/gin-admin/pkg/constant"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// CaptchaService 验证码接口
type CaptchaService interface {
	Captcha(ctx *gin.Context) (dto.Captcha, error)
	CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) error
}

// 验证码
type captchaService struct {
}

// NewCaptchaService 创建验证码对象
func NewCaptchaService() *captchaService {
	return &captchaService{}
}

// Captcha 验证码
func (h *captchaService) Captcha(ctx *gin.Context) (dto.Captcha, error) {
	result := dto.Captcha{
		CaptchaId: "",
		B64s:      "",
	}

	captchaId, b64s, err := captcha.NewCaptcha().MekeCaptcha(constant.CaptchaType)
	if err != nil {
		log.New(ctx).WithCode(errcode.CaptchaGenerateError).Errorf("%v", err)
		return result, errcode.CaptchaGenerateError
	}

	result.CaptchaId = captchaId
	result.B64s = b64s
	return result, nil
}

// CaptchaVerify 验证码验证
func (h *captchaService) CaptchaVerify(ctx *gin.Context, captchaId string, verifyValue string) error {
	// 校验验证码
	// 注意 Verify(id, VerifyValue, true) 中的 true参数
	// 当为 true 时，校验 传入的id 的验证码，校验完 这个ID的验证码就要在内存中删除
	// 当为 false 时，校验 传入的id 的验证码，校验完 这个ID的验证码不删除
	if !captcha.CaptchaStore.Verify(captchaId, verifyValue, true) {
		log.New(ctx).WithCode(errcode.CaptchaVerifyError).Error("")
		return errcode.CaptchaVerifyError
	}
	return nil
}

// 检查验证码
func chechkCaptcha(ctx *gin.Context, captchaId, captchaValue string) error {
	if captchaValue == "" {
		log.New(ctx).WithCode(errcode.SessionGetCaptchaEmptyError).Error("")
		return errcode.SessionGetCaptchaEmptyError
	}
	if captchaId == "" {
		log.New(ctx).WithCode(errcode.CaptchaNotFoundError).Error("")
		return errcode.CaptchaNotFoundError
	}

	// 校验验证码
	if !captcha.CaptchaStore.Verify(captchaId, captchaValue, true) {
		log.New(ctx).WithCode(errcode.CaptchaVerifyError).Error("")
		return errcode.CaptchaVerifyError
	}
	return nil
}
