/*用户API接口Token令牌表*/
package permission

// UserApiToken 用户API接口Token令牌表
type UserApiToken struct {
	ID         uint   `json:"id" gorm:"column:id;primaryKey"`      // 自增ID
	UserId     uint   `json:"sort" gorm:"column:sort"`             // 用户ID
	Token      string `json:"token" gorm:"column:token"`           // Token信息
	Permission string `json:"permission" gorm:"column:permission"` // 权限:GET,POST,PUT,DELETE
	Note       string `json:"note" gorm:"column:note"`             // 备注
	Status     uint   `json:"status" gorm:"column:status"`         // 状态,0:停用,1:启用
	CreatedAt  string `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

// TableName 表名重写
func (UserApiToken) TableName() string {
	return "perm_user_api_token"
}
