/*
 * @Author: silent-rain
 * @Date: 2023-01-08 13:19:16
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 17:12:44
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/user.go
 * @Descripttion: 用户 Dao
 */
package systemDao

import (
	"errors"

	"gin-admin/internal/dao"
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"
	"gin-admin/internal/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UserImpl 用户对象
var UserImpl = new(user)

// User 用户接口
type User interface {
	All() ([]systemModel.User, int64, error)
	List(req systemDto.UserQueryReq) ([]systemModel.User, int64, error)
	Info(id uint) (*systemModel.User, bool, error)
	UpdateDetails(user systemModel.User, roles []uint)
	Delete(id uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	UpdatePassword(id uint, password string) (int64, error)
	ResetPassword(id uint, password string) (int64, error)
	UpdatePhone(id uint, phone string) (int64, error)
	UpdateEmail(id uint, email string) (int64, error)
	ExistUsername(phone, email string) (bool, error)
	ExistUserPassword(userId uint, password string) (bool, error)
	GetUsername(username, password string) (*systemModel.User, bool, error)
}

// 用户
type user struct {
	dao.Transaction
}

// All 获取所有用户列表
func (d *user) All() ([]systemModel.User, int64, error) {
	var stats = func() *gorm.DB {
		stats := database.Instance()
		return stats
	}

	bean := make([]systemModel.User, 0)
	result := stats().Model(&systemModel.User{}).Order("sort DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.User{}).Count(&total)
	return bean, total, nil
}

// List 获取用户列表
func (d *user) List(req systemDto.UserQueryReq) ([]systemModel.User, int64, error) {
	var stats = func() *gorm.DB {
		stats := database.Instance()
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
		Offset(req.Offset()).Limit(req.PageSize).Order("sort DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.User{}).Count(&total)
	return bean, total, nil
}

// Info 获取用户信息
func (d *user) Info(id uint) (*systemModel.User, bool, error) {
	bean := &systemModel.User{ID: id}
	result := database.Instance().First(bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	return bean, true, nil
}

// Update 更新用户详情信息
func (d *user) UpdateDetails(user systemModel.User, roles []uint) error {
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
	// 获取用户关联的角色列表
	userRoleIds, err := d.getUserRoleIds(userId)
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

// 获取用户关联角色列表
func (d *user) getUserRoleIds(userId uint) ([]uint, error) {
	userRoles := make([]systemModel.UserRoleRel, 0)
	results := d.Tx().Where("user_id = ?", userId).Find(&userRoles)
	userRoleIds := make([]uint, 0)
	for _, userRole := range userRoles {
		userRoleIds = append(userRoleIds, userRole.RoleId)
	}
	return userRoleIds, results.Error
}

// Delete 删除用户
func (d *user) Delete(id uint) (int64, error) {
	result := database.Instance().Delete(&systemModel.User{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *user) Status(id uint, status uint) (int64, error) {
	result := database.Instance().Updates(&systemModel.User{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// UpdatePassword 更新密码
func (d *user) UpdatePassword(id uint, password string) (int64, error) {
	result := database.Instance().Table("user").Where("id = ?", id).
		Update("password", password)
	return result.RowsAffected, result.Error
}

// ResetPassword 重置密码
func (d *user) ResetPassword(id uint, password string) (int64, error) {
	result := database.Instance().Table("user").Where("id = ?", id).Update("password", password)
	return result.RowsAffected, result.Error
}

// UpdatePhone 更新手机号码
func (d *user) UpdatePhone(id uint, phone string) (int64, error) {
	result := database.Instance().Updates(&systemModel.User{
		ID:    id,
		Phone: phone,
	})
	return result.RowsAffected, result.Error
}

// UpdateEmail 更新邮箱
func (d *user) UpdateEmail(id uint, email string) (int64, error) {
	result := database.Instance().Updates(&systemModel.User{
		ID:    id,
		Email: email,
	})
	return result.RowsAffected, result.Error
}

// ExistUserName 判断用户是否存在 邮件/手机号
func (d *user) ExistUsername(phone, email string) (bool, error) {
	result := database.Instance().Where("phone = ? OR email = ?", phone, email).First(&systemModel.User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// ExistUserPassword 判断用户密码是否正确
func (d *user) ExistUserPassword(userId uint, password string) (bool, error) {
	result := database.Instance().Where("id = ? AND password = ?", userId, password).First(&systemModel.User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// GetUsername 获取用户信息 邮件/手机号
func (d *user) GetUsername(username, password string) (*systemModel.User, bool, error) {
	bean := &systemModel.User{}
	result := database.Instance().Where("(phone = ? OR email = ?) AND password = ?", username, username, password).First(bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if result.Error != nil {
		return nil, false, result.Error
	}
	return bean, true, nil
}
