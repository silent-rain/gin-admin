// Package dao 角色
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/global"
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Role 角色
type Role struct {
	mysql.DBRepo
}

// NewRoleDao 创建角色 Dao 对象
func NewRoleDao() *Role {
	return &Role{
		DBRepo: global.Instance().Mysql(),
	}
}

// All 获取所有角色列表
func (d *Role) All() ([]model.Role, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		return stats
	}

	bean := make([]model.Role, 0)
	if result := stats().Order("updated_at DESC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.Role{}).Count(&total)
	return bean, total, nil
}

// List 查询角色列表
func (d *Role) List(req dto.QueryRoleReq) ([]model.Role, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		if req.Name != "" {
			stats = stats.Where("name like ?", "%"+req.Name+"%")
		}
		return stats
	}

	bean := make([]model.Role, 0)
	result := stats().Offset(req.Offset()).Limit(req.PageSize).
		Order("sort DESC").Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.Role{}).Count(&total)
	return bean, total, nil
}

// InfoByName 获取角色信息
func (d *Role) InfoByName(name string) (model.Role, bool, error) {
	bean := model.Role{}
	result := d.GetDbR().Where("name=?", name).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Role{}, false, nil
	}
	if result.Error != nil {
		return model.Role{}, false, result.Error
	}
	return bean, true, nil
}

// Add 添加角色
func (d *Role) Add(bean model.Role) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新角色
func (d *Role) Update(bean model.Role) (int64, error) {
	result := d.GetDbW().Select("name", "status", "sort", "note").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除角色
func (d *Role) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.Role{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除角色
func (d *Role) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.Role, len(ids))
	for _, id := range ids {
		beans = append(beans, model.Role{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// UpdateStatus 更新状态
func (d *Role) UpdateStatus(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.Role{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
