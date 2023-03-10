/*角色与Http协议接口关联表*/
package apiauth

// QueryApiRoleHttpRelReq 角色Http协议接口关系查询条件
type QueryApiRoleHttpRelReq struct {
	RoleId uint `json:"role_id" form:"role_id"` // 角色ID, role_id/api_id 不能同时为空
	ApiId  uint `json:"api_id" form:"api_id"`   // Http协议接口ID
}

// UpdateApiRoleHttpRelReq 更新角色Http协议接口关系
type UpdateApiRoleHttpRelReq struct {
	RoleId uint   `json:"role_id" form:"role_id" binding:"required"` // 角色ID
	ApiIds []uint `json:"api_ids" form:"api_ids"`                    // 接口IDs
}
