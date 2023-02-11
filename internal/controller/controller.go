/*接口测试*/
package controller

import (
	"net/http"

	"gin-admin/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// Ping 服务健康检查
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{})
}

// SayHello 接口测试
func SayHello(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	response.New(ctx).WithData("hello," + name).Json()
}
