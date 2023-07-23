// Package system 配置管理
package system

import (
	"github.com/silent-rain/gin-admin/internal/app/system/controller"

	"github.com/gin-gonic/gin"
)

// InitConfigRouter 初始化配置管理路由
func InitConfigRouter(group *gin.RouterGroup) {
	// 配置管理
	config := group.Group("/config")
	{
		// 获取所有配置树
		config.GET("/allTree", controller.NewConfigController().AllTree)
		// 获取配置树
		config.GET("/tree", controller.NewConfigController().Tree)
		// 获取配置列表
		config.GET("/list", controller.NewConfigController().List)
		// 获取配置信息
		config.GET("/info", controller.NewConfigController().Info)
		// 添加配置
		config.POST("/add", controller.NewConfigController().Add)
		// 更新配置
		config.PUT("/update", controller.NewConfigController().Update)
		// 批量更新配置
		config.PUT("/batchUpdate", controller.NewConfigController().BatchUpdate)
		// 删除配置
		config.DELETE("/delete", controller.NewConfigController().Delete)
		// 批量删除配置
		config.DELETE("/batchDelete", controller.NewConfigController().BatchDelete)
		// 更新配置状态
		config.PUT("/updateStatus", controller.NewConfigController().UpdateStatus)
		// 通过上级 key 获取子配置列表
		config.GET("/childrensByKey", controller.NewConfigController().ChildrensByKey)
		// 查询网站配置列表
		config.GET("/webSiteConfigList", controller.NewConfigController().WebSiteConfigList)
	}
}
