/*角色与Http协议接口关联表*/
package apiauth

// ApiRoleHttpRel 角色与Http协议接口关联表
type ApiRoleHttpRel struct {
	ID        uint   `json:"id" gorm:"column:id;primaryKey"`      // 自增ID
	RoleId    uint   `json:"role_id" gorm:"column:role_id"`       // 角色ID
	ApiId     uint   `json:"api_id" gorm:"column:api_id"`         // Http协议接口ID
	CreatedAt string `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

// TableName 表名重写
func (ApiRoleHttpRel) TableName() string {
	return "api_role_http_rel"
}
