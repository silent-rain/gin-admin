// Package apiauth 角色与Http协议接口关联表
package apiauth

import (
	"github.com/silent-rain/gin-admin/internal/app/api_auth/controller"

	"github.com/gin-gonic/gin"
)

// InitApiRoleHttpRelRouter 初始化角色与Http协议接口关系管理路由
func InitApiRoleHttpRelRouter(group *gin.RouterGroup) {
	roleMenuRel := group.Group("/apiRoleHttpRel")
	{
		// 获取角色与Http协议接口关系列表
		roleMenuRel.GET("/list", controller.NewApiRoleHttpRelController().List)
		// 更新角色与Http协议接口关系
		roleMenuRel.PUT("/update", controller.NewApiRoleHttpRelController().Update)
	}
}
