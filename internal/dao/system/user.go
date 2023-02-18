/*用户 DAO
 */
package systemDAO

import (
	"errors"

	DAO "gin-admin/internal/dao"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/repository/mysql"
	"gin-admin/internal/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// User 用户接口
type User interface {
	All() ([]systemModel.User, int64, error)
	List(req systemDTO.QueryUserReq) ([]systemModel.User, int64, error)
	Info(id uint) (systemModel.User, bool, error)
	Add(user systemModel.User, roleIds []uint) error
	Update(user systemModel.User, roles []uint) error
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	UpdatePassword(id uint, password string) (int64, error)
	ResetPassword(id uint, password string) (int64, error)
	UpdatePhone(id uint, phone string) (int64, error)
	UpdateEmail(id uint, email string) (int64, error)
	GetUserByPhone(phone string) (systemModel.User, bool, error)
	GetUserByEmail(email string) (systemModel.User, bool, error)
	ExistUserPassword(userId uint, password string) (bool, error)
	GetUsername(username, password string) (systemModel.User, bool, error)
}

// 用户
type user struct {
	*DAO.Transaction
	db mysql.DBRepo
}

// 创建用户 Dao 对象
func NewUserDao() *user {
	return &user{
		Transaction: DAO.NewTransaction(mysql.Instance().GetDbW()),
		db:          mysql.Instance(),
	}
}

// All 获取所有用户列表
func (d *user) All() ([]systemModel.User, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		return stats
	}

	bean := make([]systemModel.User, 0)
	result := stats().Model(&systemModel.User{}).Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.User{}).Count(&total)
	return bean, total, nil
}

// List 获取用户列表
func (d *user) List(req systemDTO.QueryUserReq) ([]systemModel.User, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		if req.Nickname != "" {
			stats = stats.Where("nickname like ?", "%"+req.Nickname+"%")
		}
		if req.Phone != "" {
			stats = stats.Where("phone like ?", "%"+req.Phone+"%")
		}
		if req.Email != "" {
			stats = stats.Where("email like ?", "%"+req.Email+"%")
		}
		return stats
	}

	bean := make([]systemModel.User, 0)
	result := stats().Model(&systemModel.User{}).Preload("Roles").
		Offset(req.Offset()).Limit(req.PageSize).
		Order("sort DESC").Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.User{}).Count(&total)
	return bean, total, nil
}

// Info 获取用户信息
func (d *user) Info(id uint) (systemModel.User, bool, error) {
	bean := systemModel.User{ID: id}
	result := d.db.GetDbR().Model(&systemModel.User{}).Preload("Roles", "status=1").First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return systemModel.User{}, false, nil
	}
	if result.Error != nil {
		return systemModel.User{}, false, result.Error
	}
	return bean, true, nil
}

// Add 添加用户
func (d *user) Add(user systemModel.User, roleIds []uint) error {
	d.Begin()
	defer func() {
		if err := recover(); err != nil {
			d.Rollback()
			zap.S().Panic("注册用户异常, err: %v", err)
		}
	}()
	// 添加用户
	userId, err := d.addUser(user)
	if err != nil {
		d.Rollback()
		return err
	}
	// 添加用户角色
	if err := d.addUserRole(userId, roleIds); err != nil {
		d.Rollback()
		return err
	}
	d.Commit()
	return nil
}

// 添加用户
func (d *user) addUser(bean systemModel.User) (uint, error) {
	result := d.Tx().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// 添加用户角色关联信息
func (d *user) addUserRole(userId uint, roleIds []uint) error {
	if len(roleIds) == 0 {
		return nil
	}
	roles := make([]systemModel.UserRoleRel, 0)
	for _, roleId := range roleIds {
		roles = append(roles, systemModel.UserRoleRel{
			UserId: userId,
			RoleId: roleId,
		})
	}
	result := d.Tx().Create(&roles)
	return result.Error
}

// Update 更新用户详情信息
func (d *user) Update(user systemModel.User, roles []uint) error {
	d.Begin()
	defer func() {
		if err := recover(); err != nil {
			d.Rollback()
			zap.S().Panic("注册用户异常, err: %v", err)
		}
	}()

	// 更新用户信息
	if err := d.updateUser(user); err != nil {
		d.Rollback()
		return err
	}
	// 更新用户角色信息
	if err := d.updateUserRoles(user.ID, roles); err != nil {
		d.Rollback()
		return err
	}
	d.Commit()
	return nil
}

// 更新用户信息
func (d *user) updateUser(user systemModel.User) error {
	result := d.Tx().
		Select("*").Omit("password", "created_at").Updates(&user)
	return result.Error
}

// 更新用户角色信息
func (d *user) updateUserRoles(userId uint, roleIds []uint) error {
	// 未传入 role_ids, 不做处理
	if roleIds == nil {
		return nil
	}
	// 获取用户关联的角色 roleId 列表
	userRoleIds, err := d.getUserRoleByRoleIds(userId)
	if err != nil {
		return err
	}
	// 新增用户角色关联信息列表
	addUserRoles := make([]systemModel.UserRoleRel, 0)
	for _, roleId := range roleIds {
		if utils.IndexOfArray(userRoleIds, roleId) == -1 {
			addUserRoles = append(addUserRoles, systemModel.UserRoleRel{
				UserId: userId,
				RoleId: roleId,
			})
		}
	}

	// 删除用户角色关联信息列表
	deleteUserRoleIds := make([]uint, 0)
	for _, roleId := range userRoleIds {
		if utils.IndexOfArray(roleIds, roleId) == -1 {
			deleteUserRoleIds = append(deleteUserRoleIds, roleId)
		}
	}

	if len(addUserRoles) != 0 {
		if result := d.Tx().Create(&addUserRoles); result.Error != nil {
			return result.Error
		}
	}
	if len(deleteUserRoleIds) != 0 {
		if result := d.Tx().Where("user_id = ? AND role_id in ?", userId, deleteUserRoleIds).
			Delete(&systemModel.UserRoleRel{}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// 获取用户关联的角色 roleId 列表
func (d *user) getUserRoleByRoleIds(userId uint) ([]uint, error) {
	userRoles := make([]systemModel.UserRoleRel, 0)
	results := d.Tx().Where("user_id = ?", userId).Find(&userRoles)
	if results.Error != nil {
		return nil, results.Error
	}
	roleIds := make([]uint, 0)
	for _, item := range userRoles {
		roleIds = append(roleIds, item.RoleId)
	}
	return roleIds, nil
}

// Delete 删除用户
func (d *user) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&systemModel.User{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除用户
func (d *user) BatchDelete(ids []uint) (int64, error) {
	beans := make([]systemModel.User, len(ids))
	for _, id := range ids {
		beans = append(beans, systemModel.User{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *user) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&systemModel.User{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// UpdatePassword 更新密码
func (d *user) UpdatePassword(id uint, password string) (int64, error) {
	result := d.db.GetDbW().Model(&systemModel.User{}).Where("id = ?", id).
		Update("password", password)
	return result.RowsAffected, result.Error
}

// ResetPassword 重置密码
func (d *user) ResetPassword(id uint, password string) (int64, error) {
	result := d.db.GetDbW().Model(&systemModel.User{}).Where("id = ?", id).Update("password", password)
	return result.RowsAffected, result.Error
}

// UpdatePhone 更新手机号码
func (d *user) UpdatePhone(id uint, phone string) (int64, error) {
	result := d.db.GetDbW().Updates(&systemModel.User{
		ID:    id,
		Phone: phone,
	})
	return result.RowsAffected, result.Error
}

// UpdateEmail 更新邮箱
func (d *user) UpdateEmail(id uint, email string) (int64, error) {
	result := d.db.GetDbW().Updates(&systemModel.User{
		ID:    id,
		Email: email,
	})
	return result.RowsAffected, result.Error
}

// GetUserByPhone 获取用户信息
func (d *user) GetUserByPhone(phone string) (systemModel.User, bool, error) {
	bean := systemModel.User{}
	result := d.db.GetDbW().Where("phone=?", phone).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return systemModel.User{}, false, nil
	}
	if result.Error != nil {
		return systemModel.User{}, false, result.Error
	}
	return bean, true, nil
}

// GetUserByEmail 获取用户信息
func (d *user) GetUserByEmail(email string) (systemModel.User, bool, error) {
	bean := systemModel.User{}
	result := d.db.GetDbR().Where("email=?", email).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return systemModel.User{}, false, nil
	}
	if result.Error != nil {
		return systemModel.User{}, false, result.Error
	}
	return bean, true, nil
}

// ExistUserPassword 判断用户密码是否正确
func (d *user) ExistUserPassword(userId uint, password string) (bool, error) {
	result := d.db.GetDbR().Where("id = ? AND password = ?", userId, password).First(&systemModel.User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// GetUsername 获取用户信息 邮件/手机号
func (d *user) GetUsername(username, password string) (systemModel.User, bool, error) {
	bean := systemModel.User{}
	result := d.db.GetDbR().
		Where("(phone = ? OR email = ?) AND password = ?", username, username, password).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return systemModel.User{}, false, nil
	}
	if result.Error != nil {
		return systemModel.User{}, false, result.Error
	}
	return bean, true, nil
}
