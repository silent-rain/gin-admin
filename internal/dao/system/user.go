/*
 * @Author: silent-rain
 * @Date: 2023-01-08 13:19:16
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 23:47:57
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/user.go
 * @Descripttion: 用户 Dao
 */
package systemDao

import (
	"errors"

	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"

	"gorm.io/gorm"
)

// UserImpl 用户对象
var UserImpl = new(user)

// User 用户接口
type User interface {
	List(req systemDto.UserQueryReq) ([]systemModel.User, int64, error)
	Delete(id uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	ExistUsername(phone, email string) (bool, error)
	GetUsername(username, password string) (*systemModel.User, bool, error)
}

// 用户结构
type user struct{}

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

// Delete 删除用户
func (d *user) Delete(id uint) (int64, error) {
	result := database.Instance().Delete(&systemModel.User{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *user) Status(id uint, status uint) (int64, error) {
	result := database.Instance().Select("status").Updates(&systemModel.User{
		ID:     id,
		Status: status,
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
