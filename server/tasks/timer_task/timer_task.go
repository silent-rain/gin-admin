/*定时器任务*/
package timerTask

import (
	"gin-admin/internal/pkg/conf"
	"gin-admin/pkg/cron/timer"
)

// Init 定时器任务
func Init() {
	cfg := conf.Instance().Tasks
	// 添加任务
	timer.Add(timer.New("demo", "*/5 * * * * ?", cfg.IsEnableTimer("enable_demo"), RegisterDemoPrintln))

	// 开始执行任务
	timer.Start()
}
