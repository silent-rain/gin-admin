/*
 * @Author: silent-rain
 * @Date: 2023-01-08 12:41:32
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 22:11:22
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/model/system/user.go
 * @Descripttion: 用户模型
 */
package systemModel

import (
	"time"
)

// User 用户表
type User struct {
	ID        uint      `gorm:"column:id;primaryKey"`                   // 用户ID
	Realname  string    `gorm:"column:realname"`                        // 真实姓名
	Nickname  string    `gorm:"column:nickname"`                        // 昵称
	Gender    uint      `gorm:"column:gender"`                          // 性别: 0:女,1:男
	Age       uint8     `gorm:"column:age"`                             // 年龄
	Birthday  string    `gorm:"column:birthday"`                        // 出生日期
	Avatar    string    `gorm:"column:avatar"`                          // 用户头像URL
	Phone     string    `gorm:"column:phone"`                           // 手机号码
	Email     string    `gorm:"column:email"`                           // 邮件
	Intro     string    `gorm:"column:intro"`                           // 介绍
	Note      string    `gorm:"column:note"`                            // 备注
	Password  string    `gorm:"column:password;->:false;<-:create"`     // 密码, 仅创建（禁止从 db 读）
	Sort      uint      `gorm:"column:sort"`                            // 排序
	Status    uint      `gorm:"column:status"`                          // 是否启用,0:禁用,1:启用
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"` // 更新时间
}

// TableName 表名重写
func (User) TableName() string {
	return "user"
}

// UserRoleRel 用户角色表
type UserRoleRel struct {
	ID        uint      `gorm:"column:id;primaryKey"`                   // 自增ID
	UserId    string    `gorm:"column:user_id"`                         // 用户ID
	RoleId    string    `gorm:"column:role_id"`                         // 角色ID
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"` // 更新时间
}

// TableName 表名重写
func (UserRoleRel) TableName() string {
	return "user_role_rel"
}
