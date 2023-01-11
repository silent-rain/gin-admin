/*
 * @Author: silent-rain
 * @Date: 2023-01-10 01:25:57
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-11 21:57:33
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/verify.go
 * @Descripttion: 验证工具
 */
package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// LoginVerifyContentTypeJson 验证 Content-Type 是否为 json
// 仅验证 API
func LoginVerifyContentTypeJson(ctx *gin.Context) bool {
	contentType := ctx.Request.Header.Get("Content-Type")
	if strings.HasPrefix(ctx.Request.URL.Path, "/api") && strings.ToLower(contentType) != "application/json" {
		return false
	}
	return true
}

// VerifyContentTypeJson 验证 Content-Type 是否为 json
// 仅验证 API
func VerifyContentTypeJson(ctx *gin.Context) bool {
	contentType := ctx.Request.Header.Get("Content-Type")
	if strings.HasPrefix(ctx.Request.URL.Path, "/api") && strings.ToLower(contentType) != "application/json" {
		return false
	}
	return true
}
