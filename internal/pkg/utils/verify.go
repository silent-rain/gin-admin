/*
 * @Author: silent-rain
 * @Date: 2023-01-10 01:25:57
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-10 01:26:40
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/verify.go
 * @Descripttion: 验证工具
 */
package utils

import (
	"strings"

	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// 验证 Content-Type 是否为 json
func VerifyContentTypeJson(ctx *gin.Context) error {
	contentType := ctx.Request.Header.Get("Content-Type")
	if strings.ToLower(contentType) != "application/json" {
		ctx.Next()
		return statuscode.ReqContentTypeNotJson.Error()
	}
	return nil
}
