/*角色 DAO
 */
package permission

import (
	"errors"

	permissionDTO "gin-admin/internal/dto/permission"
	permissionModel "gin-admin/internal/model/permission"
	"gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Role 角色接口
type Role interface {
	All() ([]permissionModel.Role, int64, error)
	List(req permissionDTO.QueryRoleReq) ([]permissionModel.Role, int64, error)
	InfoByName(name string) (permissionModel.Role, bool, error)
	Add(bean permissionModel.Role) (uint, error)
	Update(bean permissionModel.Role) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// 角色
type role struct {
	db mysql.DBRepo
}

// NewRoleDao 创建角色 Dao 对象
func NewRoleDao() *role {
	return &role{
		db: mysql.Instance(),
	}
}

// All 获取所有角色列表
func (d *role) All() ([]permissionModel.Role, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		return stats
	}

	bean := make([]permissionModel.Role, 0)
	if result := stats().Order("updated_at DESC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&permissionModel.Role{}).Count(&total)
	return bean, total, nil
}

// List 查询角色列表
func (d *role) List(req permissionDTO.QueryRoleReq) ([]permissionModel.Role, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		if req.Name != "" {
			stats = stats.Where("name like ?", "%"+req.Name+"%")
		}
		return stats
	}

	bean := make([]permissionModel.Role, 0)
	result := stats().Offset(req.Offset()).Limit(req.PageSize).
		Order("sort DESC").Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&permissionModel.Role{}).Count(&total)
	return bean, total, nil
}

// Info 获取角色信息
func (d *role) InfoByName(name string) (permissionModel.Role, bool, error) {
	bean := permissionModel.Role{}
	result := d.db.GetDbR().Where("name=?", name).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return permissionModel.Role{}, false, nil
	}
	if result.Error != nil {
		return permissionModel.Role{}, false, result.Error
	}
	return bean, true, nil
}

// Add 添加角色
func (d *role) Add(bean permissionModel.Role) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新角色
func (d *role) Update(bean permissionModel.Role) (int64, error) {
	result := d.db.GetDbW().Select("name", "status", "sort", "note").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除角色
func (d *role) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&permissionModel.Role{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除角色
func (d *role) BatchDelete(ids []uint) (int64, error) {
	beans := make([]permissionModel.Role, len(ids))
	for _, id := range ids {
		beans = append(beans, permissionModel.Role{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *role) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&permissionModel.Role{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
