/**文件上传*/
package system

import (
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	service "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 上传
type uploadController struct {
	service service.UploadService
}

// NewUploadController 创建上传对象
func NewUploadController() *uploadController {
	return &uploadController{
		service: service.NewUploadService(),
	}
}

// All 获取所有角色列表
func (c *uploadController) Avatar(ctx *gin.Context) {
	// 单文件
	file, err := ctx.FormFile("file")
	if err != nil {
		log.New(ctx).WithCode(statuscode.UploadFileParserError).Errorf("%v", err)
		response.New().WithCode(statuscode.UploadFileParserError).Json(ctx)
		return
	}

	c.service.Avatar(ctx, file).Json(ctx)
}
