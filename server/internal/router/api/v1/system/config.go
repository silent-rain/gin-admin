/*配置管理*/
package system

import (
	"github.com/silent-rain/gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitConfigRouter 初始化配置管理路由
func InitConfigRouter(group *gin.RouterGroup) {
	// 配置管理
	config := group.Group("/config")
	{
		// 获取所有配置树
		config.GET("/allTree", system.NewConfigController().AllTree)
		// 获取配置树
		config.GET("/tree", system.NewConfigController().Tree)
		// 获取配置列表
		config.GET("/list", system.NewConfigController().List)
		// 获取配置信息
		config.GET("/info", system.NewConfigController().Info)
		// 添加配置
		config.POST("/add", system.NewConfigController().Add)
		// 更新配置
		config.PUT("/update", system.NewConfigController().Update)
		// 批量更新配置
		config.PUT("/batchUpdate", system.NewConfigController().BatchUpdate)
		// 删除配置
		config.DELETE("/delete", system.NewConfigController().Delete)
		// 批量删除配置
		config.DELETE("/batchDelete", system.NewConfigController().BatchDelete)
		// 更新配置状态
		config.PUT("/status", system.NewConfigController().Status)
		// 通过上级 key 获取子配置列表
		config.GET("/childrenByKey", system.NewConfigController().ChildrenByKey)
		// 查询网站配置列表
		config.GET("/webSiteConfigList", system.NewConfigController().WebSiteConfigList)
	}
}
