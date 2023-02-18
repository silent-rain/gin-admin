/*
 * @Author: silent-rain
 * @Date: 2023-01-06 00:00:18
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:34:35
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/router/router.go
 * @Descripttion:路由
 */
package router

import (
	"gin-admin/internal/controller"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init(engine *gin.Engine) {
	// 指定受信任的代理
	// engine.SetTrustedProxies([]string{"127.0.0.1"})
	engine.SetTrustedProxies(nil)

	// 设置静态资源
	setStaticApi(engine)

	// 服务健康检查
	engine.GET("/api/ping", controller.Ping)

	// 系统路由
	NewSystemApiV1(engine)
}
