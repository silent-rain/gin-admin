// Package permission 角色管理
package permission

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/controller"

	"github.com/gin-gonic/gin"
)

// InitRoleRouter 初始化角色管理路由
func InitRoleRouter(group *gin.RouterGroup) {
	// 角色管理
	role := group.Group("/role")
	{
		// 获取所有角色列表
		role.GET("/all", controller.NewRoleController().All)
		// 获取角色列表
		role.GET("/list", controller.NewRoleController().List)
		// 添加角色
		role.POST("/add", controller.NewRoleController().Add)
		// 更新角色
		role.PUT("/update", controller.NewRoleController().Update)
		// 删除角色
		role.DELETE("/delete", controller.NewRoleController().Delete)
		// 批量删除角色
		role.DELETE("/batchDelete", controller.NewRoleController().BatchDelete)
		// 更新角色状态
		role.PUT("/updateStatus", controller.NewRoleController().UpdateStatus)
	}
}
