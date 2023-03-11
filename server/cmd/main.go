/*系统入口
 *
 */
package main

import (
	"errors"
	"fmt"
	"net/http"

	"gin-admin/internal/initialize"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/middleware"
	_ "gin-admin/internal/pkg/repository/cache"
	"gin-admin/internal/pkg/repository/mysql"
	"gin-admin/internal/pkg/repository/redis"
	"gin-admin/internal/router"
	"gin-admin/pkg/plugin"
	"gin-admin/pkg/shutdown"
	"gin-admin/tasks"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	conf.Init(conf.ConfigFile)
	// 初始化日志
	log.Init()
	// 初始化 Mysql 数据库
	mysql.Init()
	// 初始化 Redis 数据库
	redis.Init()
	// 初始化定时任务
	tasks.Init()
	// 初始化数据
	initialize.Init()

	// 调试模式
	gin.SetMode(conf.Instance().Environment.Active())
	// 强制终端日志有色显示
	gin.ForceConsoleColor()

	engine := gin.Default()

	// 全局路由的中间件
	{
		// 跨域处理
		engine.Use(middleware.Cros())
		// 鉴权表
		engine.Use(middleware.AuthTable())
		// 日志链路跟踪
		engine.Use(middleware.TraceLogger())
		// Session
		engine.Use(middleware.Session())
		// 接口限流
		engine.Use(middleware.RateLimiter())
		// 检查 API 令牌鉴权中间件
		engine.Use(middleware.CheckApiLogin())
		// 登录验证
		engine.Use(middleware.CheckLogin())
		// 指标记录
		engine.Use(middleware.Metrics())
		// LoadTls 添加对 HTTPS 的支持
		// engine.Use(middleware.LoadTls())

		// 在请求的时候会在控制台打印一行请求地址的url和耗时等信息
		engine.Use(gin.Logger())
		// zap 接收 gin 框架默认的日志
		// engine.Use(middleware.GinZapLogger(), middleware.GinZapRecovery(true))
		// 接口请求日志中间件，日志输出至数据库
		engine.Use(middleware.HttpLogger())

		// 异常恢复
		engine.Use(middleware.Recover())
	}

	// 初始化路由
	router.Init(engine)

	// 初始化插件
	plugin.Init(engine)

	srv := &http.Server{
		Addr:    conf.Instance().Server.ServerAddress(),
		Handler: engine,
	}

	// 启动服务
	go func() {
		// srv.ListenAndServeTLS("ssl.pem", "ssl.key")  // 开启 SSL 服务
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("server run failed, err: %v", err))
		}
	}()

	// 关闭资源
	shutdown.NewHook().Close(
		// 关闭 Http 服务
		shutdown.WithCloseHttpServer(srv),
		// 关闭定时任务
		shutdown.WithCloseCron(),
		// 关闭 Mysql 服务
		shutdown.WithCloseMysql(),
		// 关闭 Redis 服务
		shutdown.WithCloseRedis(),
		// 服务关闭后的消息提示
		shutdown.WithCloseInfo(),
	)
}
