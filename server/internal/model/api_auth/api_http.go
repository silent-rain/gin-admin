/*Http协议接口管理表*/
package apiauth

// ApiHttp Http协议接口管理表
type ApiHttp struct {
	ID        uint   `json:"id" gorm:"column:id;primaryKey"`      // 自增ID
	Name      string `json:"name" gorm:"column:name"`             // 接口名称
	Method    string `json:"method" gorm:"column:method"`         // 请求类型:GET,POST,PUT,DELETE
	Uri       string `json:"uri" gorm:"column:uri"`               // URI资源
	Note      string `json:"note" gorm:"column:note"`             // 备注
	Status    uint   `json:"status" gorm:"column:status"`         // 状态,0:停用,1:启用
	CreatedAt string `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

// TableName 表名重写
func (ApiHttp) TableName() string {
	return "api_http"
}
