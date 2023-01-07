/**系统入口文件
 *
 */
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"gin-admin/internal/assets"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/middleware"
	"gin-admin/internal/pkg/utils"
	"gin-admin/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	conf.InitLoadConfig(conf.ConfigFile)

	// 调试模式
	gin.SetMode(conf.Instance().EnvConfig.Env())
	// 强制终端日志有色显示
	gin.ForceConsoleColor()

	engine := gin.Default()
	// 跨域处理(要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404)
	engine.Use(middleware.Cros(), gin.Recovery())
	// 在请求的时候会在控制台打印一行请求地址的url和耗时等信息
	engine.Use(gin.Logger(), gin.Recovery())

	// 加载静态资源
	engine.StaticFS("/static", http.FS(utils.NewResource()))
	// 首页模板
	templ := template.Must(template.New("").ParseFS(assets.WebAssets, "dist/*.html"))
	engine.SetHTMLTemplate(templ)

	// 路由初始化
	router.Init(engine)
	// 服务运行
	if err := engine.Run(":8080"); err != nil {
		panic(fmt.Sprintf("server run failed, err: %v", err))
	}
}
