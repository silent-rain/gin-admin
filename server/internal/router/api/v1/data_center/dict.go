/*字典维度管理*/
package datacenter

import (
	datacenter "github.com/silent-rain/gin-admin/internal/controller/data_center"

	"github.com/gin-gonic/gin"
)

// InitDictRouter 初始化字典维度管理路由
func InitDictRouter(group *gin.RouterGroup) {
	router := group.Group("/dict")
	controller := datacenter.NewDictController()
	{
		// 获取字典维度信息列表
		router.GET("/list", controller.List)
		// 添加字典维度信息
		router.POST("/add", controller.Add)
		// 更新字典维度信息
		router.PUT("/update", controller.Update)
		// 删除字典维度信息
		router.DELETE("/delete", controller.Delete)
		// 批量删除字典维度信息
		router.DELETE("/batchDelete", controller.BatchDelete)
		// 更新字典维度信息状态
		router.PUT("/status", controller.Status)
	}
}
