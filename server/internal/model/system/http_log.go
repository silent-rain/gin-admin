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
	ID         uint      `json:"id" gorm:"column:id;primaryKey"`                           // 自增ID
	UserId     uint      `json:"user_id" gorm:"column:user_id"`                            // 请求用户ID
	TraceId    string    `json:"trace_id" gorm:"column:trace_id"`                          // 上游请求traceId
	SpanId     string    `json:"span_id" gorm:"column:span_id"`                            // 请求spanId
	StatusCode int       `json:"status_code" gorm:"column:status_code"`                    // 请求状态码
	Method     string    `json:"method" gorm:"column:method"`                              // 请求方法
	Path       string    `json:"path" gorm:"column:path"`                                  // 请求地址路径
	Query      string    `json:"query" gorm:"column:query"`                                // 请求参数
	Body       string    `json:"body" gorm:"column:body;"`                                 // 请求体/响应体
	RemoteAddr string    `json:"remote_addr" gorm:"column:remote_addr"`                    // 请求IP
	UserAgent  string    `json:"user_agent" gorm:"column:user_agent"`                      // 用户代理
	Cost       int64     `json:"cost" gorm:"column:cost"`                                  // 耗时,纳秒
	HttpType   string    `json:"htpp_type" gorm:"column:htpp_type"`                        // 请求类型:req/rsp
	Note       string    `json:"note" gorm:"column:note"`                                  // 备注
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
}

// TableName 表名重写
func (HttpLog) TableName() string {
	return "sys_http_log"
}
