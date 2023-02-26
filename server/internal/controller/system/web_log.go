/*WEB 日志
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/core"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// WEB 日志
type webLogController struct {
	service systemService.WebLogService
}

// NewWebLogController 创建 WEB 日志对象
func NewWebLogController() *webLogController {
	return &webLogController{
		service: systemService.NewWebLogService(),
	}
}

// List 获取 WEB 日志列表
func (c *webLogController) List(ctx *gin.Context) {
	req := systemDTO.QueryWebLogReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Add 添加 WEB 日志
func (c *webLogController) Add(ctx *gin.Context) {
	req := systemDTO.AddWebLogReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := systemModel.WebLog{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean.UserId = core.GetContext(ctx).UserId
	bean.Nickname = core.GetContext(ctx).Nickname
	bean.TraceId = core.GetContext(ctx).TraceId

	_, err := c.service.Add(ctx, bean)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}
