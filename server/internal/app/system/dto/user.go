// Package dto 用户
package dto

// UserLoginReq 登录请求
type UserLoginReq struct {
	Username  string `json:"username" form:"username" binding:"required"` // 用户 手机号、邮箱
	Password  string `json:"password" form:"password" binding:"required"` // 密码
	CaptchaId string `json:"captcha_id" form:"captcha_id"`                // 验证码ID
	Captcha   string `json:"captcha" form:"captcha"`                      // 验证码
}
