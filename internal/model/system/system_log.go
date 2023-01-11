/*
 * @Author: silent-rain
 * @Date: 2023-01-12 00:00:25
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-12 00:04:35
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/model/system/system_log.go
 * @Descripttion: 系统日志
 */
package systemModel

import "time"

// SystemLog 系统日志
type SystemLog struct {
	ID         uint      `gorm:"column:id;primaryKey"`                   // 自增ID
	UserId     uint      `gorm:"column:user_id"`                         // 请求用户ID
	TraceId    string    `gorm:"column:trace_id"`                        // 请求traceId
	Level      string    `gorm:"column:level"`                           // 日志级别
	CallerLine string    `gorm:"column:caller_line"`                     // 日志发生位置
	ErrorCode  string    `gorm:"column:error_code"`                      // 业务错误码
	ErrorMsg   string    `gorm:"column:error_msg"`                       // 业务错误信息
	Msg        string    `gorm:"column:msg"`                             // 日志消息
	Extend     string    `gorm:"column:extend"`                          // 日志扩展信息/json
	Note       string    `gorm:"column:note"`                            // 备注
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
}

// TableName 表名重写
func (SystemLog) TableName() string {
	return "system_log"
}
