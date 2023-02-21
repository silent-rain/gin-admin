/*接口请求日志中间件，日志输出至数据库
 */
package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"time"

	systemDAO "gin-admin/internal/dao/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/core"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"

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
			UserId:     core.GetContext(ctx).UserId,
			Nickname:   core.GetContext(ctx).Nickname,
			TraceId:    core.GetContext(ctx).TraceId,
			SpanId:     core.GetContext(ctx).SpanId,
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
			systemDAO.NewHttpLogDao().Add(htppLog)
		}(htppLog)

		ctx.Next()

		// 响应
		htppLog.StatusCode = ctx.Writer.Status()
		htppLog.Cost = time.Since(start).Nanoseconds()
		htppLog.HttpType = "RESP"

		// 判断是否为接口, 当为接口时记录返回信息
		if err := json.Unmarshal(blw.Body.Bytes(), &response.ResponseAPI{}); err == nil {
			htppLog.Body = blw.Body.String()
		}
		go func(htppLog systemModel.HttpLog) {
			systemDAO.NewHttpLogDao().Add(htppLog)
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
