// Package controller 接口测试
package controller

import (
	"net/http"
	"time"

	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/conf"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	"github.com/silent-rain/gin-admin/pkg/errcode"
	timeutil "github.com/silent-rain/gin-admin/pkg/utils/time"

	"github.com/gin-gonic/gin"
)

// Ping 服务连接测试
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
	response.New(ctx).Json()
}

// Health 服务健康检查
func Health(ctx *gin.Context) {
	result := dto.Health{
		Timestamp:   time.Now().Local().Format(timeutil.CSTMilliLayout),
		Environment: conf.Instance().Environment.Env,
		Host:        ctx.Request.Host,
		Status:      "ok",
	}
	response.New(ctx).WithData(result).Json()
}

// NotFound 404 接口不存在
func NotFound(ctx *gin.Context) {
	response.New(ctx).WithHttpStatus(http.StatusNotFound).WithCode(errcode.RouteNotFoundError).Json()
}
