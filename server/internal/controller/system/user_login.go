/*用户登录/登出
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// 用户登录/登出
type userLoginController struct {
	service systemService.UserLoginService
}

// NewUserLoginController 创建用户登录/登出 对象
func NewUserLoginController() *userLoginController {
	return &userLoginController{
		service: systemService.NewUserLoginService(),
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
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}
