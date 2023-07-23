// Package service 系统日志
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/log/dao"
	"github.com/silent-rain/gin-admin/internal/app/log/dto"
	"github.com/silent-rain/gin-admin/internal/app/log/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// SystemLogService 系统日志
type SystemLogService struct {
	dao *dao.SystemLog
}

// NewSystemLogService 创建系统日志对象
func NewSystemLogService() *SystemLogService {
	return &SystemLogService{
		dao: dao.NewSystemLogDao(),
	}
}

// List 获取系统日志列表
func (s *SystemLogService) List(ctx *gin.Context, req dto.QuerySystemLogReq) ([]model.SystemLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}
