/*
 * @Author: silent-rain
 * @Date: 2023-01-08 14:31:45
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 21:59:12
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/user_register.go
 * @Descripttion: 用户注册数据传输
 */
package systemDto

// RegisterUserReq 用户注册
type RegisterUserReq struct {
	Realname string `json:"realname" form:"realname"`                    // 真实姓名
	Nickname string `json:"nickname" form:"nickname" binding:"required"` // 昵称
	Gender   uint   `json:"gender" form:"gender"`                        // 性别: 0:女,1:男
	Age      uint8  `json:"age" form:"age" binding:"required"`           // 年龄
	Birthday string `json:"birthday" form:"birthday" binding:"required"` // 出生日期
	Avatar   string `json:"avatar" form:"avatar"`                        // 用户头像URL
	Phone    string `json:"phone" form:"phone" binding:"required"`       // 手机号码
	Email    string `json:"email" form:"email"`                          // 邮件
	Intro    string `json:"intro" form:"intro"`                          // 介绍
	Note     string `json:"note" form:"note"`                            // 备注
	Password string `json:"password" form:"password" binding:"required"` // 密码, 仅创建（禁止从 db 读）
	RoleIds  []uint `json:"role_ids" form:"role_ids"`                    // 角色IDs
}
