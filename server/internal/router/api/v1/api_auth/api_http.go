/*Http协议接口管理表*/
package apiauth

import (
	apiauth "gin-admin/internal/controller/api_auth"

	"github.com/gin-gonic/gin"
)

// InitApiHttpRouter 初始化Http协议接口管理路由
func InitApiHttpRouter(group *gin.RouterGroup) {
	role := group.Group("/apiHttp")
	{
		// 获取所有Http协议接口信息列表
		role.GET("/all", apiauth.NewApiHttpController().All)
		// 获取Http协议接口信息列表
		role.GET("/list", apiauth.NewApiHttpController().List)
		// 添加Http协议接口信息
		role.POST("/add", apiauth.NewApiHttpController().Add)
		// 更新Http协议接口信息
		role.PUT("/update", apiauth.NewApiHttpController().Update)
		// 删除Http协议接口信息
		role.DELETE("/delete", apiauth.NewApiHttpController().Delete)
		// 批量删除Http协议接口信息
		role.DELETE("/batchDelete", apiauth.NewApiHttpController().BatchDelete)
		// 更新Http协议接口信息状态
		role.PUT("/status", apiauth.NewApiHttpController().Status)
	}
}
