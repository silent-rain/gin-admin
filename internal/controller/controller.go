/*接口测试*/
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 服务健康检查
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
