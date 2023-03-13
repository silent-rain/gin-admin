/*用户管理*/
package permission

import (
	"github.com/silent-rain/gin-admin/internal/controller/permission"

	"github.com/gin-gonic/gin"
)

// InitUserRouter 初始化用户管理路由
func InitUserRouter(group *gin.RouterGroup) {
	// 用户管理
	user := group.Group("/user")
	{
		// 获取用户信息
		user.GET("/info", permission.NewUserController().Info)
		// 获取所有用户列表
		user.GET("/all", permission.NewUserController().All)
		// 获取用户列表
		user.GET("/list", permission.NewUserController().List)
		// 添加用户
		user.POST("/add", permission.NewUserController().Add)
		// 更新用户详情信息
		user.PUT("/update", permission.NewUserController().Update)
		// 删除用户
		user.DELETE("/delete", permission.NewUserController().Delete)
		// 批量删除用户
		user.DELETE("/batchDelete", permission.NewUserController().BatchDelete)
	}

	{
		// 更新用户状态
		user.PUT("/status", permission.NewUserController().Status)
		// 重置用户密码
		user.PUT("/resetPwd", permission.NewUserController().ResetPassword)
		// 更新用户密码
		user.PUT("/updatePwd", permission.NewUserController().UpdatePassword)
		// 更新用户手机号码
		user.PUT("/updatePhone", permission.NewUserController().UpdatePhone)
		// 更新用户邮箱
		user.PUT("/updateEmail", permission.NewUserController().UpdateEmail)
	}
}
