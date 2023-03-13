/*网络请求日志
 */
package log

import (
	logDTO "gin-admin/internal/dto/log"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	logService "gin-admin/internal/service/log"

	"github.com/gin-gonic/gin"
)

// 网络请求日志
type httpLogController struct {
	service logService.HttpLogService
}

// NewHttpLogController 创建网络请求日志对象
func NewHttpLogController() *httpLogController {
	return &httpLogController{
		service: logService.NewHttpLogService(),
	}
}

// List 获取网络请求日志列表
func (c *httpLogController) List(ctx *gin.Context) {
	req := logDTO.QueryHttpLogReq{}
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

// GetBody 获取网络请求日志的 body 信息
// 由于该信息过长，单独获取
func (c *httpLogController) GetBody(ctx *gin.Context) {
	req := logDTO.QueryHttpLogBodyReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	result, err := c.service.GetBody(ctx, req.Id)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}
