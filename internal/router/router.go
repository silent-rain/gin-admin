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
	"gin-admin/internal/handler"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
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
