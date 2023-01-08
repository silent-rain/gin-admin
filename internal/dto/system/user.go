/*
 * @Author: silent-rain
 * @Date: 2023-01-08 17:14:54
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:40:33
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/user.go
 * @Descripttion:
 */
package systemDto

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
