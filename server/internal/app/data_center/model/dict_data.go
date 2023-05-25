/*字典数据管理*/
package model

// DictData 字典数据表
type DictData struct {
	ID        uint   `json:"id" gorm:"column:id;primaryKey"`      // 字典项ID
	DictId    uint   `json:"dict_id" gorm:"column:dict_id"`       // 字典维度ID
	Name      string `json:"name" gorm:"column:name"`             // 字典项名称
	Value     string `json:"value" gorm:"column:value"`           // 字典项值
	Note      string `json:"note" gorm:"column:note"`             // 备注
	Status    uint   `json:"status" gorm:"column:status"`         // 状态,0:停用,1:启用
	CreatedAt string `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

// TableName 表名重写
func (DictData) TableName() string {
	return "dc_dict_data"
}
