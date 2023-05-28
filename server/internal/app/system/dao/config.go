// Package dao 应用配置表
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Config 配置接口
type Config interface {
	All() ([]model.Config, int64, error)
	List(req dto.QueryConfigReq) ([]model.Config, int64, error)
	Info(key string) (model.Config, bool, error)
	Add(bean model.Config) (uint, error)
	Update(bean model.Config) (int64, error)
	BatchUpdate(beans []model.Config) error
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	Children(parentId uint) ([]model.Config, error)
	ChildrenByKey(key string) ([]model.Config, error)
}

// 配置
type config struct {
	mysql.DBRepo
}

// NewConfigDao 创建配置对象
func NewConfigDao() *config {
	return &config{
		DBRepo: mysql.Instance(),
	}
}

// All 获取所有配置列表
func (d *config) All() ([]model.Config, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		return stats
	}

	beans := make([]model.Config, 0)
	if result := stats().Order("sort ASC").Order("id ASC").Find(&beans); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.Config{}).Count(&total)
	return beans, total, nil
}

// List 查询配置列表
func (d *config) List(req dto.QueryConfigReq) ([]model.Config, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		if req.Name != "" {
			stats = stats.Where("name like ?", "%"+req.Name+"%")
		}
		if req.Key != "" {
			stats = stats.Where("key like ?", "%"+req.Key+"%")
		}
		return stats
	}

	beans := make([]model.Config, 0)
	result := stats().
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.Config{}).Count(&total)
	return beans, total, nil
}

// Info 获取配置信息
func (d *config) Info(key string) (model.Config, bool, error) {
	bean := model.Config{}
	result := d.GetDbR().Where("status=1").Where("`key` = ?", key).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加配置
func (d *config) Add(bean model.Config) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新配置
func (d *config) Update(bean model.Config) (int64, error) {
	result := d.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// BatchUpdate 批量更新配置
func (d *config) BatchUpdate(beans []model.Config) error {
	tx := d.GetDbW().Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			zap.S().Panic("注批量更新配置异常, err: %v", err)
		}
	}()

	for _, bean := range beans {
		bean := bean
		result := d.GetDbW().Omit("created_at").UpdateColumns(&bean)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}

// Delete 删除配置
func (d *config) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.Config{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除配置
func (d *config) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.Config, len(ids))
	for _, id := range ids {
		beans = append(beans, model.Config{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *config) Status(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.Config{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// Children 通过父 ID 获取子配置列表
func (d *config) Children(parentId uint) ([]model.Config, error) {
	beans := make([]model.Config, 0)
	result := d.GetDbR().Where("status=1").Where("parent_id=?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}

// ChildrenByKey 通过父 key 获取子配置列表
func (d *config) ChildrenByKey(key string) ([]model.Config, error) {
	beans := make([]model.Config, 0)
	subQuery := d.GetDbR().Model(&model.Config{}).Where("status=1").Where("`key` = ?", key).Select("id")
	result := d.GetDbR().Model(&model.Config{}).Where("status=1").Where("parent_id = (?)", subQuery).
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}