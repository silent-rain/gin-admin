/*角色
 */
package permission

import "gin-admin/internal/dto"

// QueryRoleReq 查询条件
type QueryRoleReq struct {
	dto.Pagination        // 分页
	Name           string `json:"name" form:"name"` // 角色名称
}

// AddRoleReq 添加角色
type AddRoleReq struct {
	Name   string `json:"name" form:"name" binding:"required"` // 角色名称
	Status uint   `json:"status" form:"status"`                // 角色状态,0:停用,1:启用
	Sort   uint   `json:"sort" form:"sort"`                    // 排序
	Note   string `json:"note" form:"note"`                    // 备注
}

// UpdateRoleReq 更新角色
type UpdateRoleReq struct {
	ID     uint   `json:"id" form:"id" binding:"required"`     // 角色ID
	Name   string `json:"name" form:"name" binding:"required"` // 角色名称
	Status uint   `json:"status" form:"status"`                // 角色状态,0:停用,1:启用
	Sort   uint   `json:"sort"  form:"sort"`                   // 排序
	Note   string `json:"note"  form:"note"`                   // 备注
}
