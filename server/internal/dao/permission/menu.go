/* 菜单 DAO
 */
package permission

import (
	permissionDTO "github.com/silent-rain/gin-admin/internal/dto/permission"
	permissionModel "github.com/silent-rain/gin-admin/internal/model/permission"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Menu 菜单接口
type Menu interface {
	All() ([]permissionModel.Menu, int64, error)
	List(req permissionDTO.QueryMenuReq) ([]permissionModel.Menu, int64, error)
	Add(bean permissionModel.Menu) (uint, error)
	Update(bean permissionModel.Menu) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	ListByRoleIds(roleIds []uint) ([]permissionModel.Menu, error)
	ChildrenMenu(parentId uint) ([]permissionModel.Menu, error)
}

// 菜单
type menu struct {
	db mysql.DBRepo
}

// NewMenuDao 创建菜单 Dao 对象
func NewMenuDao() *menu {
	return &menu{
		db: mysql.Instance(),
	}
}

// All 获取所有菜单列表
func (d *menu) All() ([]permissionModel.Menu, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		return stats
	}

	bean := make([]permissionModel.Menu, 0)
	if result := stats().Order("sort ASC").Order("id ASC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&permissionModel.Menu{}).Count(&total)
	return bean, total, nil
}

// List 查询菜单列表
func (d *menu) List(req permissionDTO.QueryMenuReq) ([]permissionModel.Menu, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		if req.Title != "" {
			stats = stats.Where("title like ?", "%"+req.Title+"%")
		}
		return stats
	}

	bean := make([]permissionModel.Menu, 0)
	result := stats().
		Order("sort ASC").Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&permissionModel.Menu{}).Count(&total)
	return bean, total, nil
}

// Add 添加菜单
func (d *menu) Add(bean permissionModel.Menu) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新菜单
func (d *menu) Update(bean permissionModel.Menu) (int64, error) {
	result := d.db.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除菜单
func (d *menu) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&permissionModel.Menu{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除菜单
func (d *menu) BatchDelete(ids []uint) (int64, error) {
	beans := make([]permissionModel.Menu, len(ids))
	for _, id := range ids {
		beans = append(beans, permissionModel.Menu{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *menu) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&permissionModel.Menu{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// 通过 role_ids 获取菜单列表, 菜单去重
func (d *menu) ListByRoleIds(roleIds []uint) ([]permissionModel.Menu, error) {
	beans := make([]permissionModel.Menu, 0)
	result := d.db.GetDbR().Model(&permissionModel.Menu{}).
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
func (d *menu) ChildrenMenu(parentId uint) ([]permissionModel.Menu, error) {
	bean := make([]permissionModel.Menu, 0)
	result := d.db.GetDbR().Where("status=1").Where("parent_id=?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, result.Error
	}
	return bean, nil
}
