/*ICON 图标
 */
package system

// Icon ICON 图标
type Icon struct {
	ID        uint   `json:"id" gorm:"column:id;primaryKey"`      // 配置ID
	Name      string `json:"name" gorm:"column:name"`             // 配置名称
	Value     string `json:"value" gorm:"column:value"`           // 配置值
	Category  uint   `json:"category" gorm:"column:category"`     // 图标类型,1:element,2:custom
	Note      string `json:"note" gorm:"column:note"`             // 备注
	CreatedAt string `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

// TableName 表名重写
func (Icon) TableName() string {
	return "sys_icon"
}
