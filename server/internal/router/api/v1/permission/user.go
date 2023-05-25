// Package permission 用户管理
package permission

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/controller"

	"github.com/gin-gonic/gin"
)

// InitUserRouter 初始化用户管理路由
func InitUserRouter(group *gin.RouterGroup) {
	// 用户管理
	user := group.Group("/user")
	{
		// 获取用户信息
		user.GET("/info", controller.NewUserController().Info)
		// 获取所有用户列表
		user.GET("/all", controller.NewUserController().All)
		// 获取用户列表
		user.GET("/list", controller.NewUserController().List)
		// 添加用户
		user.POST("/add", controller.NewUserController().Add)
		// 更新用户详情信息
		user.PUT("/update", controller.NewUserController().Update)
		// 删除用户
		user.DELETE("/delete", controller.NewUserController().Delete)
		// 批量删除用户
		user.DELETE("/batchDelete", controller.NewUserController().BatchDelete)
	}

	{
		// 更新用户状态
		user.PUT("/status", controller.NewUserController().Status)
		// 重置用户密码
		user.PUT("/resetPwd", controller.NewUserController().ResetPassword)
		// 更新用户密码
		user.PUT("/updatePwd", controller.NewUserController().UpdatePassword)
		// 更新用户手机号码
		user.PUT("/updatePhone", controller.NewUserController().UpdatePhone)
		// 更新用户邮箱
		user.PUT("/updateEmail", controller.NewUserController().UpdateEmail)
	}
}
