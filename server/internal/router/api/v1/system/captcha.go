// Package system 验证码
package system

import (
	"github.com/silent-rain/gin-admin/internal/app/system/controller"

	"github.com/gin-gonic/gin"
)

// InitCaptchaRouter 初始化验证码路由
func InitCaptchaRouter(group *gin.RouterGroup) {
	captcha := group.Group("/captcha")
	{
		// 验证码
		captcha.GET("", controller.NewCaptchaController().Captcha)
		// 验证码验证
		captcha.GET("/captcha/verify", controller.NewCaptchaController().CaptchaVerify)
	}
}
