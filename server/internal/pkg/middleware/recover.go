/*Http 服务异常处理中间件*/
package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"gin-admin/internal/pkg/response"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recover 异常恢复
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误堆栈信息
				zap.S().Error("got panic",
					zap.String("panic", fmt.Sprintf("%+v", err)),
					zap.String("stack", string(debug.Stack())),
				)
				response.New(ctx).WithHttpStatus(http.StatusInternalServerError).
					WithCode(errcode.InternalServerError).Json()
			}
		}()

		// 加载完 defer recover，继续后续接口调用
		ctx.Next()
	}
}
