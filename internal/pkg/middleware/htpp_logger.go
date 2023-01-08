/*
 * @Author: silent-rain
 * @Date: 2023-01-08 00:47:40
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-09 01:02:01
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/middleware/htpp_logger.go
 * @Descripttion: 接口请求日志中间件，日志输出至数据库
 */
package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// HttpLogger 日志中间件
// 日志输出至数据库
func HttpLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		// 读取 body 数据
		bodyBytes, err := ctx.GetRawData()
		if err != nil {
			zap.S().Errorf("读取 body 失败, err: %v", err)
		} else {
			// gin body 只能获取一次，上面获取之后，一定要 再次给 context 赋值 不然 后面接口就获取不到了。
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
		}
		ctx.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("body", string(bodyBytes)),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
