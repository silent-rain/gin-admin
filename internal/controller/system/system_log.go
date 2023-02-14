/*系统日志
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	service "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 系统日志
type systemLogController struct {
	service service.SystemLogService
}

// NewSystemLogController 创建系统日志对象
func NewSystemLogController() *systemLogController {
	return &systemLogController{
		service: service.NewSystemLogService(),
	}
}

// List 获取系统日志列表
func (c *systemLogController) List(ctx *gin.Context) {
	req := systemDTO.QuerySystemLogReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.List(ctx, req).Json(ctx)
}
