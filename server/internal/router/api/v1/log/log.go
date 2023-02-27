/*日志*/
package log

import (
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// InitLogRouter 初始化日志管理路由
func InitLogRouter(group *gin.RouterGroup) {
	// 网络请求管理
	httpLog := group.Group("/httpLog")
	{
		// 获取网络请求日志列表
		httpLog.GET("/list", system.NewHttpLogController().List)
		httpLog.GET("/body", system.NewHttpLogController().GetBody)
	}

	// 系统日志管理
	systemLog := group.Group("/systemLog")
	{
		// 获取系统日志列表
		systemLog.GET("/list", system.NewSystemLogController().List)
	}

	// WEB 日志管理
	webLog := group.Group("/webLog")
	{
		// 获取 WEB 日志列表
		webLog.GET("/list", system.NewWebLogController().List)
		// 添加 WEB 日志
		webLog.POST("/add", system.NewWebLogController().Add)
	}
}
