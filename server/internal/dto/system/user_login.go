/*用户登录信息表*/
package system

import "gin-admin/internal/dto"

// QueryUserLoginReq 查询条件
type QueryUserLoginReq struct {
	dto.Pagination        // 分页
	Nickname       string `json:"nickname" form:"nickname"`       // 昵称
	RemoteAddr     string `json:"remote_addr" form:"remote_addr"` // 请求IP
}
