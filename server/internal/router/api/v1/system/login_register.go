/*注册/登录/登出/验证码*/
package system

import (
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitLoginRegisterRouter 初始化注册/登录/登出/验证码路由
func InitLoginRegisterRouter(group *gin.RouterGroup) {
	// 注册/登录/登出/验证码
	userLogin := group.Group("/")
	{
		// 注册
		userLogin.POST("/register", system.NewUserController().Add)
		// 验证码
		userLogin.GET("/captcha", system.NewUserLoginController().Captcha)
		// 验证码验证
		userLogin.GET("/captcha/verify", system.NewUserLoginController().CaptchaVerify)
		// 登录
		userLogin.POST("/login", system.NewUserLoginController().Login)
		// 登出
		userLogin.POST("/logout", system.NewUserLoginController().Logout)
	}
}
