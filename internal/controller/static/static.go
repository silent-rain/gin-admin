/*静态资源
 */
package controller

import (
	"net/http"

	"gin-admin/assets"

	"github.com/gin-gonic/gin"
)

// FaviconIco 网站 favicon.ico
func FaviconIco(ctx *gin.Context) {
	ctx.Header("content-type", "image/x-icon")
	ctx.FileFromFS("dist/favicon.ico", http.FS(assets.WebAssets))
	ctx.Status(200)
}

// Index 网站首页
func Index(ctx *gin.Context) {
	ctx.Header("content-type", "text/html;charset=utf-8")
	ctx.HTML(http.StatusOK, "index.html", nil)
}
