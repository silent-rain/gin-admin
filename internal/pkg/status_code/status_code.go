/*
 * @Author: silent-rain
 * @Date: 2023-01-07 16:35:07
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 15:03:17
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/status_code/status_code.go
 * @Descripttion:业务状态码
 */
package statuscode

import "errors"

// 业务状态码
type StatuScode uint

const (
	Ok            StatuScode = iota + 10000 // 访问正常
	InternalError                           // 内部错误
	Unknown                                 // 未知错误
)

// 请求解析
const (
	ReqParameterParsingError StatuScode = iota + 10100 // 请求参数解析错误
)

// 数据解析
const (
	JsonDataEncodeError StatuScode = iota + 10100 // json 数据编码错误
	JsonDataDecodeError                           // json 数据解码错误
)

// 数据库

// 上游服务

// 系统管理
const (
	UserRegisterError StatuScode = iota + 10800 // 用户注册失败
)

// 状态码映射具体消息
var statusCodeMsg = map[StatuScode]error{
	Ok:            errors.New("Ok"),
	InternalError: errors.New("内部错误"),
	Unknown:       errors.New("未知错误"),
	// 请求解析
	ReqParameterParsingError: errors.New("请求参数解析错误"),
	// 系统管理
	UserRegisterError: errors.New("用户注册失败"),
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
