/*应用配置表*/
package systemDAO

import (
	"errors"
	DAO "gin-admin/internal/dao"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/repository/mysql"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Config 配置接口
type Config interface {
	All() ([]systemModel.Config, int64, error)
	List(req systemDTO.QueryConfigReq) ([]systemModel.Config, int64, error)
	InfoByKey(key string) (systemModel.Config, bool, error)
	Add(bean systemModel.Config) (uint, error)
	Update(bean systemModel.Config) (int64, error)
	BatchUpdate(beans []systemModel.Config) error
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	Children(parentId uint) ([]systemModel.Config, error)
	ChildrenByKey(key string) ([]systemModel.Config, error)
}

// 配置
type config struct {
	*DAO.Transaction
	db mysql.DBRepo
}

// 创建配置对象
func NewConfigDao() *config {
	return &config{
		Transaction: DAO.NewTransaction(mysql.Instance().GetDbW()),
		db:          mysql.Instance(),
	}
}

// All 获取所有配置列表
func (d *config) All() ([]systemModel.Config, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
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
		stats := d.db.GetDbR().Debug()
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

// Info 获取配置信息
func (d *config) InfoByKey(key string) (systemModel.Config, bool, error) {
	bean := systemModel.Config{}
	result := d.db.GetDbR().Where("`key` = ?", key).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加配置
func (d *config) Add(bean systemModel.Config) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新配置
func (d *config) Update(bean systemModel.Config) (int64, error) {
	result := d.db.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// BatchUpdate 批量更新配置
func (d *config) BatchUpdate(beans []systemModel.Config) error {
	d.Begin()
	defer func() {
		if err := recover(); err != nil {
			d.Rollback()
			zap.S().Panic("注批量更新配置异常, err: %v", err)
		}
	}()

	for _, bean := range beans {
		bean := bean
		result := d.db.GetDbW().Omit("created_at").UpdateColumns(&bean)
		if result.Error != nil {
			d.Rollback()
			return result.Error
		}
	}
	d.Commit()
	return nil
}

// Delete 删除配置
func (d *config) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&systemModel.Config{
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
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *config) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&systemModel.Config{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// Children 通过父 ID 获取子配置列表
func (d *config) Children(parentId uint) ([]systemModel.Config, error) {
	beans := make([]systemModel.Config, 0)
	result := d.db.GetDbR().Where("parent_id=?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}

// ChildrenByKey 通过父 key 获取子配置列表
func (d *config) ChildrenByKey(key string) ([]systemModel.Config, error) {
	beans := make([]systemModel.Config, 0)
	subQuery := d.db.GetDbR().Model(&systemModel.Config{}).Where("`key` = ?", key).Select("id")
	result := d.db.GetDbR().Model(&systemModel.Config{}).Where("parent_id = (?)", subQuery).
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}
