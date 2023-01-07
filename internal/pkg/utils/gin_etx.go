/*
 * @Author: silent-rain
 * @Date: 2023-01-06 01:22:41
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 21:28:41
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/gin_etx.go
 * @Descripttion:
 */
/**gin 服务扩展工具*/
package utils

import (
	"fmt"
	"io"
	"os"

	"gin-admin/internal/pkg/conf"

	"github.com/gin-gonic/gin"
)

func init() {
	ginLogFile()
}

// 将 gin 服务产生的日志输出至文件
func ginLogFile() {
	// debug 模式输出至控制台
	if conf.Instance().EnvConfig.Env() == gin.DebugMode {
		return
	}

	f, err := os.OpenFile("../logs/server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("服务日志文件打开失败, err: %v", err))
	}
	// 把日志信息输出到f文件中
	gin.DefaultWriter = io.MultiWriter(f)
	// 把错误信息也输出到f文件中
	gin.DefaultErrorWriter = io.MultiWriter(f)
}
