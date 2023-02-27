/* 菜单 DAO
 */
package systemDAO

import (
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Menu 菜单接口
type Menu interface {
	All() ([]systemModel.Menu, int64, error)
	List(req systemDTO.QueryMenuReq) ([]systemModel.Menu, int64, error)
	Add(bean systemModel.Menu) (uint, error)
	Update(bean systemModel.Menu) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	ListByRoleIds(roleIds []uint) ([]systemModel.Menu, error)
	ChildrenMenu(parentId uint) ([]systemModel.Menu, error)
}

// 菜单
type menu struct {
	db mysql.DBRepo
}

// 创建菜单 Dao 对象
func NewMenuDao() *menu {
	return &menu{
		db: mysql.Instance(),
	}
}

// All 获取所有菜单列表
func (d *menu) All() ([]systemModel.Menu, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		return stats
	}

	bean := make([]systemModel.Menu, 0)
	if result := stats().Order("sort ASC").Order("id ASC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.Menu{}).Count(&total)
	return bean, total, nil
}

// List 查询菜单列表
func (d *menu) List(req systemDTO.QueryMenuReq) ([]systemModel.Menu, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		if req.Title != "" {
			stats = stats.Where("title like ?", "%"+req.Title+"%")
		}
		return stats
	}

	bean := make([]systemModel.Menu, 0)
	result := stats().
		Order("sort ASC").Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.Menu{}).Count(&total)
	return bean, total, nil
}

// Add 添加菜单
func (d *menu) Add(bean systemModel.Menu) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新菜单
func (d *menu) Update(bean systemModel.Menu) (int64, error) {
	result := d.db.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除菜单
func (d *menu) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&systemModel.Menu{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除菜单
func (d *menu) BatchDelete(ids []uint) (int64, error) {
	beans := make([]systemModel.Menu, len(ids))
	for _, id := range ids {
		beans = append(beans, systemModel.Menu{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *menu) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&systemModel.Menu{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// 通过 role_ids 获取菜单列表, 菜单去重
func (d *menu) ListByRoleIds(roleIds []uint) ([]systemModel.Menu, error) {
	beans := make([]systemModel.Menu, 0)
	result := d.db.GetDbR().Model(&systemModel.Menu{}).
		Joins("left join sys_role_menu_rel on sys_role_menu_rel.menu_id = sys_menu.id").
		Where("sys_role_menu_rel.role_id in ?", roleIds).
		Where("sys_menu.status = 1").
		Order("sort ASC").Order("id ASC").
		Distinct("sys_menu.*").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}

// ChildrenMenu 通过父 ID 获取子菜单列表
func (d *menu) ChildrenMenu(parentId uint) ([]systemModel.Menu, error) {
	bean := make([]systemModel.Menu, 0)
	result := d.db.GetDbR().Where("status=1").Where("parent_id=?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, result.Error
	}
	return bean, nil
}
