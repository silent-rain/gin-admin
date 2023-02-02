/*
 * @Author: silent-rain
 * @Date: 2023-01-08 17:14:54
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 17:51:59
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/user.go
 * @Descripttion:
 */
package systemDto

import "gin-admin/internal/dto"

// UserLoginReq 登录请求
type UserLoginReq struct {
	Username string `json:"username" form:"username" binding:"required"` // 用户 手机号、邮箱
	Password string `json:"password" form:"password" binding:"required"` // 密码
}

// UserLoginRsp 登录响应
type UserLoginRsp struct {
	Token string `json:"token"` // 令牌
}

// UserInfoRsp 用户信息响应
type UserInfoRsp struct {
	ID       uint   `json:"id"`       // 用户ID
	Nickname string `json:"nickname"` // 昵称
	Phone    string `json:"phone"`    // 手机号码
	Email    string `json:"email"`    // 邮件
	Avatar   string `json:"avatar"`   // 用户头像URL
	RoleIds  []uint `json:"role_ids"` // 角色 ID 列表
}

// UserQueryReq 查询条件
type UserQueryReq struct {
	dto.Pagination        // 分页
	Nickname       string `json:"nickname" form:"nickname"` // 用户昵称
	Phone          string `json:"phone" form:"phone"`       // 手机号码
	Email          string `json:"email" form:"email"`       // 邮件
}

// UserUpdateDetailsReq 更新用户详情信息
type UserUpdateDetailsReq struct {
	ID       uint   `json:"id" form:"id" binding:"required"` // 用户ID
	Realname string `json:"realname" form:"realname"`        // 真实姓名
	Nickname string `json:"nickname" form:"nickname"`        // 昵称
	Gender   uint   `json:"gender" form:"gender"`            // 性别: 0:女,1:男
	Age      uint8  `json:"age" form:"age"`                  // 年龄
	Birthday string `json:"birthday" form:"birthday"`        // 出生日期
	Avatar   string `json:"avatar" form:"avatar"`            // 用户头像URL
	Phone    string `json:"phone" form:"phone"`              // 手机号码
	Email    string `json:"email" form:"email"`              // 邮件
	Intro    string `json:"intro" form:"intro"`              // 介绍
	Note     string `json:"note" form:"note"`                // 备注
	RoleIds  []uint `json:"role_ids" form:"role_ids"`        // 角色IDs
}

// UserDeleteReq 删除用户
type UserDeleteReq struct {
	ID uint `json:"id" form:"id" binding:"required"` // 用户ID
}

// UserBatchDeleteReq 批量删除用户
type UserBatchDeleteReq struct {
	Ids []uint `json:"ids" form:"ids" binding:"required"` // 用户ID列表
}

// UserStatusReq 更新用户状态
type UserStatusReq struct {
	ID     uint `json:"id" form:"id" binding:"required"`
	Status uint `json:"status" form:"status"`
}

// UserResetPasswordReq 重置用户密码
type UserResetPasswordReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

// UserUpdatePasswordReq 用户密码更新
type UserUpdatePasswordReq struct {
	ID          uint   `json:"id" form:"id" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"` // 旧密码
	NewPassword string `json:"new_password" form:"new_password" binding:"required"` // 新密码
}

// UserUpdatePhoneReq 用户更新手机号码
type UserUpdatePhoneReq struct {
	ID    uint   `json:"id" form:"id" binding:"required"`
	Phone string `json:"phone" form:"phone"` // 手机号码
}

// UserUpdateEmailReq 用户更新邮箱
type UserUpdateEmailReq struct {
	ID    uint   `json:"id" form:"id" binding:"required"`
	Email string `json:"email" form:"email"` // 邮件
}
