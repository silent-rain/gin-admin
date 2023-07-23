// Package dao 用户管理
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"
	"github.com/silent-rain/gin-admin/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	mysql.DBRepo
}

// NewUserDao 创建用户 Dao 对象
func NewUserDao() *User {
	return &User{
		DBRepo: global.Instance().Mysql(),
	}
}

// All 获取所有用户列表
func (d *User) All() ([]model.User, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
		return stats
	}

	bean := make([]model.User, 0)
	result := stats().Model(&model.User{}).Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.User{}).Count(&total)
	return bean, total, nil
}

// List 获取用户列表
func (d *User) List(req dto.QueryUserReq) ([]model.User, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.GetDbR()
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

	bean := make([]model.User, 0)
	result := stats().Model(&model.User{}).Preload("Roles").
		Offset(req.Offset()).Limit(req.PageSize).
		Order("sort DESC").Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&model.User{}).Count(&total)
	return bean, total, nil
}

// Info 获取用户信息
func (d *User) Info(id uint) (model.User, bool, error) {
	bean := model.User{ID: id}
	result := d.GetDbR().Model(&model.User{}).Preload("Roles", "status=1").First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.User{}, false, nil
	}
	if result.Error != nil {
		return model.User{}, false, result.Error
	}
	return bean, true, nil
}

// Add 添加用户
func (d *User) Add(User model.User, roleIds []uint) error {
	tx := d.GetDbW().Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			zap.S().Panic("注册用户异常, err: %v", err)
		}
	}()
	// 添加用户
	userId, err := d.addUser(tx, User)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 添加用户角色
	if err := d.addUserRole(tx, userId, roleIds); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 添加用户
func (d *User) addUser(tx *gorm.DB, bean model.User) (uint, error) {
	result := tx.Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// 添加用户角色关联信息
func (d *User) addUserRole(tx *gorm.DB, userId uint, roleIds []uint) error {
	if len(roleIds) == 0 {
		return nil
	}
	roles := make([]model.UserRoleRel, 0)
	for _, roleId := range roleIds {
		roles = append(roles, model.UserRoleRel{
			UserId: userId,
			RoleId: roleId,
		})
	}
	result := tx.Create(&roles)
	return result.Error
}

// Update 更新用户详情信息
func (d *User) Update(User model.User, roles []uint) error {
	tx := d.GetDbW().Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			zap.S().Panic("注册用户异常, err: %v", err)
		}
	}()

	// 更新用户信息
	if err := d.updateUser(tx, User); err != nil {
		tx.Rollback()
		return err
	}
	// 更新用户角色信息
	if err := d.updateUserRoles(tx, User.ID, roles); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 更新用户信息
func (d *User) updateUser(tx *gorm.DB, User model.User) error {
	result := tx.
		Select("*").Omit("password", "created_at").Updates(&User)
	return result.Error
}

// 更新用户角色信息
func (d *User) updateUserRoles(tx *gorm.DB, userId uint, roleIds []uint) error {
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
	addUserRoles := make([]model.UserRoleRel, 0)
	for _, roleId := range roleIds {
		if utils.IndexOfArray(userRoleIds, roleId) == -1 {
			addUserRoles = append(addUserRoles, model.UserRoleRel{
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
		if result := tx.Create(&addUserRoles); result.Error != nil {
			return result.Error
		}
	}
	if len(deleteUserRoleIds) != 0 {
		if result := tx.Where("user_id = ? AND role_id in ?", userId, deleteUserRoleIds).
			Delete(&model.UserRoleRel{}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// 获取用户关联的角色 roleId 列表
func (d *User) getUserRoleByRoleIds(userId uint) ([]uint, error) {
	userRoles := make([]model.UserRoleRel, 0)
	results := d.GetDbR().Where("user_id = ?", userId).Find(&userRoles)
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
func (d *User) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.User{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除用户
func (d *User) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.User, len(ids))
	for _, id := range ids {
		beans = append(beans, model.User{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// UpdateStatus 更新状态
func (d *User) UpdateStatus(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.User{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// UpdatePassword 更新密码
func (d *User) UpdatePassword(id uint, password string) (int64, error) {
	result := d.GetDbW().Model(&model.User{}).Where("id = ?", id).
		Update("password", password)
	return result.RowsAffected, result.Error
}

// ResetPassword 重置密码
func (d *User) ResetPassword(id uint, password string) (int64, error) {
	result := d.GetDbW().Model(&model.User{}).Where("id = ?", id).Update("password", password)
	return result.RowsAffected, result.Error
}

// UpdatePhone 更新手机号码
func (d *User) UpdatePhone(id uint, phone string) (int64, error) {
	result := d.GetDbW().Updates(&model.User{
		ID:    id,
		Phone: phone,
	})
	return result.RowsAffected, result.Error
}

// UpdateEmail 更新邮箱
func (d *User) UpdateEmail(id uint, email string) (int64, error) {
	result := d.GetDbW().Updates(&model.User{
		ID:    id,
		Email: email,
	})
	return result.RowsAffected, result.Error
}

// GetUserByPhone 获取用户信息
func (d *User) GetUserByPhone(phone string) (model.User, bool, error) {
	bean := model.User{}
	result := d.GetDbW().Where("phone=?", phone).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.User{}, false, nil
	}
	if result.Error != nil {
		return model.User{}, false, result.Error
	}
	return bean, true, nil
}

// GetUserByEmail 获取用户信息
func (d *User) GetUserByEmail(email string) (model.User, bool, error) {
	bean := model.User{}
	result := d.GetDbR().Where("email=?", email).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.User{}, false, nil
	}
	if result.Error != nil {
		return model.User{}, false, result.Error
	}
	return bean, true, nil
}

// ExistUserPassword 判断用户密码是否正确
func (d *User) ExistUserPassword(userId uint, password string) (bool, error) {
	result := d.GetDbR().Where("id = ? AND password = ?", userId, password).First(&model.User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// InfoByApiToken 通过 api token 获取用户信息
func (d *User) InfoByApiToken(token string) (model.User, bool, error) {
	bean := model.User{}
	result := d.GetDbR().Model(&model.User{}).
		Select("perm_user.*").
		Joins("LEFT JOIN perm_user_api_token ON perm_user_api_token.user_id = perm_user.id").
		Where("perm_user_api_token.token = ?", token).
		First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}
