/*
 * @Author: silent-rain
 * @Date: 2023-01-09 23:09:04
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 22:49:29
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/model/system/http_log.go
 * @Descripttion: 网络请求日志模型
 */
package systemModel

import (
	"time"
)

// HttpLog 网络请求日志
type HttpLog struct {
	ID         uint      `gorm:"column:id;primaryKey"`                   // 自增ID
	UserId     uint      `gorm:"column:user_id"`                         // 请求用户ID
	TraceId    string    `gorm:"column:trace_id"`                        // 请求traceId
	StatusCode int       `gorm:"column:status_code"`                     // 请求状态码
	Method     string    `gorm:"column:method"`                          // 请求方法
	Path       string    `gorm:"column:path"`                            // 请求地址路径
	Query      string    `gorm:"column:query"`                           // 请求参数
	Body       string    `gorm:"column:body"`                            // 请求体/响应体
	RemoteAddr string    `gorm:"column:remote_addr"`                     // 请求IP
	UserAgent  string    `gorm:"column:user_agent"`                      // 用户代理
	Cost       int64     `gorm:"column:cost"`                            // 耗时,纳秒
	HttpType   string    `gorm:"column:htpp_type"`                       // 日志类型:req/rsp
	Note       string    `gorm:"column:note"`                            // 备注
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
}

// TableName 表名重写
func (HttpLog) TableName() string {
	return "sys_http_log"
}
