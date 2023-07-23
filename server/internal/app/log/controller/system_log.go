// Package controller 系统日志
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/log/dto"
	"github.com/silent-rain/gin-admin/internal/app/log/service"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 系统日志
type systemLogController struct {
	service *service.SystemLogService
}

// NewSystemLogController 创建系统日志对象
func NewSystemLogController() *systemLogController {
	return &systemLogController{
		service: service.NewSystemLogService(),
	}
}

// List 获取系统日志列表
func (c *systemLogController) List(ctx *gin.Context) {
	req := dto.QuerySystemLogReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}
