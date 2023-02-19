/*路由相关插件*/
package plugin

import (
	"gin-admin/internal/pkg/conf"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RegisterPprof pprof 性能剖析工具
func RegisterPprof(engine *gin.Engine) {
	if !conf.Instance().Server.Base.EnablePprof {
		return
	}
	pprof.Register(engine)
}

// RegisterPrometheus Prometheus 监控指标工具
func RegisterPrometheus(engine *gin.Engine) {
	if !conf.Instance().Server.Base.EnablePrometheus {
		return
	}
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

// RegisterSwagger swagger API 文档
func RegisterSwagger(engine *gin.Engine) {
	if !conf.Instance().Server.Base.EnableSwagger {
		return
	}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
