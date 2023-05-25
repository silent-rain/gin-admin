/*WEB 日志
 */
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/log/dto"
	"github.com/silent-rain/gin-admin/internal/app/log/model"
	"github.com/silent-rain/gin-admin/internal/app/log/service"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// WEB 日志
type webLogController struct {
	service service.WebLogService
}

// NewWebLogController 创建 WEB 日志对象
func NewWebLogController() *webLogController {
	return &webLogController{
		service: service.NewWebLogService(),
	}
}

// List 获取 WEB 日志列表
func (c *webLogController) List(ctx *gin.Context) {
	req := dto.QueryWebLogReq{}
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

// Add 添加 WEB 日志
func (c *webLogController) Add(ctx *gin.Context) {
	req := dto.AddWebLogReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := model.WebLog{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean.UserId = core.GetContext(ctx).UserId
	bean.Nickname = core.GetContext(ctx).Nickname
	bean.TraceId = core.GetContext(ctx).TraceId

	_, err := c.service.Add(ctx, bean)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
