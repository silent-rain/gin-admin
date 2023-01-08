/*
 * @Author: silent-rain
 * @Date: 2023-01-08 13:19:16
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 17:19:55
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/user.go
 * @Descripttion: 用户 Dao
 */
package systemDao

import (
	"errors"

	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"

	"gorm.io/gorm"
)

// 用户对象
var UserImpl = new(user)

// 用户接口
type User interface {
	ExistUsername(phone, email string) (bool, error)
	GetUsername(username, password string) (*systemModel.User, bool, error)
}

// 用户结构
type user struct{}

// GetList 获取用户列表
func (d *user) GetList() {

}

// ExistUserName 判断用户是否存在 邮件/手机号
func (d *user) ExistUsername(phone, email string) (bool, error) {
	result := database.Instance().Debug().Where("phone = ? OR email = ?", phone, email).First(&systemModel.User{})
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
	result := database.Instance().Debug().Where("(phone = ? OR email = ?) AND password = ?", username, username, password).First(bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if result.Error != nil {
		return nil, false, result.Error
	}
	return bean, true, nil
}

// GetList 添加用户
func (d *user) Add() {

}
