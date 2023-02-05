/* 角色菜单
 */
package systemDto

// QueryRoleMenuRelReq 角色菜单关系查询条件
type QueryRoleMenuRelReq struct {
	RoleId uint `json:"role_id" form:"role_id"` // 角色ID, role_id/menu_id 不能同时为空
	MenuId uint `json:"menu_id" form:"menu_id"` // 菜单ID
}

// UpdateRoleMenuRelReq 更新角色对应的菜单关系
type UpdateRoleMenuRelReq struct {
	RoleId  uint   `json:"role_id" form:"role_id" binding:"required"` // 角色ID
	MenuIds []uint `json:"menu_ids" form:"menu_ids"`                  // 菜单IDs
}
