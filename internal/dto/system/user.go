/*
 * @Author: silent-rain
 * @Date: 2023-01-08 17:14:54
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 17:32:50
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/user.go
 * @Descripttion:
 */
package systemDto

// 登录请求
type UserLoginReq struct {
	Username string `form:"username" binding:"required"` // 用户 手机号、邮箱
	Captcha  string `form:"captcha" binding:"required"`  // 验证码
	Password string `form:"password" binding:"required"` // 密码
}

// 登录响应
type UserLoginRsp struct {
	Token string `json:"token"` // 令牌
}
