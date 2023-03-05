/*验证码*/
package system

import (
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitCaptchaRouter 初始化验证码路由
func InitCaptchaRouter(group *gin.RouterGroup) {
	captcha := group.Group("/")
	{
		// 验证码
		captcha.GET("/captcha", system.NewCaptchaController().Captcha)
		// 验证码验证
		captcha.GET("/captcha/verify", system.NewCaptchaController().CaptchaVerify)
	}
}
