/*
 * @Author: silent-rain
 * @Date: 2023-01-08 14:31:45
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 14:46:08
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/user_register.go
 * @Descripttion: 用户注册数据传输
 */
package systemDto

// 用户注册
type UserRegisterReq struct {
	Realname string `form:"realname" binding:"required"` // 真实姓名
	Nickname string `form:"nickname" binding:"required"` // 昵称
	Gender   uint   `form:"gender" binding:"required"`   // 性别: 0:女,1:男
	Age      uint8  `form:"age" binding:"required"`      // 年龄
	Birthday string `form:"birthday" binding:"required"` // 出生日期
	Avatar   string `form:"avatar"`                      // 用户头像URL
	Phone    string `form:"phone"`                       // 手机号码
	Email    string `form:"email"`                       // 邮件
	Intro    string `form:"intro"`                       // 介绍
	Note     string `form:"note"`                        // 备注
	Password string `form:"password" binding:"required"` // 密码, 仅创建（禁止从 db 读）
	RoleIds  []uint `form:"role_ids"`                    // 角色IDs
}
