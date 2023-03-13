/*菜单管理*/
package permission

import (
	"github.com/silent-rain/gin-admin/internal/controller/permission"

	"github.com/gin-gonic/gin"
)

// InitMenuRouter 初始化菜单管理路由
func InitMenuRouter(group *gin.RouterGroup) {
	// 菜单管理
	menu := group.Group("/menu")
	{
		// 获取所有菜单树
		menu.GET("/allTree", permission.NewMenuController().AllTree)
		// 获取菜单树
		menu.GET("/tree", permission.NewMenuController().Tree)
		// 添加菜单
		menu.POST("/add", permission.NewMenuController().Add)
		// 更新菜单
		menu.PUT("/update", permission.NewMenuController().Update)
		// 删除菜单
		menu.DELETE("/delete", permission.NewMenuController().Delete)
		// 批量删除菜单
		menu.DELETE("/batchDelete", permission.NewMenuController().BatchDelete)
		// 更新菜单状态
		menu.PUT("/status", permission.NewMenuController().Status)
		// 通过父 ID 获取子配置列表
		menu.GET("/childrenMenu", permission.NewMenuController().ChildrenMenu)
	}
}
