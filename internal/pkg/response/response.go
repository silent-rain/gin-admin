/*API 返回结构
 */
package response

import (
	"net/http"

	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// ResponseAPI API响应结构
type ResponseAPI struct {
	Code    statuscode.StatusCode `json:"code"` // 状态码
	Msg     string                `json:"msg"`  // 状态码信息
	Data    interface{}           `json:"data"` // 返回数据
	context *gin.Context          `json:"-"`    // gin Context
}

// New 返回 API 响应结构对象
//
// 返回默认 Ok 状态码及对应的状态码信息
func New(c *gin.Context) *ResponseAPI {
	return &ResponseAPI{
		context: c,
		Code:    statuscode.Ok,
		Msg:     statuscode.Ok.Msg(),
	}
}

// WithCode 添加响应状态码及状态码对应的信息
func (r *ResponseAPI) WithCode(code statuscode.StatusCode) *ResponseAPI {
	r.Code = code
	r.Msg = code.Msg()
	return r
}

// WithMsg 添加响应信息
func (r *ResponseAPI) WithMsg(msg string) *ResponseAPI {
	r.Msg = msg
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

// 返回状态码错误
func (r *ResponseAPI) Error() error {
	return r.Code.Error()
}

// Json 正常返回值
func (r *ResponseAPI) Json() {
	r.context.JSON(http.StatusOK, r)
}
