/*系统入口
 *
 */
package main

import (
	"fmt"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/middleware"
	"gin-admin/internal/pkg/plugin"
	"gin-admin/internal/pkg/repository/mysql"
	"gin-admin/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	conf.Init(conf.ConfigFile)
	// 初始化日志
	log.Init()
	// 初始化数据库
	mysql.Init()

	// 调试模式
	gin.SetMode(conf.Instance().Environment.Active())
	// 强制终端日志有色显示
	gin.ForceConsoleColor()

	engine := gin.Default()

	// 针对路由组的中间件
	{
		// 异常恢复
		engine.Use(gin.Recovery())
		// 鉴权表
		engine.Use(middleware.AuthTable())
		// 接口限流
		engine.Use(middleware.RateLimiter())
		// 跨域处理
		engine.Use(middleware.Cros())
		// 登录验证
		engine.Use(middleware.CheckLogin())
		// Session
		engine.Use(middleware.Session())

		// 在请求的时候会在控制台打印一行请求地址的url和耗时等信息
		engine.Use(gin.Logger())
		// zap 接收 gin 框架默认的日志
		// engine.Use(middleware.GinZapLogger(), middleware.GinZapRecovery(true))
		// 日志链路跟踪中间件
		engine.Use(middleware.TraceLogger())
		// 接口请求日志中间件，日志输出至数据库
		engine.Use(middleware.HttpLogger())
	}

	// 路由初始化
	router.Init(engine)

	// 插件
	plugin.Init(engine)

	// 服务运行
	if err := engine.Run(conf.Instance().Server.ServerAddress()); err != nil {
		panic(fmt.Sprintf("server run failed, err: %v", err))
	}
}
