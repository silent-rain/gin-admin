/*Http 服务异常处理中间件*/
package middleware

import (
	"net/http"

	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Recover 异常恢复
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误堆栈信息
				log.New(ctx).WithCode(errcode.InternalServerError).
					WithStack().
					Errorf("%v", err)
				response.New(ctx).WithHttpStatus(http.StatusInternalServerError).
					WithCode(errcode.InternalServerError).Json()
			}
		}()

		// 加载完 defer recover，继续后续接口调用
		ctx.Next()
	}
}
