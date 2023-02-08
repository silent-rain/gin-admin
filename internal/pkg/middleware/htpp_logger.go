/*
 * @Author: silent-rain
 * @Date: 2023-01-08 00:47:40
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-12 22:05:28
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/middleware/htpp_logger.go
 * @Descripttion: 接口请求日志中间件，日志输出至数据库
 */
package middleware

import (
	"bytes"
	"io"
	"strings"
	"time"

	systemDao "gin-admin/internal/dao/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// HttpLogger 日志中间件
// 日志输出至数据库
func HttpLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 响应
		start := time.Now()
		// 读取 body 数据
		bodyBytes, err := ctx.GetRawData()
		if err != nil {
			log.New(ctx).Errorf("读取请求体失败, %v", err)
		} else {
			// gin body 只能获取一次，上面获取之后，一定要 再次给 context 赋值 不然 后面接口就获取不到了。
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
		}

		// record response info
		blw := &ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		htppLog := systemModel.HttpLog{
			UserId:     utils.GetUserId(ctx),
			TraceId:    utils.GetTraceId(ctx),
			StatusCode: ctx.Writer.Status(),
			Method:     ctx.Request.Method,
			Path:       ctx.Request.URL.Path,
			Query:      ctx.Request.URL.RawQuery,
			Body:       string(bodyBytes),
			RemoteAddr: ctx.ClientIP(),
			UserAgent:  ctx.Request.UserAgent(),
			Cost:       time.Since(start).Nanoseconds(),
			HttpType:   "REQ",
		}
		// 文件上传的 API 需要过滤 body
		if !strings.HasPrefix(ctx.Request.URL.Path, "/api/upload") ||
			!strings.HasPrefix(ctx.Request.URL.Path, "/upload") {
			htppLog.Body = ""
		}

		go func(htppLog systemModel.HttpLog) {
			systemDao.NewHttpLogDao().Add(htppLog)
		}(htppLog)

		ctx.Next()

		// 响应
		htppLog.StatusCode = ctx.Writer.Status()
		htppLog.Cost = time.Since(start).Nanoseconds()
		htppLog.HttpType = "RSP"
		htppLog.Body = blw.Body.String()
		go func(htppLog systemModel.HttpLog) {
			systemDao.NewHttpLogDao().Add(htppLog)
		}(htppLog)
	}
}

// 自定义响应接口
type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer // 缓存
}

// Write 写入 []byte
func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Write 写入 string
func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
