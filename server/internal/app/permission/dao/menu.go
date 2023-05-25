/* 菜单 DAO
 */
package dao

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Menu 菜单接口
type Menu interface {
	All() ([]model.Menu, int64, error)
	List(req dto.QueryMenuReq) ([]model.Menu, int64, error)
	Add(bean model.Menu) (uint, error)
	Update(bean model.Menu) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	ListByRoleIds(roleIds []uint) ([]model.Menu, error)
	ChildrenMenu(parentId uint) ([]model.Menu, error)
}

// 菜单
type menu struct {
	mysql.DBRepo
}

// NewMenuDao 创建菜单 Dao 对象
func NewMenuDao() *menu {
	return &menu{
		DBRepo: mysql.Instance(),
	}
}

// All 获取所有菜单列表
func (d *menu) All() ([]model.Menu, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		return stats
	}

	bean := make([]model.Menu, 0)
	if result := stats().Order("sort ASC").Order("id ASC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.Menu{}).Count(&total)
	return bean, total, nil
}

// List 查询菜单列表
func (d *menu) List(req dto.QueryMenuReq) ([]model.Menu, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		if req.Title != "" {
			stats = stats.Where("title like ?", "%"+req.Title+"%")
		}
		return stats
	}

	bean := make([]model.Menu, 0)
	result := stats().
		Order("sort ASC").Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.Menu{}).Count(&total)
	return bean, total, nil
}

// Add 添加菜单
func (d *menu) Add(bean model.Menu) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新菜单
func (d *menu) Update(bean model.Menu) (int64, error) {
	result := d.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除菜单
func (d *menu) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.Menu{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除菜单
func (d *menu) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.Menu, len(ids))
	for _, id := range ids {
		beans = append(beans, model.Menu{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *menu) Status(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.Menu{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// 通过 role_ids 获取菜单列表, 菜单去重
func (d *menu) ListByRoleIds(roleIds []uint) ([]model.Menu, error) {
	beans := make([]model.Menu, 0)
	result := d.GetDbR().Model(&model.Menu{}).
		Joins("left join perm_role_menu_rel on perm_role_menu_rel.menu_id = perm_menu.id").
		Where("perm_role_menu_rel.role_id in ?", roleIds).
		Where("perm_menu.status = 1").
		Order("sort ASC").Order("id ASC").
		Distinct("perm_menu.*").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}

// ChildrenMenu 通过父 ID 获取子菜单列表
func (d *menu) ChildrenMenu(parentId uint) ([]model.Menu, error) {
	bean := make([]model.Menu, 0)
	result := d.GetDbR().Where("status=1").Where("parent_id=?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, result.Error
	}
	return bean, nil
}
