// Package router 路由
package router

import (
	"github.com/silent-rain/gin-admin/internal/app/public/controller"
	apiauth "github.com/silent-rain/gin-admin/internal/router/api/v1/api_auth"
	datacenter "github.com/silent-rain/gin-admin/internal/router/api/v1/data_center"
	"github.com/silent-rain/gin-admin/internal/router/api/v1/log"
	"github.com/silent-rain/gin-admin/internal/router/api/v1/permission"
	"github.com/silent-rain/gin-admin/internal/router/api/v1/system"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init(engine *gin.Engine) {
	// 指定受信任的代理
	// engine.SetTrustedProxies([]string{"127.0.0.1"})
	engine.SetTrustedProxies(nil)

	// 初始化静态资源路由
	InitStaticRouter(engine)
	// 404 接口不存在
	engine.NoRoute(controller.NotFound)

	// 后端路由
	beGroup := engine.Group("/")

	// 服务连接测试
	beGroup.GET("/api/ping", controller.Ping)
	// 健康检查
	beGroup.GET("/api/health", controller.Health)

	// 公开根路由组
	PublicGroup := beGroup.Group("/api/v1")
	{
		system.InitCaptchaRouter(PublicGroup)       // 初始化验证码路由
		system.InitLoginRegisterRouter(PublicGroup) // 初始化注册/登录/登出路由
	}

	// 私有路由组
	privateGroup := beGroup.Group("/api/v1")
	{
		// 权限管理
		permission.InitUserRouter(privateGroup)         // 初始化用户管理路由
		permission.InitRoleRouter(privateGroup)         // 初始化角色管理路由
		permission.InitMenuRouter(privateGroup)         // 初始化菜单管理路由
		permission.InitRoleMenuRelRouter(privateGroup)  // 初始化角色菜单关系管理路由
		permission.InitUserApiTokenRouter(privateGroup) // 初始化用户API接口Token令牌管理路由
		// 系统管理
		system.InitUploadRouter(privateGroup)    // 初始化上传管理路由
		system.InitConfigRouter(privateGroup)    // 初始化配置管理路由
		system.InitUserLoginRouter(privateGroup) // 初始化配置管理路由
		// API 管理
		apiauth.InitApiHttpRouter(privateGroup)        // 初始化Http协议接口管理路由
		apiauth.InitApiRoleHttpRelRouter(privateGroup) // 初始化角色与Http协议接口关系管理路由
		// 数据中心
		datacenter.InitDictRouter(privateGroup)     // 初始化字典维度管理路由
		datacenter.InitDictDataRouter(privateGroup) // 初始化字典数据管理路由

		log.InitLogRouter(privateGroup) // 初始化日志管理路由
	}
}
