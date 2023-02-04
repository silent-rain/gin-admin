/*
 * @Author: silent-rain
 * @Date: 2023-01-09 23:08:33
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-12 00:06:00
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/system_log.go
 * @Descripttion: 系统日志
 */
package systemDao

import (
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"

	"gorm.io/gorm"
)

// HttpLog 系统日志接口
type SystemLog interface {
	Add(bean systemModel.SystemLog) (uint, error)
}

// 系统日志
type systemLog struct {
	db *gorm.DB
}

// 创建系统日志 Dao 对象
func NewDaoSystemLog() *systemLog {
	return &systemLog{
		db: database.Instance(),
	}
}

// Add 添加系统日志
func (d *systemLog) Add(bean systemModel.SystemLog) (uint, error) {
	result := d.db.Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}
