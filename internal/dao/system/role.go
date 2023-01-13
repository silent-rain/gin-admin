/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:24:36
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 21:54:02
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/role.go
 * @Descripttion: 角色
 */
package systemDao

import (
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"
)

// RoleImpl 角色对象
var RoleImpl = new(role)

// Role 角色接口
type Role interface {
	List(req systemDto.RoleQueryReq) ([]systemModel.Role, int64, error)
	Add(bean systemModel.Role) (uint, error)
	Update(bean systemModel.Role) (int64, error)
	Delete(id uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// 角色结构
type role struct{}

// List 查询角色列表
func (d *role) List(req systemDto.RoleQueryReq) ([]systemModel.Role, int64, error) {
	var db = database.Instance()
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}
	bean := make([]systemModel.Role, 0)
	if result := db.Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	db.Count(&total)
	return bean, total, nil
}

// Add 添加角色
func (d *role) Add(bean systemModel.Role) (uint, error) {
	result := database.Instance().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新角色
func (d *role) Update(bean systemModel.Role) (int64, error) {
	result := database.Instance().Select("name", "status", "sort", "note").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除角色
func (d *role) Delete(id uint) (int64, error) {
	result := database.Instance().Delete(&systemModel.Role{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *role) Status(id uint, status uint) (int64, error) {
	result := database.Instance().Select("status").Updates(&systemModel.Role{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
