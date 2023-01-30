/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:20:26
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 22:49:48
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/model/system/role.go
 * @Descripttion: 角色
 */
package systemModel

import (
	"gin-admin/internal/pkg/utils"
)

// Role 角色表
type Role struct {
	ID     uint   `json:"id" gorm:"column:id;primaryKey"` // 角色ID
	Name   string `json:"name" gorm:"column:name"`        // 角色名称
	Sort   uint   `json:"sort" gorm:"column:sort"`        // 排序
	Note   string `json:"note" gorm:"column:note"`        // 备注
	Status uint   `json:"status" gorm:"column:status"`    // 角色状态,0:停用,1:启用
	// 创建时间
	CreatedAt *utils.LocalTime `json:"created_at" gorm:"column:created_at;autoCreateTime:milli;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
	// 更新时间
	UpdatedAt *utils.LocalTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp"`
}

// TableName 表名重写
func (Role) TableName() string {
	return "sys_role"
}
