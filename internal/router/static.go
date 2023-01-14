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
	"gin-admin/internal/handler/static"

	"github.com/gin-gonic/gin"
)

// NewStaticApi 静态资源路由
func NewStaticApi(engine *gin.Engine) {
	engine.GET("/", static.Index)
	engine.GET("/favicon.ico", static.FaviconIco)
}
