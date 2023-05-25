// Package model 应用配置表
package model

// Config 应用配置
type Config struct {
	ID        uint     `json:"id" gorm:"column:id;primaryKey"`      // 配置ID
	ParentId  *uint    `json:"parent_id" gorm:"column:parent_id"`   // 父配置ID
	Name      string   `json:"name" gorm:"column:name"`             // 配置名称
	Key       string   `json:"key" gorm:"column:key"`               // 配置KEY
	Value     string   `json:"value" gorm:"column:value"`           // 配置值
	Sort      uint     `json:"sort" gorm:"column:sort"`             // 排序
	Note      string   `json:"note" gorm:"column:note"`             // 备注
	Status    uint     `json:"status" gorm:"column:status"`         // 状态,0:停用,1:启用
	CreatedAt string   `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt string   `json:"updated_at" gorm:"column:updated_at"` // 更新时间
	Children  []Config `json:"children" gorm:"foreignKey:ParentId"` // 子配置列表
}

// TableName 表名重写
func (Config) TableName() string {
	return "sys_config"
}
