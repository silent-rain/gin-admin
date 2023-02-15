/**文件上传*/
package system

import (
	"gin-admin/internal/pkg/code_errors"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 上传
type uploadController struct {
	service systemService.UploadService
}

// NewUploadController 创建上传对象
func NewUploadController() *uploadController {
	return &uploadController{
		service: systemService.NewUploadService(),
	}
}

// All 获取所有角色列表
func (c *uploadController) Avatar(ctx *gin.Context) {
	// 单文件
	file, err := ctx.FormFile("file")
	if err != nil {
		log.New(ctx).WithCode(code_errors.UploadFileParserError).Errorf("%v", err)
		response.New(ctx).WithCode(code_errors.UploadFileParserError).Json()
		return
	}

	result, err := c.service.Avatar(ctx, file)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}
