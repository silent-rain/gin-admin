// Package plugin 插件初始化
package plugin

import "github.com/gin-gonic/gin"

// Init 插件初始化
func Init(engine *gin.Engine) {
	// Pprof 性能剖析工具
	RegisterPprof(engine)
	// Prometheus 监控指标
	RegisterPrometheus(engine)
	// swagger API 文档
	RegisterSwagger(engine)
	// 服务启动后显示 logo
	RegisterLogo()
	//  服务启动后显示 IP 地址
	RegisterAddr()
	// 服务启动后在浏览器中打开 URI
	RegisterOpenBrowser()
}
