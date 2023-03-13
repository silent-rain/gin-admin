/*角色菜单关系管理*/
package permission

import (
	"github.com/silent-rain/gin-admin/internal/controller/permission"

	"github.com/gin-gonic/gin"
)

// InitRoleMenuRelRouter 初始化角色菜单关系管理路由
func InitRoleMenuRelRouter(group *gin.RouterGroup) {
	// 角色菜单关系管理
	roleMenuRel := group.Group("/roleMenuRel")
	{
		// 获取角色关联的菜单列表
		roleMenuRel.GET("/list", permission.NewRoleMenuRelController().List)
		// 更新角色菜单关联关系
		roleMenuRel.PUT("/update", permission.NewRoleMenuRelController().Update)
	}
}
