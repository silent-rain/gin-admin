/*
 * @Author: silent-rain
 * @Date: 2023-01-05 00:22:11
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 00:50:09
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/cmd/main.go
 * @Descripttion:
 */
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
	"gin-admin/internal/pkg/database"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/middleware"
	"gin-admin/internal/pkg/utils"
	"gin-admin/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 配置初始化
	conf.InitLoadConfig(conf.ConfigFile)
	// 数据库初始化
	database.Init()
	// 日志初始化
	log.Init()

	// 调试模式
	gin.SetMode(conf.Instance().EnvConfig.Env())
	// 强制终端日志有色显示
	gin.ForceConsoleColor()

	engine := gin.Default()
	// 跨域处理(要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404)
	engine.Use(middleware.Cros(), gin.Recovery())
	// 日志中间件，日志输出至数据库
	engine.Use(middleware.GinLogger())
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
