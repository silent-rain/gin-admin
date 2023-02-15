/*系统日志
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 系统日志
type systemLogController struct {
	service systemService.SystemLogService
}

// NewSystemLogController 创建系统日志对象
func NewSystemLogController() *systemLogController {
	return &systemLogController{
		service: systemService.NewSystemLogService(),
	}
}

// List 获取系统日志列表
func (c *systemLogController) List(ctx *gin.Context) {
	req := systemDTO.QuerySystemLogReq{}
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
