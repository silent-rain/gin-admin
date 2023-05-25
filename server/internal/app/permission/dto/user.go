/*用户
 */
package dto

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	DTO "github.com/silent-rain/gin-admin/internal/dto"
)

// AddUserReq 添加用户
type AddUserReq struct {
	Realname  string `json:"realname" form:"realname"`                    // 真实姓名
	Nickname  string `json:"nickname" form:"nickname" binding:"required"` // 昵称
	Gender    uint   `json:"gender" form:"gender"`                        // 性别: 0:女,1:男
	Age       uint8  `json:"age" form:"age" binding:"required"`           // 年龄
	Birthday  string `json:"birthday" form:"birthday" binding:"required"` // 出生日期
	Avatar    string `json:"avatar" form:"avatar"`                        // 用户头像URL
	Phone     string `json:"phone" form:"phone" binding:"required"`       // 手机号码
	Email     string `json:"email" form:"email"`                          // 邮件
	Intro     string `json:"intro" form:"intro"`                          // 介绍
	Note      string `json:"note" form:"note"`                            // 备注
	Password  string `json:"password" form:"password" binding:"required"` // 密码, 仅创建（禁止从 db 读）
	RoleIds   []uint `json:"role_ids" form:"role_ids"`                    // 角色IDs
	CaptchaId string `json:"captcha_id" form:"captcha_id"`                // 验证码ID
	Captcha   string `json:"captcha" form:"captcha"`                      // 验证码，只有注册的时候才需要
}

// UserInfoRsp 用户信息响应
type UserInfoRsp struct {
	User        model.User         `json:"user"`        // 用户信息
	Roles       []model.Role       `json:"roles"`       // 角色列表
	Menus       []model.Menu       `json:"menus"`       // 菜单路由列表
	Permissions []ButtonPermission `json:"permissions"` // 按钮权限列表
}

// QueryUserReq 查询条件
type QueryUserReq struct {
	DTO.Pagination        // 分页
	Nickname       string `json:"nickname" form:"nickname"` // 用户昵称
	Phone          string `json:"phone" form:"phone"`       // 手机号码
	Email          string `json:"email" form:"email"`       // 邮件
}

// UpdateUserReq 更新用户详情信息
type UpdateUserReq struct {
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
	Sort     uint   `json:"sort" form:"sort"`                // 排序
	Status   uint   `json:"status" form:"status"`            // 状态
	RoleIds  []uint `json:"role_ids" form:"role_ids"`        // 角色IDs
}

// ResetUserPasswordReq 重置用户密码
type ResetUserPasswordReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

// UpdateUserPasswordReq 用户密码更新
type UpdateUserPasswordReq struct {
	ID          uint   `json:"id" form:"id" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"` // 旧密码
	NewPassword string `json:"new_password" form:"new_password" binding:"required"` // 新密码
}

// UpdateUserPhoneReq 用户更新手机号码
type UpdateUserPhoneReq struct {
	ID       uint   `json:"id" form:"id" binding:"required"`             // 用户ID
	Phone    string `json:"phone" form:"phone" binding:"required"`       // 手机号码
	Password string `json:"password" form:"password" binding:"required"` // 密码
}

// UpdateUserEmailReq 用户更新邮箱
type UpdateUserEmailReq struct {
	ID       uint   `json:"id" form:"id" binding:"required"`             // 用户ID
	Email    string `json:"email" form:"email" binding:"required"`       // 邮件
	Password string `json:"password" form:"password" binding:"required"` // 密码
}
