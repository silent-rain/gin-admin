// Package timertask 测试定时器任务
package timertask

import (
	"fmt"
)

// RegisterDemoPrintln 注册日志打印任务
func RegisterDemoPrintln() error {
	fmt.Println("timer_task demo======")
	return nil
}
