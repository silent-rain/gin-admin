// Package dao 角色菜单
package dao

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"
	"github.com/silent-rain/gin-admin/pkg/slices"

	"gorm.io/gorm"
)

// RoleMenuRel 角色菜单关系
type RoleMenuRel struct {
	mysql.DBRepo
}

// NewRoleMenuRelDao 创建角色菜单关系 Dao 对象
func NewRoleMenuRelDao() *RoleMenuRel {
	return &RoleMenuRel{
		DBRepo: global.Instance().Mysql(),
	}
}

// List 角色关联的菜单列表
func (d *RoleMenuRel) List(req dto.QueryRoleMenuRelReq) ([]model.RoleMenuRel, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		if req.MenuId != 0 {
			stats = stats.Where("menu_id = ?", req.MenuId)
		}
		if req.RoleId != 0 {
			stats = stats.Where("role_id = ?", req.RoleId)
		}
		return stats
	}

	bean := make([]model.RoleMenuRel, 0)
	result := stats().Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.RoleMenuRel{}).Count(&total)
	return bean, total, nil
}

// Update 更新角色关联的菜单
func (d *RoleMenuRel) Update(roleId uint, menuIds []uint) error {
	// 未传入 menu_ids, 不做处理
	if menuIds == nil {
		return nil
	}
	// 获取角色关联菜单的 menuId 列表
	roleMenuIds, err := d.getRoleMenuRelByMenuIds(roleId)
	if err != nil {
		return err
	}

	tx := d.GetDbW().Begin()

	// 批量删除关系
	if err := d.deleteRoleMenuIds(tx, menuIds, roleMenuIds, roleId); err != nil {
		tx.Rollback()
		return err
	}
	// 批量添加关系
	if err := d.addRoleMenuRels(tx, menuIds, roleMenuIds, roleId); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 获取角色关联的菜单 menuId 列表
func (d *RoleMenuRel) getRoleMenuRelByMenuIds(roleId uint) ([]uint, error) {
	beans := make([]model.RoleMenuRel, 0)
	results := d.GetDbR().Where("role_id = ?", roleId).Find(&beans)
	if results.Error != nil {
		return nil, results.Error
	}

	roleMenuRelIds := make([]uint, 0)
	for _, item := range beans {
		roleMenuRelIds = append(roleMenuRelIds, item.MenuId)
	}
	return roleMenuRelIds, nil
}

// 批量添加关系
func (d *RoleMenuRel) addRoleMenuRels(tx *gorm.DB, menuIds, roleMenuIds []uint, roleId uint) error {
	// 新增的 menuId 列表
	addRoleMenuRels := make([]model.RoleMenuRel, 0)
	for _, menuId := range menuIds {
		if slices.IndexOfArray(roleMenuIds, menuId) == -1 {
			addRoleMenuRels = append(addRoleMenuRels, model.RoleMenuRel{
				RoleId: roleId,
				MenuId: menuId,
			})
		}
	}

	if len(addRoleMenuRels) == 0 {
		return nil
	}
	if result := tx.Create(&addRoleMenuRels); result.Error != nil {
		return result.Error
	}
	return nil
}

// 批量删除关系
func (d *RoleMenuRel) deleteRoleMenuIds(tx *gorm.DB, menuIds, roleMenuIds []uint, roleId uint) error {
	// 删除的 menuId 列表
	deleteRoleMenuIds := make([]uint, 0)
	for _, roleId := range roleMenuIds {
		if slices.IndexOfArray(menuIds, roleId) == -1 {
			deleteRoleMenuIds = append(deleteRoleMenuIds, roleId)
		}
	}

	if len(deleteRoleMenuIds) == 0 {
		return nil
	}
	if result := tx.Where("role_id = ? AND menu_id in ?", roleId, deleteRoleMenuIds).
		Delete(&model.RoleMenuRel{}); result.Error != nil {
		return result.Error
	}
	return nil
}
