/*网络请求日志
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 网络请求日志
type httpLogController struct {
	service systemService.HttpLogService
}

// NewHttpLogController 创建网络请求日志对象
func NewHttpLogController() *httpLogController {
	return &httpLogController{
		service: systemService.NewHttpLogService(),
	}
}

// List 获取网络请求日志列表
func (c *httpLogController) List(ctx *gin.Context) {
	req := systemDTO.QueryHttpLogReq{}
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
