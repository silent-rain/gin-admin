// Package dao WEB 日志
package dao

import (
	"github.com/silent-rain/gin-admin/internal/app/log/dto"
	"github.com/silent-rain/gin-admin/internal/app/log/model"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// WebLog 系统日志接口
type WebLog interface {
	List(req dto.QueryWebLogReq) ([]model.WebLog, int64, error)
	Add(bean model.WebLog) (uint, error)
}

// WEB 日志
type webLog struct {
	mysql.DBRepo
}

// NewWebLogDao 创建 WEB 日志对象
func NewWebLogDao() *webLog {
	return &webLog{
		DBRepo: mysql.Instance(),
	}
}

// List 查询 WEB 日志列表
func (d *webLog) List(req dto.QueryWebLogReq) ([]model.WebLog, int64, error) {
	tx := d.GetDbR()
	if req.Nickname != "" {
		tx = tx.Where("nickname = ?", req.Nickname)
	}
	if req.OsType != 0 {
		tx = tx.Where("os_type = ?", req.OsType)
	}
	if req.ErrorType != 0 {
		tx = tx.Where("error_type = ?", req.ErrorType)
	}
	if req.Level != "" {
		tx = tx.Where("level = ?", req.Level)
	}
	if req.Url != "" {
		tx = tx.Where("url LIKE ?", req.Url+"%")
	}
	if req.Msg != "" {
		tx = tx.Where("msg LIKE ?", req.Msg+"%")
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&model.WebLog{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	beans := make([]model.WebLog, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).
		Order("created_at DESC").
		Find(&beans)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return beans, total, nil
}

// Add 添加 WEB 日志
func (d *webLog) Add(bean model.WebLog) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}
