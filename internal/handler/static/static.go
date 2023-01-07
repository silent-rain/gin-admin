/**静态资源
 */
package static

import (
	"net/http"

	"gin-admin/internal/assets"

	"github.com/gin-gonic/gin"
)

// FaviconIco 网站 favicon.ico
func FaviconIco(c *gin.Context) {
	c.Header("content-type", "image/x-icon")
	c.FileFromFS("dist/favicon.ico", http.FS(assets.WebAssets))
	c.Status(200)
}

// Index 网站首页
func Index(c *gin.Context) {
	c.Header("content-type", "text/html;charset=utf-8")
	c.HTML(http.StatusOK, "index.html", nil)
}
