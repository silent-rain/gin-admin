// Package dao 网络请求日志
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/global"
	"github.com/silent-rain/gin-admin/internal/app/log/dto"
	"github.com/silent-rain/gin-admin/internal/app/log/model"
	"github.com/silent-rain/gin-admin/pkg/repository/mysql"

	"gorm.io/gorm"
)

// HttpLog 网络请求日志结构
type HttpLog struct {
	mysql.DBRepo
}

// NewHttpLogDao 创建网络请求日志 Dao 对象
func NewHttpLogDao() *HttpLog {
	return &HttpLog{
		DBRepo: global.Instance().Mysql(),
	}
}

// List 查询网络请求日志列表
func (d *HttpLog) List(req dto.QueryHttpLogReq) ([]model.HttpLog, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR().Omit("body")
		if req.UserId != 0 {
			stats = stats.Where("user_id = ?", req.UserId)
		}
		if req.TraceId != "" {
			stats = stats.Where("trace_id = ? OR parent_trace_id = ?", req.TraceId, req.TraceId)
		}
		if req.StatusCode != 0 {
			stats = stats.Where("status_code = ?", req.StatusCode)
		}
		if req.Method != "" {
			stats = stats.Where("method = ?", req.Method)
		}
		if req.RemoteAddr != "" {
			stats = stats.Where("remote_addr = ?", req.RemoteAddr)
		}
		if req.HttpType != "" {
			stats = stats.Where("htpp_type = ?", req.HttpType)
		}
		if req.Path != "" {
			stats = stats.Where("path LIKE ?", "%"+req.Path+"%")
		}
		return stats
	}

	beans := make([]model.HttpLog, 0)
	result := stats().Offset(req.Offset()).Limit(req.PageSize).
		Order("created_at DESC").
		Find(&beans)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.HttpLog{}).Count(&total)
	return beans, total, nil
}

// Add 添加网络请求日志
func (d *HttpLog) Add(bean model.HttpLog) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Info 获取详情信息
func (d *HttpLog) Info(id uint) (model.HttpLog, bool, error) {
	bean := model.HttpLog{
		ID: id,
	}
	result := d.GetDbR().First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}
