/*用户管理*/
package system

import (
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitUserRouter 初始化用户管理路由
func InitUserRouter(group *gin.RouterGroup) {
	// 用户管理
	user := group.Group("/user")
	{
		// 获取用户信息
		user.GET("/info", system.NewUserController().Info)
		// 获取所有用户列表
		user.GET("/all", system.NewUserController().All)
		// 获取用户列表
		user.GET("/list", system.NewUserController().List)
		// 添加用户
		user.POST("/add", system.NewUserController().Add)
		// 更新用户详情信息
		user.PUT("/update", system.NewUserController().Update)
		// 删除用户
		user.DELETE("/delete", system.NewUserController().Delete)
		// 批量删除用户
		user.DELETE("/batchDelete", system.NewUserController().BatchDelete)
	}

	{
		// 更新用户状态
		user.PUT("/status", system.NewUserController().Status)
		// 重置用户密码
		user.PUT("/resetPwd", system.NewUserController().ResetPassword)
		// 更新用户密码
		user.PUT("/updatePwd", system.NewUserController().UpdatePassword)
		// 更新用户手机号码
		user.PUT("/updatePhone", system.NewUserController().UpdatePhone)
		// 更新用户邮箱
		user.PUT("/updateEmail", system.NewUserController().UpdateEmail)
	}
}
