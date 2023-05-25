/*WEB 日志
 */
package model

// WebLog WEB 日志
type WebLog struct {
	ID         uint   `json:"id" gorm:"column:id;primaryKey"`        // 自增ID
	UserId     uint   `json:"user_id" gorm:"column:user_id"`         // 用户ID
	Nickname   string `json:"nickname" gorm:"column:nickname"`       // 用户昵称
	TraceId    string `json:"trace_id" gorm:"column:trace_id"`       // 请求traceId
	OsType     uint   `json:"os_type" gorm:"column:os_type"`         // 终端类型: 0: 未知,1: 安卓,2 :ios,3 :web
	ErrorType  uint   `json:"error_type" gorm:"column:error_type"`   // 错误类型: 1:接口报错,2:代码报错
	Level      string `json:"level" gorm:"column:level"`             // 日志级别
	CallerLine string `json:"caller_line" gorm:"column:caller_line"` // 日发生位置
	Url        string `json:"url" gorm:"column:url"`                 // 错误页面
	Msg        string `json:"msg" gorm:"column:msg"`                 // 日志消息
	Stack      string `json:"stack" gorm:"column:stack"`             // 堆栈信息
	Note       string `json:"note" gorm:"column:note"`               // 备注
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`   // 创建时间
}

// TableName 表名重写
func (WebLog) TableName() string {
	return "log_web"
}
