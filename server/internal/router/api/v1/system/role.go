/*角色管理*/
package system

import (
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitRoleRouter 初始化角色管理路由
func InitRoleRouter(group *gin.RouterGroup) {
	// 角色管理
	role := group.Group("/role")
	{
		// 获取所有角色列表
		role.GET("/all", system.NewRoleController().All)
		// 获取角色列表
		role.GET("/list", system.NewRoleController().List)
		// 添加角色
		role.POST("/add", system.NewRoleController().Add)
		// 更新角色
		role.PUT("/update", system.NewRoleController().Update)
		// 删除角色
		role.DELETE("/delete", system.NewRoleController().Delete)
		// 批量删除角色
		role.DELETE("/batchDelete", system.NewRoleController().BatchDelete)
		// 更新角色状态
		role.PUT("/status", system.NewRoleController().Status)
	}
}
