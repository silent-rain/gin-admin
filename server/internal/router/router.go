/*路由
 */
package router

import (
	"gin-admin/internal/controller"
	"gin-admin/internal/router/apiv1"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init(engine *gin.Engine) {
	// 指定受信任的代理
	// engine.SetTrustedProxies([]string{"127.0.0.1"})
	engine.SetTrustedProxies(nil)

	// 设置静态资源路由
	setStaticRouter(engine)

	// 服务连接测试
	engine.GET("/api/ping", controller.Ping)
	// 健康检查
	engine.GET("/api/health", controller.Health)

	// 系统路由
	apiv1.NewSystemApi(engine)

	// 404 接口不存在
	engine.NoRoute(controller.NotFound)
}
