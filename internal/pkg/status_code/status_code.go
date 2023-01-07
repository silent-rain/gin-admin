/*
 * @Author: silent-rain
 * @Date: 2023-01-07 16:35:07
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 18:02:27
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/status_code/status_code.go
 * @Descripttion:业务状态码
 */
package statuscode

import "errors"

// 业务状态码
type StatuScode uint

const (
	Ok      StatuScode = iota + 10000 // 正常
	Unknown                           // 未知状态
)

// 状态码映射具体消息
var statusCodeMsg = map[StatuScode]error{
	Ok:      errors.New("Ok"),
	Unknown: errors.New("Unknown"),
}

// Error 返回状态码错误信息
func (r StatuScode) Error() error {
	msg, ok := statusCodeMsg[r]
	if !ok {
		return statusCodeMsg[Unknown]
	}
	return msg
}

// Msg 返回状态码信息
func (r StatuScode) Msg() string {
	msg, ok := statusCodeMsg[r]
	if !ok {
		return statusCodeMsg[Unknown].Error()
	}
	return msg.Error()
}
