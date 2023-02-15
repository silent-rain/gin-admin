/*文件上传*/
package service

import (
	"errors"
	"fmt"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"

	"gin-admin/internal/pkg/code_errors"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/utils"
	systemVO "gin-admin/internal/vo/system"

	"github.com/gin-gonic/gin"
)

// UploadService 上传
type UploadService interface {
	Avatar(ctx *gin.Context, file *multipart.FileHeader) (systemVO.Avatar, error)
}

// 上传
type uploadService struct {
}

// NewUploadService 创建上传对象
func NewUploadService() *uploadService {
	return &uploadService{}
}

// All 获取所有角色列表
func (h *uploadService) Avatar(ctx *gin.Context, file *multipart.FileHeader) (systemVO.Avatar, error) {
	result := systemVO.Avatar{
		Url: "",
	}

	ext := path.Ext(file.Filename)
	filename := utils.Md5(file.Filename+strconv.Itoa(int(file.Size))+time.Now().Local().String()) + ext
	// 上传文件到指定的 dst
	dst := conf.Instance().UploadConfig.FilePath + "/avatar/" + filename

	err := ctx.SaveUploadedFile(file, dst)
	if errors.Is(err, fs.ErrNotExist) {
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			log.New(ctx).WithCode(code_errors.DirNotFoundError).Errorf("%v", err)
			return result, code_errors.New(code_errors.DirNotFoundError).
				WithMsg(fmt.Sprintf("%s not found", dst))
		}
	}
	if err != nil {
		log.New(ctx).WithCode(code_errors.UploadFileSaveError).Errorf("%v", err)
		return result, code_errors.New(code_errors.UploadFileSaveError)
	}

	return result, nil
}
