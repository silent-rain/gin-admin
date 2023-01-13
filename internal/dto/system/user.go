/*
 * @Author: silent-rain
 * @Date: 2023-01-08 17:14:54
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 23:52:03
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/user.go
 * @Descripttion:
 */
package systemDto

import "gin-admin/internal/dto"

// UserLoginReq 登录请求
type UserLoginReq struct {
	Username string `form:"username" binding:"required"` // 用户 手机号、邮箱
	Captcha  string `form:"captcha" binding:"required"`  // 验证码
	Password string `form:"password" binding:"required"` // 密码
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

// UserDeleteReq 删除用户
type UserDeleteReq struct {
	ID uint `json:"id" form:"id"` // 用户ID
}

// 更新用户状态
type UserStatusReq struct {
	ID     uint `form:"id" binding:"required"`
	Status uint `form:"status"`
}
