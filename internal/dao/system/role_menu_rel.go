/* 角色菜单 DAO
 */
package systemDAO

import (
	DAO "gin-admin/internal/dao"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"
	"gin-admin/internal/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// RoleMenuRel 角色菜单关系接口
type RoleMenuRel interface {
	List(req systemDTO.QueryRoleMenuRelReq) ([]systemModel.RoleMenuRel, int64, error)
	Update(roleId uint, menuIds []uint) error
}

// 角色菜单关系
type roleMenuRel struct {
	*DAO.Transaction
	db *gorm.DB
}

// 创建角色菜单关系 Dao 对象
func NewRoleMenuRelDao() *roleMenuRel {
	return &roleMenuRel{
		Transaction: DAO.NewTransaction(database.Instance()),
		db:          database.Instance(),
	}
}

// List 角色关联的菜单列表
func (d *roleMenuRel) List(req systemDTO.QueryRoleMenuRelReq) ([]systemModel.RoleMenuRel, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db
		if req.MenuId != 0 {
			stats = stats.Where("menu_id = ?", req.MenuId)
		}
		if req.RoleId != 0 {
			stats = stats.Where("role_id = ?", req.RoleId)
		}
		return stats
	}

	bean := make([]systemModel.RoleMenuRel, 0)
	result := stats().
		Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.RoleMenuRel{}).Count(&total)
	return bean, total, nil
}

// Update 更新角色关联的菜单
func (d *roleMenuRel) Update(roleId uint, menuIds []uint) error {
	d.Begin()
	defer func() {
		if err := recover(); err != nil {
			d.Rollback()
			zap.S().Panic("更新角色关联关系异常, err: %v", err)
		}
	}()

	// 未传入 menu_ids, 不做处理
	if menuIds == nil {
		return nil
	}
	// 获取角色关联菜单的 menuId 列表
	roleMenuIds, err := d.getRoleMenuRelByMenuIds(roleId)
	if err != nil {
		return err
	}
	// 新增的 menuId 列表
	addRoleMenuRels := make([]systemModel.RoleMenuRel, 0)
	for _, menuId := range menuIds {
		if utils.IndexOfArray(roleMenuIds, menuId) == -1 {
			addRoleMenuRels = append(addRoleMenuRels, systemModel.RoleMenuRel{
				RoleId: roleId,
				MenuId: menuId,
			})
		}
	}

	// 删除的 menuId 列表
	deleteRoleMenuIds := make([]uint, 0)
	for _, roleId := range roleMenuIds {
		if utils.IndexOfArray(menuIds, roleId) == -1 {
			deleteRoleMenuIds = append(deleteRoleMenuIds, roleId)
		}
	}

	if len(addRoleMenuRels) != 0 {
		if result := d.Tx().Create(&addRoleMenuRels); result.Error != nil {
			return result.Error
		}
	}
	if len(deleteRoleMenuIds) != 0 {
		if result := d.Tx().Where("role_id = ? AND menu_id in ?", roleId, deleteRoleMenuIds).
			Delete(&systemModel.RoleMenuRel{}); result.Error != nil {
			return result.Error
		}
	}
	d.Commit()
	return nil
}

// 获取角色关联的菜单 menuId 列表
func (d *roleMenuRel) getRoleMenuRelByMenuIds(roleId uint) ([]uint, error) {
	beans := make([]systemModel.RoleMenuRel, 0)
	results := d.Tx().Where("role_id = ?", roleId).Find(&beans)
	if results.Error != nil {
		return nil, results.Error
	}

	roleMenuRelIds := make([]uint, 0)
	for _, item := range beans {
		roleMenuRelIds = append(roleMenuRelIds, item.MenuId)
	}
	return roleMenuRelIds, nil
}
