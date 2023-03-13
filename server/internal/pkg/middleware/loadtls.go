/*
添加对 HTTPS 的支持

利用工具可以生成私钥key.pem和证书cert.pem

Golang标准库crypto/tls里有generate_cert.go，可以生成私钥key.pem和证书cert.pem，host参数是必须的，需要注意的是默认有效期是1年。

windows平台运行如下命令：
go.exe run  C:\Go\src\crypto\tls\generate_cert.go --host="localhost"

linux平台运行如下命令：
go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
*/
package middleware

import (
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// LoadTls 添加对 HTTPS 的支持
func LoadTls() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mid := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		// If there was an error, do not continue.
		if err := mid.Process(ctx.Writer, ctx.Request); err != nil {
			log.New(ctx).WithCode(errcode.LoadSSLError).
				WithStack().
				Panicf("%v", err)
			return
		}
		ctx.Next()
	}
}
