/*用户登录表*/
package system

// UserLogin 用户登录表-用于登录
type UserLogin struct {
	ID         uint   `json:"id" gorm:"column:id;primaryKey"`        // 自增ID
	UserId     uint   `json:"user_id" gorm:"column:user_id"`         // 用户ID
	Nickname   string `json:"nickname" gorm:"column:nickname"`       // 用户昵称
	RemoteAddr string `json:"remote_addr" gorm:"column:remote_addr"` // 登录IP
	UserAgent  string `json:"user_agent" gorm:"column:user_agent"`   // 用户代理
	Status     uint   `json:"status" gorm:"column:status"`           // 登录状态,0:停用,1:启用
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`   // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at"`   // 更新时间
}

// TableName 表名重写
func (UserLogin) TableName() string {
	return "sys_user_login"
}
