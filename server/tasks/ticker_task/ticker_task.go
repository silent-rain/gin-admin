/*即时器任务*/
package tickertask

import (
	"gin-admin/internal/pkg/conf"
	"gin-admin/pkg/cron/ticker"
)

// Init 即时器任务
func Init() {
	cfg := conf.Instance().Tasks
	// 添加任务
	ticker.Add(ticker.New("demo", 10, cfg.IsEnableTicker("enable_demo"), RegisterDemoPrintln))

	// 开始执行任务
	ticker.Start()
}
