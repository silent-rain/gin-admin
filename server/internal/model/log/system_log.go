/*系统日志
 */
package logModel

// SystemLog 系统日志
type SystemLog struct {
	ID         uint   `json:"id" gorm:"column:id;primaryKey"`        // 自增ID
	UserId     uint   `json:"user_id" gorm:"column:user_id"`         // 用户ID
	Nickname   string `json:"nickname" gorm:"column:nickname"`       // 用户昵称
	TraceId    string `json:"trace_id" gorm:"column:trace_id"`       // 请求traceId
	SpanId     string `json:"span_id" gorm:"column:span_id"`         // 埋点spanId
	Level      string `json:"level" gorm:"column:level"`             // 日志级别
	CallerLine string `json:"caller_line" gorm:"column:caller_line"` // 日志发生位置
	ErrorCode  uint   `json:"error_code" gorm:"column:error_code"`   // 业务错误码
	ErrorMsg   string `json:"error_msg" gorm:"column:error_msg"`     // 业务错误信息
	Msg        string `json:"msg" gorm:"column:msg"`                 // 日志消息
	Stack      string `json:"stack" gorm:"column:stack"`             // 堆栈信息
	Extend     string `json:"extend" gorm:"column:extend"`           // 日志扩展信息/json
	Note       string `json:"note" gorm:"column:note"`               // 备注
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`   // 创建时间
}

// TableName 表名重写
func (SystemLog) TableName() string {
	return "sys_system_log"
}
