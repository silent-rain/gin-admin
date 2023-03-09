/*用户API接口Token令牌表*/
package permission

// UserApiToken 用户API接口Token令牌表
type UserApiToken struct {
	ID         uint   `json:"id" gorm:"column:id;primaryKey"`      // 自增ID
	UserId     uint   `json:"user_id" gorm:"column:user_id"`       // 用户ID
	Nickname   string `json:"nickname" gorm:"nickname"`            // 用户昵称
	Permission string `json:"permission" gorm:"column:permission"` // 权限标识:GET,POST,PUT,DELETE
	Token      string `json:"token" gorm:"column:token"`           // 令牌
	Passphrase string `json:"passphrase" gorm:"column:passphrase"` // 口令
	Note       string `json:"note" gorm:"column:note"`             // 备注
	Status     uint   `json:"status" gorm:"column:status"`         // 状态,0:停用,1:启用
	CreatedAt  string `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}
