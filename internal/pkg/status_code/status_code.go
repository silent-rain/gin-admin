/*
 * @Author: silent-rain
 * @Date: 2023-01-07 16:35:07
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 18:14:01
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
	JsonDataEncodeError StatuScode = iota + 10200 // json 数据编码错误
	JsonDataDecodeError                           // json 数据解码错误
)

// 数据库
const (
	DbQueryError      StatuScode = iota + 10300 // 数据查询错误
	DbQueryEmptyError                           // 数据查询空
	DbDataExistError                            // 数据已存在
)

// 鉴权
const (
	TokenGenerateError    StatuScode = iota + 10400 // 生成 Token 失败
	TokenParsingError                               // 解析 Token 失败
	TokenInvalid                                    // 无效鉴权
	TokenExpiredSignature                           // 鉴权过期
)

// 上游服务

// 系统管理
const (
	UserRegisterError StatuScode = iota + 11000 // 用户注册失败
	UserLoginError                              // 用户登录失败
	UserLogoutError                             // 用户注销失败
	UserDisableError                            // 您的账号已被禁用,请联系管理员
)

// 状态码映射具体消息
var statusCodeMsg = map[StatuScode]error{
	Ok:            errors.New("Ok"),
	InternalError: errors.New("内部错误"),
	Unknown:       errors.New("未知错误"),
	// 请求解析
	ReqParameterParsingError: errors.New("请求参数解析错误"),
	// 数据解析
	JsonDataEncodeError: errors.New("数据编码错误"),
	JsonDataDecodeError: errors.New("数据解码错误"),
	// 数据库
	DbQueryError:      errors.New("数据查询错误"),
	DbQueryEmptyError: errors.New("数据查询空"),
	DbDataExistError:  errors.New("数据已存在"),
	// 鉴权
	TokenGenerateError:    errors.New("生成 Token 失败"),
	TokenParsingError:     errors.New("解析 Token 失败"),
	TokenInvalid:          errors.New("无效鉴权"),
	TokenExpiredSignature: errors.New("鉴权过期"),
	// 系统管理
	UserRegisterError: errors.New("用户注册失败"),
	UserLoginError:    errors.New("用户登录失败"),
	UserLogoutError:   errors.New("用户注销失败"),
	UserDisableError:  errors.New("您的账号已被禁用,请联系管理员"),
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
