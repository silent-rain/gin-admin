// Package permission 用户API接口Token令牌表
package permission

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/controller"

	"github.com/gin-gonic/gin"
)

// InitUserApiTokenRouter 初始化用户API接口Token令牌管理路由
func InitUserApiTokenRouter(group *gin.RouterGroup) {
	userApiToken := group.Group("/userApiToken")
	api := controller.NewUserApiTokenController()
	{
		// 获取 Token 令牌列表
		userApiToken.GET("/list", api.List)
		// 添加 Token 令牌
		userApiToken.POST("/add", api.Add)
		// 更新 Token 令牌
		userApiToken.PUT("/update", api.Update)
		// 删除 Token 令牌
		userApiToken.DELETE("/delete", api.Delete)
		// 批量删除 Token 令牌
		userApiToken.DELETE("/batchDelete", api.BatchDelete)
		// 更新 Token 令牌状态
		userApiToken.PUT("/status", api.Status)
	}
}
