// Package router 静态资源路由
package router

import (
	"html/template"
	"net/http"

	"github.com/silent-rain/gin-admin/assets"
	"github.com/silent-rain/gin-admin/internal/app/public/controller"
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/pkg/resource"

	"github.com/gin-gonic/gin"
)

// 初始化静态资源路由
func InitStaticRouter(engine *gin.Engine) {
	// 加载静态资源
	engine.StaticFS("/static", http.FS(resource.NewResource()))
	// Api Docs 静态内嵌资源
	engine.StaticFS("/docs", http.FS(resource.NewDocsResource()))
	// 本地静态资源
	engine.Static("/upload", global.Instance().Config().Server.Upload.FilePath)
	// WEB 首页模板
	templ := template.Must(template.New("").ParseFS(assets.WebAssets, "dist/*.html"))
	engine.SetHTMLTemplate(templ)

	engine.GET("/", controller.Index)
	engine.GET("/favicon.ico", controller.FaviconIco)
}
