// Package ticker 测试即时器
package ticker

import (
	"fmt"
)

// RegisterDemoPrintln 注册日志打印任务
func RegisterDemoPrintln() error {
	fmt.Println("ticker_task demo======")
	return nil
}
