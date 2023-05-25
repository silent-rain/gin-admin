// Package dto 用户登录信息表
package dto

import DTO "github.com/silent-rain/gin-admin/internal/dto"

// QueryUserLoginReq 查询条件
type QueryUserLoginReq struct {
	DTO.Pagination        // 分页
	Nickname       string `json:"nickname" form:"nickname"`       // 昵称
	RemoteAddr     string `json:"remote_addr" form:"remote_addr"` // 请求IP
}

// UpdateUserLoginStatusReq 更新数据状态
type UpdateUserLoginStatusReq struct {
	ID     uint `json:"id" form:"id" binding:"required"`           // 数据 ID
	UserId uint `json:"user_id" form:"user_id" binding:"required"` // 用户 ID
	Status uint `json:"status" form:"status"`                      // 状态
}

// UserLogin 登录响应
type UserLogin struct {
	Token string `json:"token"` // 令牌
}
