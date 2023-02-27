/*路由
 */
package router

import (
	"gin-admin/internal/controller"
	"gin-admin/internal/router/api/v1/log"
	"gin-admin/internal/router/api/v1/system"

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
		system.InitLoginRegisterRouter(PublicGroup) // 初始化注册/登录/登出/验证码路由
	}

	// 私有路由组
	privateGroup := beGroup.Group("/api/v1")
	// 路由
	{
		system.InitUploadRouter(privateGroup)      // 初始化上传管理路由
		system.InitUserRouter(privateGroup)        // 初始化用户管理路由
		system.InitRoleRouter(privateGroup)        // 初始化角色管理路由
		system.InitMenuRouter(privateGroup)        // 初始化菜单管理路由
		system.InitRoleMenuRelRouter(privateGroup) // 初始化角色菜单关系管理路由
		system.InitConfigRouter(privateGroup)      // 初始化配置管理路由
		log.InitLogRouter(privateGroup)            // 初始化日志管理路由
	}
}
