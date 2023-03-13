/*注册/登录/登出/验证码*/
package system

import (
	"github.com/silent-rain/gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitLoginRegisterRouter 初始化注册/登录/登出/验证码路由
func InitLoginRegisterRouter(group *gin.RouterGroup) {
	// 注册/登录/登出/验证码
	userLogin := group.Group("/")
	{
		// 注册
		userLogin.POST("/register", system.NewUserLoginRegisterController().Register)
		// 登录
		userLogin.POST("/login", system.NewUserLoginRegisterController().Login)
		// 登出
		userLogin.POST("/logout", system.NewUserLoginRegisterController().Logout)
	}
}
