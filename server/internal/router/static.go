/*
 * @Author: silent-rain
 * @Date: 2023-01-07 16:24:58
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 18:49:41
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/router/static.go
 * @Descripttion:静态资源路由
 */
package router

import (
	"html/template"
	"net/http"

	"gin-admin/assets"
	"gin-admin/internal/controller"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 设置静态资源路由
func setStaticRouter(engine *gin.Engine) {
	// 加载静态资源
	engine.StaticFS("/static", http.FS(utils.NewResource()))
	// Api Docs 静态内嵌资源
	engine.StaticFS("/docs", http.FS(utils.NewDocsResource()))
	// 本地静态资源
	engine.Static("/upload", conf.Instance().Server.Upload.FilePath)
	// WEB 首页模板
	templ := template.Must(template.New("").ParseFS(assets.WebAssets, "dist/*.html"))
	engine.SetHTMLTemplate(templ)

	engine.GET("/", controller.Index)
	engine.GET("/favicon.ico", controller.FaviconIco)
}
