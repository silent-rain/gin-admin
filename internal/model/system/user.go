/*
 * @Author: silent-rain
 * @Date: 2023-01-08 12:41:32
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 16:50:31
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
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`                           // 用户ID
	Realname  string    `json:"realname" gorm:"column:realname"`                          // 真实姓名
	Nickname  string    `json:"nickname" gorm:"column:nickname"`                          // 昵称
	Gender    uint      `json:"gender" gorm:"column:gender"`                              // 性别: 0:女,1:男
	Age       uint8     `json:"age" gorm:"column:age"`                                    // 年龄
	Birthday  string    `json:"birthday" gorm:"column:birthday"`                          // 出生日期
	Avatar    string    `json:"avatar" gorm:"column:avatar"`                              // 用户头像URL
	Phone     string    `json:"phone" gorm:"column:phone"`                                // 手机号码
	Email     string    `json:"email" gorm:"column:email"`                                // 邮件
	Intro     string    `json:"intro" gorm:"column:intro"`                                // 介绍
	Note      string    `json:"note" gorm:"column:note"`                                  // 备注
	Password  string    `json:"password" gorm:"column:password;->:false;<-:create"`       // 密码, 仅创建（禁止从 db 读）
	Sort      uint      `json:"sort" gorm:"column:sort"`                                  // 排序
	Status    uint      `json:"status" gorm:"column:status"`                              // 是否启用,0:禁用,1:启用
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"` // 更新时间
	Roles     []Role    `gorm:"many2many:user_role_rel;"`                                 // Many To Many, 关联 role 表
}

// TableName 表名重写
func (User) TableName() string {
	return "user"
}

// UserRoleRel 用户角色表
type UserRoleRel struct {
	ID        uint      `gorm:"column:id;primaryKey"`                   // 自增ID
	UserId    uint      `gorm:"column:user_id"`                         // 用户ID
	RoleId    uint      `gorm:"column:role_id"`                         // 角色ID
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"` // 更新时间
}

// TableName 表名重写
func (UserRoleRel) TableName() string {
	return "user_role_rel"
}
