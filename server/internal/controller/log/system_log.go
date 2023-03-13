/*系统日志
 */
package log

import (
	logDTO "github.com/silent-rain/gin-admin/internal/dto/log"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	logService "github.com/silent-rain/gin-admin/internal/service/log"

	"github.com/gin-gonic/gin"
)

// 系统日志
type systemLogController struct {
	service logService.SystemLogService
}

// NewSystemLogController 创建系统日志对象
func NewSystemLogController() *systemLogController {
	return &systemLogController{
		service: logService.NewSystemLogService(),
	}
}

// List 获取系统日志列表
func (c *systemLogController) List(ctx *gin.Context) {
	req := logDTO.QuerySystemLogReq{}
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
