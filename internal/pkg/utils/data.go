/*
 * @Author: silent-rain
 * @Date: 2023-01-08 15:34:09
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:33:08
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/data.go
 * @Descripttion: 数据处理工具
 */
package utils

import (
	"encoding/json"

	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ParsingReqParams 将请求参数解析到结构体
func ParsingReqParams(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBind(&req); err != nil {
		zap.S().Errorf("参数解析失败, err: %v", err)
		response.New(ctx).WithCode(statuscode.ReqParameterParsingError).Json()
		return err
	}
	return nil
}

// JsonConvertJson 结构体转换
func JsonConvertJson(ctx *gin.Context, req interface{}, user interface{}) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		zap.S().Errorf("数据编码失败, data: %v, err: %v", req, err)
		response.New(ctx).WithCode(statuscode.JsonDataEncodeError).Json()
		return err
	}
	if err := json.Unmarshal(bytes, user); err != nil {
		zap.S().Errorf("数据解码失败, data: %v, err: %v", req, err)
		response.New(ctx).WithCode(statuscode.JsonDataDecodeError).Json()
		return err
	}
	return nil
}
