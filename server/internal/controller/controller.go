/*接口测试*/
package controller

import (
	"net/http"
	"time"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/response"
	systemVO "gin-admin/internal/vo/system"

	"github.com/gin-gonic/gin"
)

// Ping 服务连接测试
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
	response.New(ctx).Json()
}

// Health 服务健康检查
func Health(ctx *gin.Context) {
	result := systemVO.Health{
		Timestamp:   time.Now(),
		Environment: conf.Instance().Environment.Active(),
		Host:        ctx.Request.Host,
		Status:      "ok",
	}
	response.New(ctx).WithData(result).Json()
}