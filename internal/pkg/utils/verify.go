/*
 * @Author: silent-rain
 * @Date: 2023-01-10 01:25:57
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-10 21:59:07
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/verify.go
 * @Descripttion: 验证工具
 */
package utils

import (
	"strings"

	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// LoginVerifyContentTypeJson 验证 Content-Type 是否为 json
// 仅验证 API
func LoginVerifyContentTypeJson(ctx *gin.Context) {
	contentType := ctx.Request.Header.Get("Content-Type")
	if strings.HasPrefix(ctx.Request.URL.Path, "/api") && strings.ToLower(contentType) != "application/json" {
		response.New(ctx).WithCode(statuscode.ReqContentTypeNotJson).Json()
		ctx.Abort()
		return
	}
}

// VerifyContentTypeJson 验证 Content-Type 是否为 json
// 仅验证 API
func VerifyContentTypeJson(ctx *gin.Context) error {
	contentType := ctx.Request.Header.Get("Content-Type")
	if strings.HasPrefix(ctx.Request.URL.Path, "/api") && strings.ToLower(contentType) != "application/json" {
		return statuscode.ReqContentTypeNotJson.Error()
	}
	return nil
}
