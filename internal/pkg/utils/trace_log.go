/*
 * @Author: silent-rain
 * @Date: 2023-01-10 21:29:22
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-10 22:13:10
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/trace_log.go
 * @Descripttion:  trace 日志链路
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"gin-admin/internal/pkg/conf"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateTraceId 生成 traceId
func GenerateTraceId(ctx *gin.Context) string {
	data := fmt.Sprintf(`unix_nano: %v
					rand: %v,
					status: %v,
					method: %v,
					path: %v,
					query: %v,
					remote_addr: %v,
					user_agent: %v`,
		time.Now().UnixNano(),
		rand.Int63n(5000000),
		ctx.Writer.Status(),
		ctx.Request.URL.Path,
		ctx.Request.URL.Path,
		ctx.Request.URL.RawQuery,
		ctx.ClientIP(),
		ctx.Request.UserAgent(),
	)
	m := md5.New()
	m.Write([]byte(conf.Secret))
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

// GetTraceId 获取请求 TraceTd
func GetTraceId(ctx *gin.Context) string {
	var traceTd = ctx.Request.Header.Get(GinContextTraceTd)
	if traceTd == "" {
		traceTd = GenerateTraceId(ctx)
	}
	return traceTd
}
