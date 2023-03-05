/*WEB 日志
 */
package log

import (
	logDTO "gin-admin/internal/dto/log"
	logModel "gin-admin/internal/model/log"
	"gin-admin/internal/pkg/core"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	logService "gin-admin/internal/service/log"

	"github.com/gin-gonic/gin"
)

// WEB 日志
type webLogController struct {
	service logService.WebLogService
}

// NewWebLogController 创建 WEB 日志对象
func NewWebLogController() *webLogController {
	return &webLogController{
		service: logService.NewWebLogService(),
	}
}

// List 获取 WEB 日志列表
func (c *webLogController) List(ctx *gin.Context) {
	req := logDTO.QueryWebLogReq{}
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
	req := logDTO.AddWebLogReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := logModel.WebLog{}
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