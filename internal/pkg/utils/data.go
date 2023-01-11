/*
 * @Author: silent-rain
 * @Date: 2023-01-08 15:34:09
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-11 21:34:32
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
)

// ParsingReqParams 将请求参数解析到结构体
func ParsingReqParams(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBind(&req); err != nil {
		response.New(ctx).WithCode(statuscode.ReqParameterParsingError).Json()
		return err
	}
	return nil
}

// ApiJsonConvertJson 结构体转换
func ApiJsonConvertJson(ctx *gin.Context, req interface{}, user interface{}) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		response.New(ctx).WithCode(statuscode.JsonDataEncodeError).Json()
		return err
	}
	if err := json.Unmarshal(bytes, user); err != nil {
		response.New(ctx).WithCode(statuscode.JsonDataDecodeError).Json()
		return err
	}
	return nil
}

// IndexOfArray 元素在字符串切片中的位置
func IndexOfArray(arr []string, target string) int {
	for i, item := range arr {
		if item == target {
			return i
		}
	}
	return -1
}
