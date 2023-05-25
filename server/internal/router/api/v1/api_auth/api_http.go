/*Http协议接口管理表*/
package apiauth

import (
	"github.com/silent-rain/gin-admin/internal/app/api_auth/controller"

	"github.com/gin-gonic/gin"
)

// InitApiHttpRouter 初始化Http协议接口管理路由
func InitApiHttpRouter(group *gin.RouterGroup) {
	role := group.Group("/apiHttp")
	{
		// 获取所有Http协议接口信息树
		role.GET("/allTree", controller.NewApiHttpController().AllTree)
		// 获取Http协议接口信息树
		role.GET("/tree", controller.NewApiHttpController().Tree)
		// 添加Http协议接口信息
		role.POST("/add", controller.NewApiHttpController().Add)
		// 更新Http协议接口信息
		role.PUT("/update", controller.NewApiHttpController().Update)
		// 删除Http协议接口信息
		role.DELETE("/delete", controller.NewApiHttpController().Delete)
		// 批量删除Http协议接口信息
		role.DELETE("/batchDelete", controller.NewApiHttpController().BatchDelete)
		// 更新Http协议接口信息状态
		role.PUT("/status", controller.NewApiHttpController().Status)
	}
}
