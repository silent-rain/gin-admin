/*错误封装*/
package code_errors

// CError 自定义错误类型
type CError struct {
	Code StatusCode `json:"code"` // 状态码
	Msg  string     `json:"msg"`  // 状态码信息
}

// New 返回自定义错误类型对象
func New(code StatusCode) *CError {
	return &CError{
		Code: code,
		Msg:  code.Msg(),
	}
}

// WithMsg 添加返回信息
func (e *CError) WithMsg(msg string) *CError {
	e.Msg = msg
	return e
}

// Error 实现 error 接口
func (e *CError) Error() string {
	return e.Msg
}
