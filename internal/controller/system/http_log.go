/*网络请求日志
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	service "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 网络请求日志
type httpLogController struct {
	service service.HttpLogService
}

// NewHttpLogController 创建网络请求日志对象
func NewHttpLogController() *httpLogController {
	return &httpLogController{
		service: service.NewHttpLogService(),
	}
}

// List 获取网络请求日志列表
func (c *httpLogController) List(ctx *gin.Context) {
	req := systemDTO.QueryHttpLogReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.List(ctx, req).Json(ctx)
}
