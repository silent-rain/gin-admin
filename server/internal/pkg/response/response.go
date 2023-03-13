/*API 返回结构
 */
package response

import (
	"net/http"

	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// ResponseAPI API响应结构
type ResponseAPI struct {
	Code       errcode.ErrorCode `json:"code"` // 状态码
	Msg        string            `json:"msg"`  // 状态码信息
	Data       interface{}       `json:"data"` // 返回数据
	httpStatus int               `json:"-"`    // HTTP 状态码
	ctx        *gin.Context      `json:"-"`
}

// New 返回 API 响应结构对象
// 返回默认 Ok 状态码及对应的状态码信息
func New(ctx *gin.Context) *ResponseAPI {
	return &ResponseAPI{
		Code:       errcode.Ok,
		Msg:        errcode.Ok.Error(),
		ctx:        ctx,
		httpStatus: http.StatusOK,
	}
}

// WithHttpStatus 添加请求状态码
func (r *ResponseAPI) WithHttpStatus(code int) *ResponseAPI {
	r.httpStatus = code
	return r
}

// WithMsg 添加返回信息
func (r *ResponseAPI) WithMsg(msg string) *ResponseAPI {
	r.Msg = msg
	return r
}

// WithCode 添加响应状态码及状态码对应的信息
func (r *ResponseAPI) WithCode(code errcode.ErrorCode) *ResponseAPI {
	r.Code = code
	r.Msg = code.Error()
	return r
}

// WithCodeError 添加响应状态码及状态码对应的信息
func (r *ResponseAPI) WithError(err error) *ResponseAPI {
	// 业务错误码 error code
	if code, ok := err.(errcode.ErrorCode); ok {
		r.Code = code
		r.Msg = code.Error()
		return r
	}

	// 业务错误码附加信息 erro code
	if msg, ok := err.(*errcode.ErrorMsg); ok {
		r.Code = msg.Code
		r.Msg = msg.Err.Error()
		return r
	}

	// 原始错误
	r.Code = errcode.UnknownError
	r.Msg = err.Error()
	return r
}

// WithData 添加响应数据
func (r *ResponseAPI) WithData(data interface{}) *ResponseAPI {
	r.Data = data
	return r
}

// WithData 添加列表响应数据及列表总数
func (r *ResponseAPI) WithDataList(data interface{}, total int64) *ResponseAPI {
	r.Data = map[string]interface{}{
		"data_list": data,
		"tatol":     total,
	}
	return r
}

// Json 返回接口
func (r *ResponseAPI) Json() {
	r.ctx.JSON(r.httpStatus, r)
}
