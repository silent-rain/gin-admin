/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:24:36
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 17:16:27
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/role.go
 * @Descripttion: 角色
 */
package systemDao

import (
	"errors"
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"

	"gorm.io/gorm"
)

// Role 角色接口
type Role interface {
	All() ([]systemModel.Role, int64, error)
	List(req systemDto.QueryRoleReq) ([]systemModel.Role, int64, error)
	InfoByName(name string) (systemModel.Role, bool, error)
	Add(bean systemModel.Role) (uint, error)
	Update(bean systemModel.Role) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// 角色
type role struct {
	db *gorm.DB
}

// 创建角色 Dao 对象
func NewRoleDao() *role {
	return &role{
		db: database.Instance(),
	}
}

// All 获取所有角色列表
func (d *role) All() ([]systemModel.Role, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db
		return stats
	}

	bean := make([]systemModel.Role, 0)
	if result := stats().Order("updated_at DESC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.Role{}).Count(&total)
	return bean, total, nil
}

// List 查询角色列表
func (d *role) List(req systemDto.QueryRoleReq) ([]systemModel.Role, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db
		if req.Name != "" {
			stats = stats.Where("name like ?", "%"+req.Name+"%")
		}
		return stats
	}

	bean := make([]systemModel.Role, 0)
	if result := stats().Offset(req.Offset()).Limit(req.PageSize).Order("sort DESC").Order("updated_at DESC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.Role{}).Count(&total)
	return bean, total, nil
}

// Info 获取角色信息
func (d *role) InfoByName(name string) (systemModel.Role, bool, error) {
	bean := systemModel.Role{}
	result := d.db.Where("name=?", name).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return systemModel.Role{}, false, nil
	}
	if result.Error != nil {
		return systemModel.Role{}, false, result.Error
	}
	return bean, true, nil
}

// Add 添加角色
func (d *role) Add(bean systemModel.Role) (uint, error) {
	result := d.db.Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新角色
func (d *role) Update(bean systemModel.Role) (int64, error) {
	result := d.db.Select("name", "status", "sort", "note").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除角色
func (d *role) Delete(id uint) (int64, error) {
	result := d.db.Delete(&systemModel.Role{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除角色
func (d *role) BatchDelete(ids []uint) (int64, error) {
	beans := make([]systemModel.Role, len(ids))
	for _, id := range ids {
		beans = append(beans, systemModel.Role{
			ID: id,
		})
	}
	result := d.db.Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *role) Status(id uint, status uint) (int64, error) {
	result := d.db.Select("status").Updates(&systemModel.Role{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
