/*测试即时器*/
package tickerTask

import (
	"fmt"
	"gin-admin/internal/pkg/conf"
)

// RegisterDemoPrintln 注册日志打印任务
func RegisterDemoPrintln() error {
	if !conf.Instance().Tasks.IsEnableTicker("enable_demo") {
		return nil
	}
	fmt.Println("ticker_task demo======")
	return nil
}
