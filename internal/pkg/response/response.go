/*
 * @Author: silent-rain
 * @Date: 2023-01-07 17:45:55
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 19:16:38
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/response/response.go
 * @Descripttion: API 返回结构
 */
package response

import (
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// ResponseAPI API响应结构
type ResponseAPI struct {
	Code    statuscode.StatuScode `json:"code"` // 状态码
	Msg     string                `json:"msg"`  // 状态码信息
	Data    interface{}           `json:"data"` // 返回数据
	context *gin.Context          `json:"-"`    // gin Context
}

// 返回 API 响应结构对象
//
// 返回默认 Ok 状态码及对应的状态码信息
func New(c *gin.Context) *ResponseAPI {
	return &ResponseAPI{
		context: c,
		Code:    statuscode.Ok,
		Msg:     statuscode.Ok.Msg(),
	}
}

// WithMsg 添加响应状态码及状态码对应的信息
func (r *ResponseAPI) WithCode(code statuscode.StatuScode) *ResponseAPI {
	r.Code = code
	r.Msg = code.Msg()
	return r
}

// WithMsg 添加响应信息
func (r *ResponseAPI) WithMsg(msg string) *ResponseAPI {
	r.Msg = msg
	return r
}

// WithMsg 添加响应数据
func (r *ResponseAPI) WithData(data interface{}) *ResponseAPI {
	r.Data = data
	return r
}

// Json gin json 返回值
func (r *ResponseAPI) Json() {
	r.context.JSON(200, r)
}
