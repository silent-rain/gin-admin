/*HTTP 服务上下文封装*/
package http

import (
	"encoding/json"

	"gin-admin/internal/pkg/code_errors"
	"gin-admin/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// ParsingReqParams 将请求参数解析到结构体
func ParsingReqParams(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBind(req); err != nil {
		log.New(ctx).WithCode(code_errors.ReqParameterParsingError).Errorf("%v", err)
		return code_errors.New(code_errors.ReqParameterParsingError)
	}
	return nil
}

// ApiJsonConvertJson 结构体转换
func ApiJsonConvertJson(ctx *gin.Context, src interface{}, dst interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		log.New(ctx).WithCode(code_errors.JsonDataEncodeError).Errorf("%v", err)
		return code_errors.New(code_errors.JsonDataEncodeError)
	}
	if err := json.Unmarshal(bytes, dst); err != nil {
		log.New(ctx).WithCode(code_errors.JsonDataDecodeError).Errorf("%v", err)
		return code_errors.New(code_errors.JsonDataDecodeError)
	}
	return nil
}
