/*角色与Http协议接口关联表*/
package apiauth

import (
	apiauth "gin-admin/internal/controller/api_auth"

	"github.com/gin-gonic/gin"
)

// InitApiRoleHttpRelRouter 初始化角色与Http协议接口关系管理路由
func InitApiRoleHttpRelRouter(group *gin.RouterGroup) {
	roleMenuRel := group.Group("/apiRoleHttpRel")
	{
		// 获取角色与Http协议接口关系列表
		roleMenuRel.GET("/list", apiauth.NewApiRoleHttpRelController().List)
		// 更新角色与Http协议接口关系
		roleMenuRel.PUT("/update", apiauth.NewApiRoleHttpRelController().Update)
	}
}
