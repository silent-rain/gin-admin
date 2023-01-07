/**路由
 */
package router

import (
	"gin-admin/internal/handler"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	// 指定受信任的代理
	// engine.SetTrustedProxies([]string{"127.0.0.1"})
	engine.SetTrustedProxies(nil)

	// 服务健康检查
	engine.GET("/ping", handler.Ping)

	//  静态资源路由
	NewStaticApi(engine)

	// api v1
	NewApiV1(engine)
}
