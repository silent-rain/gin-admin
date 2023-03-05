/*用户登录信息表*/
package system

import (
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitUserLoginRouter 初始化配置管理路由
func InitUserLoginRouter(group *gin.RouterGroup) {
	// 配置管理
	config := group.Group("/userLogin")
	{
		// 获取用户登录信息列表
		config.GET("/list", system.NewUserLoginController().List)
		// 更新用户登录信息状态
		config.PUT("/status", system.NewUserLoginController().Status)
	}
}
