/*字典数据管理*/
package datacenter

import (
	datacenter "gin-admin/internal/controller/data_center"

	"github.com/gin-gonic/gin"
)

// InitDictDataRouter 初始化字典数据管理路由
func InitDictDataRouter(group *gin.RouterGroup) {
	router := group.Group("/dictData")
	controller := datacenter.NewDictDataController()
	{
		// 获取字典数据信息列表
		router.GET("/list", controller.List)
		// 添加字典数据信息
		router.POST("/add", controller.Add)
		// 更新字典数据信息
		router.PUT("/update", controller.Update)
		// 删除字典数据信息
		router.DELETE("/delete", controller.Delete)
		// 批量删除字典数据信息
		router.DELETE("/batchDelete", controller.BatchDelete)
		// 更新字典数据信息状态
		router.PUT("/status", controller.Status)
	}
}
