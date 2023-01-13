/*
 * @Author: silent-rain
 * @Date: 2023-01-08 13:43:50
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 22:05:42
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/user_register.go
 * @Descripttion: 用户注册 Dao
 */
package systemDao

import (
	"gin-admin/internal/dao"
	systemModel "gin-admin/internal/model/system"

	"go.uber.org/zap"
)

// UserRegisterImpl 用户注册对象
var UserRegisterImpl = &userRegister{
	Transaction: dao.NewTransaction(),
}

// UserRegister 用户接口
type UserRegister interface {
	Add(user *systemModel.User, roleIds []uint) error
}

// 用户注册结构
type userRegister struct {
	*dao.Transaction
}

// 注册用户
func (d *userRegister) Add(user *systemModel.User, roleIds []uint) error {
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
func (d *userRegister) addUser(bean *systemModel.User) (uint, error) {
	result := d.Tx().Create(bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// 添加用户角色
func (d *userRegister) addUserRole(userId uint, roleIds []uint) error {
	if len(roleIds) == 0 {
		return nil
	}
	roles := make([]systemModel.Role, len(roleIds))
	result := d.Tx().Create(&roles)
	return result.Error
}
