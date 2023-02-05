/* 角色菜单
 */
package systemModel

import "gin-admin/internal/pkg/utils"

// RoleMenuRel 角色菜单关联表
type RoleMenuRel struct {
	ID        uint             `json:"id" gorm:"column:id;primaryKey"`                           // 自增ID
	RoleId    uint             `json:"role_id" gorm:"column:role_id"`                            // 角色ID
	MenuId    uint             `json:"menu_id" gorm:"column:menu_id"`                            // 菜单ID
	CreatedAt *utils.LocalTime `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
	UpdatedAt *utils.LocalTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"` // 更新时间
}

// TableName 表名重写
func (RoleMenuRel) TableName() string {
	return "sys_role_menu_rel"
}
