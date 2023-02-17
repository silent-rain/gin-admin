/*HTTP 服务上下文封装*/
package http

import (
	"encoding/json"

	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// ParsingReqParams 将请求参数解析到结构体
func ParsingReqParams(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBind(req); err != nil {
		log.New(ctx).WithCode(errcode.ReqParameterParsingError).Errorf("%v", err)
		return errcode.New(errcode.ReqParameterParsingError)
	}
	return nil
}

// ApiJsonConvertJson 结构体转换
func ApiJsonConvertJson(ctx *gin.Context, src interface{}, dst interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		log.New(ctx).WithCode(errcode.JsonDataEncodeError).Errorf("%v", err)
		return errcode.New(errcode.JsonDataEncodeError)
	}
	if err := json.Unmarshal(bytes, dst); err != nil {
		log.New(ctx).WithCode(errcode.JsonDataDecodeError).Errorf("%v", err)
		return errcode.New(errcode.JsonDataDecodeError)
	}
	return nil
}
