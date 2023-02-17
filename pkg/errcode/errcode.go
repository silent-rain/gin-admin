/*错误封装*/
package errcode

// Error 自定义错误类型
type Error struct {
	Code ErrorCode `json:"code"` // 状态码
	Msg  string    `json:"msg"`  // 状态码信息
}

// New 返回自定义错误类型对象
func New(code ErrorCode) *Error {
	return &Error{
		Code: code,
		Msg:  code.Msg(),
	}
}

// WithMsg 添加返回信息
func (e *Error) WithMsg(msg string) *Error {
	e.Msg = msg
	return e
}

// Error 实现 error 接口
func (e *Error) Error() string {
	return e.Msg
}
