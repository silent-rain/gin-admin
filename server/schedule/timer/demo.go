// Package timer 测试定时器任务
package timer

import (
	"fmt"
)

// RegisterDemoPrintln 注册日志打印任务
func RegisterDemoPrintln() error {
	fmt.Println("timer_task demo======")
	return nil
}
