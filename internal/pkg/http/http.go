/*HTTP 服务上下文封装*/
package http

import (
	"encoding/json"

	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// ParsingReqParams 将请求参数解析到结构体
func ParsingReqParams(ctx *gin.Context, req interface{}) *response.ResponseAPI {
	if err := ctx.ShouldBind(req); err != nil {
		log.New(ctx).WithCode(statuscode.ReqParameterParsingError).Errorf("%v", err)
		return response.New().WithCode(statuscode.ReqParameterParsingError)
	}
	return response.New()
}

// ApiJsonConvertJson 结构体转换
func ApiJsonConvertJson(ctx *gin.Context, src interface{}, dst interface{}) *response.ResponseAPI {
	bytes, err := json.Marshal(src)
	if err != nil {
		log.New(ctx).WithCode(statuscode.JsonDataEncodeError).Errorf("%v", err)
		return response.New().WithCode(statuscode.JsonDataEncodeError)
	}
	if err := json.Unmarshal(bytes, dst); err != nil {
		log.New(ctx).WithCode(statuscode.JsonDataDecodeError).Errorf("%v", err)
		return response.New().WithCode(statuscode.JsonDataDecodeError)
	}
	return response.New()
}
