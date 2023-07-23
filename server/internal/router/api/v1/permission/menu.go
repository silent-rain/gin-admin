// Package permission 菜单管理
package permission

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/controller"

	"github.com/gin-gonic/gin"
)

// InitMenuRouter 初始化菜单管理路由
func InitMenuRouter(group *gin.RouterGroup) {
	// 菜单管理
	menu := group.Group("/menu")
	{
		// 获取所有菜单树
		menu.GET("/allTree", controller.NewMenuController().AllTree)
		// 获取菜单树
		menu.GET("/tree", controller.NewMenuController().Tree)
		// 添加菜单
		menu.POST("/add", controller.NewMenuController().Add)
		// 更新菜单
		menu.PUT("/update", controller.NewMenuController().Update)
		// 删除菜单
		menu.DELETE("/delete", controller.NewMenuController().Delete)
		// 批量删除菜单
		menu.DELETE("/batchDelete", controller.NewMenuController().BatchDelete)
		// 更新菜单状态
		menu.PUT("/updateStatus", controller.NewMenuController().UpdateStatus)
		// 通过父 ID 获取子配置列表
		menu.GET("/childrenMenus", controller.NewMenuController().ChildrenMenus)
	}
}
