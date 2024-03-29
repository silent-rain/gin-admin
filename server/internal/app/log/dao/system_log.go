// Package dao 系统日志
package dao

import (
	"github.com/silent-rain/gin-admin/global"
	"github.com/silent-rain/gin-admin/internal/app/log/dto"
	"github.com/silent-rain/gin-admin/internal/app/log/model"
	"github.com/silent-rain/gin-admin/pkg/repository/mysql"

	"gorm.io/gorm"
)

// SystemLog 系统日志
type SystemLog struct {
	mysql.DBRepo
}

// NewSystemLogDao 创建系统日志对象
func NewSystemLogDao() *SystemLog {
	return &SystemLog{
		DBRepo: global.Instance().Mysql(),
	}
}

// List 查询系统日志列表
func (d *SystemLog) List(req dto.QuerySystemLogReq) ([]model.SystemLog, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		if req.UserId != 0 {
			stats = stats.Where("user_id = ?", req.UserId)
		}
		if req.TraceId != "" {
			stats = stats.Where("trace_id = ?", req.TraceId)
		}
		if req.Level != "" {
			stats = stats.Where("level = ?", req.Level)
		}
		if req.ErrorCode != 0 {
			stats = stats.Where("error_code = ?", req.ErrorCode)
		}
		if req.ErrorMsg != "" {
			stats = stats.Where("error_msg LIKE ?", req.ErrorMsg+"%")
		}
		if req.Msg != "" {
			stats = stats.Where("msg LIKE ?", req.Msg+"%")
		}
		return stats
	}

	beans := make([]model.SystemLog, 0)
	result := stats().Offset(req.Offset()).Limit(req.PageSize).
		Order("created_at DESC").
		Find(&beans)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.SystemLog{}).Count(&total)
	return beans, total, nil
}

// Add 添加系统日志
func (d *SystemLog) Add(bean model.SystemLog) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}
