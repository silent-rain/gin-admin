// Package dao 登录
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/global"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Login 登录
type Login struct {
	mysql.DBRepo
}

// NewLoginDao 创建登录对象
func NewLoginDao() *Login {
	return &Login{
		DBRepo: global.Instance().Mysql(),
	}
}

// Login 查询登录用户信息 邮件/手机号
func (d *Login) Login(username, password string) (model.User, bool, error) {
	bean := model.User{}
	result := d.GetDbR().
		Where("(phone = ? OR email = ?) AND password = ?", username, username, password).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.User{}, false, nil
	}
	if result.Error != nil {
		return model.User{}, false, result.Error
	}
	return bean, true, nil
}
