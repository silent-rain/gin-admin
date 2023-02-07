/**文件上传*/
package system

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 上传
type uploadHandler struct {
}

// 创建上传 Handler 对象
func NewUploadHandler() *uploadHandler {
	return &uploadHandler{}
}

// All 获取所有角色列表
func (h *uploadHandler) Avatar(ctx *gin.Context) {
	// 单文件
	file, err := ctx.FormFile("file")
	if err != nil {
		log.New(ctx).WithCode(statuscode.UploadFileParserError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.UploadFileParserError).Json()
		return
	}

	zap.S().Errorf("============= %#v", file.Filename)

	// 上传文件到指定的 dst
	dst := conf.Instance().UploadConfig.FilePath + "/avatar/" + file.Filename
	err = ctx.SaveUploadedFile(file, dst)
	if errors.Is(err, fs.ErrNotExist) {
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			log.New(ctx).WithCode(statuscode.DirNotFoundError).Errorf("%v", err)
			response.New(ctx).WithCode(statuscode.DirNotFoundError).
				WithMsg(fmt.Sprintf("%s not found", dst)).Json()
			return
		}
	}
	if err != nil {
		zap.S().Errorf("============%#v", err)
		log.New(ctx).WithCode(statuscode.UploadFileSaveError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.UploadFileSaveError).Json()
		return
	}

	response.New(ctx).WithData(map[string]string{
		"url": dst,
	}).Json()
}
