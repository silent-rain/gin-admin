// Package model 字典维度管理
package model

// Dict 字典维度表
type Dict struct {
	ID        uint   `json:"id" gorm:"column:id;primaryKey"`      // 字典ID
	Name      string `json:"name" gorm:"column:name"`             // 字典名称
	Code      string `json:"code" gorm:"column:code"`             // 字典编码
	Note      string `json:"note" gorm:"column:note"`             // 备注
	Status    uint   `json:"status" gorm:"column:status"`         // 状态,0:停用,1:启用
	CreatedAt string `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

// TableName 表名重写
func (Dict) TableName() string {
	return "dc_dict"
}
