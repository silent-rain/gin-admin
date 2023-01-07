/**静态资源路由
 */
package router

import (
	"gin-admin/internal/handler/static"

	"github.com/gin-gonic/gin"
)

// NewStaticApi 静态资源路由
func NewStaticApi(engine *gin.Engine) {
	engine.GET("/", static.Index)
	engine.GET("/favicon.ico", static.FaviconIco)
}
