/**文件上传*/
package system

import (
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"
	"gin-admin/pkg/errcode"

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

// Avatar 上传用户头像
func (c *uploadController) Avatar(ctx *gin.Context) {
	// 单文件
	file, err := ctx.FormFile("file")
	if err != nil {
		log.New(ctx).WithCode(errcode.UploadFileParserError).Errorf("%v", err)
		response.New(ctx).WithCode(errcode.UploadFileParserError).Json()
		return
	}

	result, err := c.service.Avatar(ctx, file)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}

// Image 上传图片
func (c *uploadController) Image(ctx *gin.Context) {
	// 单文件
	file, err := ctx.FormFile("file")
	if err != nil {
		log.New(ctx).WithCode(errcode.UploadFileParserError).Errorf("%v", err)
		response.New(ctx).WithCode(errcode.UploadFileParserError).Json()
		return
	}

	result, err := c.service.Image(ctx, file)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}

// Images 上传图片列表
func (c *uploadController) Images(ctx *gin.Context) {
	// 多文件
	form, err := ctx.MultipartForm()
	if err != nil {
		log.New(ctx).WithCode(errcode.UploadFileParserError).Errorf("%v", err)
		response.New(ctx).WithCode(errcode.UploadFileParserError).Json()
		return
	}
	files := form.File["file"]

	result, err := c.service.Images(ctx, files)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(result, int64(len(result))).Json()
}
