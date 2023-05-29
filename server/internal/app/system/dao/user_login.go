// Package dao 用户登录信息表
package dao

import (
	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// UserLogin 用户登录信息接口
type UserLogin interface {
	List(req dto.QueryUserLoginReq) ([]model.UserLogin, int64, error)
	Add(bean model.UserLogin) (uint, error)
	Status(id uint, status uint) (int64, error)
}

// 用户登录信息
type userLogin struct {
	mysql.DBRepo
}

// NewUserLoginDao 创建用户登录信息对象
func NewUserLoginDao() *userLogin {
	return &userLogin{
		DBRepo: global.Instance().Mysql(),
	}
}

// List 查询用户登录信息列表
func (d *userLogin) List(req dto.QueryUserLoginReq) ([]model.UserLogin, int64, error) {
	tx := d.GetDbR()
	if req.Nickname != "" {
		tx = tx.Where("nickname like ?", req.Nickname+"%")
	}
	if req.RemoteAddr != "" {
		tx = tx.Where("remote_addr like ?", req.RemoteAddr+"%")
	}
	tx = tx.Session(&gorm.Session{})

	bean := make([]model.UserLogin, 0)
	var total int64 = 0
	tx.Model(&model.UserLogin{}).Count(&total)

	result := tx.Offset(req.Offset()).Limit(req.PageSize).
		Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// Add 添加用户登录信息
func (d *userLogin) Add(bean model.UserLogin) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Status 更新用户登录信息状态
func (d *userLogin) Status(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.UserLogin{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
