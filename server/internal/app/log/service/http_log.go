// Package service 网络请求日志
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/log/dao"
	"github.com/silent-rain/gin-admin/internal/app/log/dto"
	"github.com/silent-rain/gin-admin/internal/app/log/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// HttpLogService 网络请求日志
type HttpLogService struct {
	dao *dao.HttpLog
}

// NewHttpLogService 创建网络请求日志对象
func NewHttpLogService() *HttpLogService {
	return &HttpLogService{
		dao: dao.NewHttpLogDao(),
	}
}

// List 获取网络请求日志列表
func (s *HttpLogService) List(ctx *gin.Context, req dto.QueryHttpLogReq) ([]model.HttpLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}

// GetBody 获取 body 信息
func (h *HttpLogService) GetBody(ctx *gin.Context, id uint) (dto.QueryHttpLogBody, error) {
	result := dto.QueryHttpLogBody{
		Body: "",
	}
	resp, ok, err := h.dao.Info(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return result, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Errorf("%v", err)
		return result, errcode.DBQueryEmptyError
	}

	result.Body = resp.Body
	return result, nil
}
