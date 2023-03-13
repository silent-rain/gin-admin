/*用户API接口Token令牌表*/
package permission

import "github.com/silent-rain/gin-admin/internal/dto"

// QueryUserApiTokenReq 查询条件
type QueryUserApiTokenReq struct {
	dto.Pagination        // 分页
	UserId         *uint  `json:"user_id" form:"user_id"`   // 用户ID
	Nickname       string `json:"nickname" form:"nickname"` // 用户昵称
	Status         *uint  `json:"status" form:"status"`     // 状态,0:停用,1:启用
}

// AddUserApiTokenReq 添加 Token 令牌
type AddUserApiTokenReq struct {
	UserId     uint   `json:"user_id" form:"user_id"`       // 用户ID
	Permission string `json:"permission" form:"permission"` // 权限:GET,POST,PUT,DELETE
	Passphrase string `json:"passphrase" form:"passphrase"` // 口令
	Note       string `json:"note" form:"note"`             // 备注
	Status     uint   `json:"status" form:"status"`         // 状态,0:停用,1:启用
}

// UpdateUserApiTokenReq 更新 Token 令牌
type UpdateUserApiTokenReq struct {
	ID         uint   `json:"id" form:"id" binding:"required"` // 令牌ID
	Permission string `json:"permission" form:"permission"`    // 权限:GET,POST,PUT,DELETE
	Passphrase string `json:"passphrase" form:"passphrase"`    // 口令
	Note       string `json:"note" form:"note"`                // 备注
	Status     uint   `json:"status" form:"status"`            // 状态,0:停用,1:启用
}
