// Package service 文件上传
package service

import (
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/conf"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"
	"github.com/silent-rain/gin-admin/pkg/utils"

	"github.com/gin-gonic/gin"
)

var (
	// 头像文件夹名称
	avatarDirName = "avatar"
	// 图片文件夹名称
	imagesDirName = "images"
)

// UploadService 上传
type UploadService interface {
	Avatar(ctx *gin.Context, file *multipart.FileHeader) (dto.Image, error)
	Image(ctx *gin.Context, file *multipart.FileHeader) (dto.Image, error)
	Images(ctx *gin.Context, files []*multipart.FileHeader) ([]dto.Image, error)
}

// 上传
type uploadService struct {
}

// NewUploadService 创建上传对象
func NewUploadService() *uploadService {
	return &uploadService{}
}

// Avatar 上传用户头像
func (h *uploadService) Avatar(ctx *gin.Context, file *multipart.FileHeader) (dto.Image, error) {
	result := dto.Image{}

	// 文件夹是否存在, 不存在则创建
	dstPath := filepath.Join(conf.Instance().Server.Upload.FilePath, avatarDirName)
	if err := utils.DirNotExistCreate(dstPath); err != nil {
		log.New(ctx).WithCode(errcode.DirCreateError).Errorf("%v", err)
		return result, errcode.DirCreateError.
			WithMsg(fmt.Sprintf("err: %v", err))
	}

	// 上传文件到指定位置
	ext := path.Ext(file.Filename)
	filename := utils.Md5(file.Filename+strconv.Itoa(int(file.Size))+time.Now().Local().String()) + ext
	dst := filepath.Join(dstPath, filename)

	// 文件夹不存在则创建
	if err := utils.DirNotExistCreate(dstPath); err != nil {
		log.New(ctx).WithCode(errcode.DirCreateError).Errorf("%v", err)
		return result, errcode.DirCreateError.
			WithMsg(fmt.Sprintf("err: %v", err))
	}

	// 保存文件
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		log.New(ctx).WithCode(errcode.UploadFileSaveError).Errorf("%v", err)
		return result, errcode.UploadFileSaveError
	}

	result.Name = file.Filename
	result.Url = "/" + dst
	return result, nil
}

// Image 上传图片
func (h *uploadService) Image(ctx *gin.Context, file *multipart.FileHeader) (dto.Image, error) {
	result := dto.Image{}

	// 文件夹是否存在, 不存在则创建
	timePath := time.Now().Format("2006-01-02")
	dstPath := filepath.Join(conf.Instance().Server.Upload.FilePath, imagesDirName, timePath)
	if err := utils.DirNotExistCreate(dstPath); err != nil {
		log.New(ctx).WithCode(errcode.DirCreateError).Errorf("%v", err)
		return result, errcode.DirCreateError.
			WithMsg(fmt.Sprintf("err: %v", err))
	}

	// 上传文件到指定位置
	ext := path.Ext(file.Filename)
	filename := utils.Md5(timePath+file.Filename+strconv.Itoa(int(file.Size))+time.Now().Local().String()) + ext
	dst := filepath.Join(dstPath, filename)

	// 保存文件
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		log.New(ctx).WithCode(errcode.UploadFileSaveError).Errorf("%v", err)
		return result, errcode.UploadFileSaveError
	}

	result.Name = file.Filename
	result.Url = "/" + dst
	return result, nil
}

// Images 上传图片列表
func (h *uploadService) Images(ctx *gin.Context, files []*multipart.FileHeader) ([]dto.Image, error) {
	results := make([]dto.Image, 0)

	timePath := time.Now().Format("2006-01-02")
	dstPath := filepath.Join(conf.Instance().Server.Upload.FilePath, imagesDirName, timePath)
	// 文件夹是否存在, 不存在则创建
	if err := utils.DirNotExistCreate(dstPath); err != nil {
		log.New(ctx).WithCode(errcode.DirCreateError).Errorf("%v", err)
		return nil, errcode.DirCreateError.
			WithMsg(fmt.Sprintf("err: %v", err))
	}

	for _, file := range files {
		// 上传文件到指定位置
		ext := path.Ext(file.Filename)
		filename := utils.Md5(timePath+file.Filename+strconv.Itoa(int(file.Size))+time.Now().Local().String()) + ext
		dst := filepath.Join(dstPath, filename)

		// 保存文件
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			log.New(ctx).WithCode(errcode.UploadFileSaveError).Errorf("%v", err)
			return nil, errcode.UploadFileSaveError
		}

		results = append(results, dto.Image{
			Name: file.Filename,
			Url:  "/" + dst,
		})
	}

	return results, nil
}
