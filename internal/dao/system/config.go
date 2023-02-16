/*应用配置表*/
package systemDAO

import (
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"

	"gorm.io/gorm"
)

// Config 配置接口
type Config interface {
	All() ([]systemModel.Config, int64, error)
	List(req systemDTO.QueryConfigReq) ([]systemModel.Config, int64, error)
	Add(bean systemModel.Config) (uint, error)
	Update(bean systemModel.Config) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	Children(parentId uint) ([]systemModel.Config, error)
}

// 配置
type config struct {
	db *gorm.DB
}

// 创建配置对象
func NewConfigDao() *config {
	return &config{
		db: database.Instance(),
	}
}

// All 获取所有配置列表
func (d *config) All() ([]systemModel.Config, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db
		return stats
	}

	beans := make([]systemModel.Config, 0)
	if result := stats().Order("sort ASC").Order("id ASC").Find(&beans); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.Config{}).Count(&total)
	return beans, total, nil
}

// List 查询配置列表
func (d *config) List(req systemDTO.QueryConfigReq) ([]systemModel.Config, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db
		if req.Name != "" {
			stats = stats.Where("name like ?", "%"+req.Name+"%")
		}
		if req.Key != "" {
			stats = stats.Where("key like ?", "%"+req.Key+"%")
		}
		return stats
	}

	beans := make([]systemModel.Config, 0)
	result := stats().
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.Config{}).Count(&total)
	return beans, total, nil
}

// Add 添加配置
func (d *config) Add(bean systemModel.Config) (uint, error) {
	result := d.db.Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新配置
func (d *config) Update(bean systemModel.Config) (int64, error) {
	result := d.db.Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除配置
func (d *config) Delete(id uint) (int64, error) {
	result := d.db.Delete(&systemModel.Config{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除配置
func (d *config) BatchDelete(ids []uint) (int64, error) {
	beans := make([]systemModel.Config, len(ids))
	for _, id := range ids {
		beans = append(beans, systemModel.Config{
			ID: id,
		})
	}
	result := d.db.Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *config) Status(id uint, status uint) (int64, error) {
	result := d.db.Select("status").Updates(&systemModel.Config{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// Children 通过父 ID 获取子配置列表
func (d *config) Children(parentId uint) ([]systemModel.Config, error) {
	beans := make([]systemModel.Config, 0)
	result := d.db.Where("parent_id=?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}