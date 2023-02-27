/*菜单管理*/
package system

import (
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitMenuRouter 初始化菜单管理路由
func InitMenuRouter(group *gin.RouterGroup) {
	// 菜单管理
	menu := group.Group("/menu")
	{
		// 获取所有菜单树
		menu.GET("/allTree", system.NewMenuController().AllTree)
		// 获取菜单树
		menu.GET("/tree", system.NewMenuController().Tree)
		// 添加菜单
		menu.POST("/add", system.NewMenuController().Add)
		// 更新菜单
		menu.PUT("/update", system.NewMenuController().Update)
		// 删除菜单
		menu.DELETE("/delete", system.NewMenuController().Delete)
		// 批量删除菜单
		menu.DELETE("/batchDelete", system.NewMenuController().BatchDelete)
		// 更新菜单状态
		menu.PUT("/status", system.NewMenuController().Status)
	}
}
