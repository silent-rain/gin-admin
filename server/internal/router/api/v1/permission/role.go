/*角色管理*/
package permission

import (
	"github.com/silent-rain/gin-admin/internal/controller/permission"

	"github.com/gin-gonic/gin"
)

// InitRoleRouter 初始化角色管理路由
func InitRoleRouter(group *gin.RouterGroup) {
	// 角色管理
	role := group.Group("/role")
	{
		// 获取所有角色列表
		role.GET("/all", permission.NewRoleController().All)
		// 获取角色列表
		role.GET("/list", permission.NewRoleController().List)
		// 添加角色
		role.POST("/add", permission.NewRoleController().Add)
		// 更新角色
		role.PUT("/update", permission.NewRoleController().Update)
		// 删除角色
		role.DELETE("/delete", permission.NewRoleController().Delete)
		// 批量删除角色
		role.DELETE("/batchDelete", permission.NewRoleController().BatchDelete)
		// 更新角色状态
		role.PUT("/status", permission.NewRoleController().Status)
	}
}
