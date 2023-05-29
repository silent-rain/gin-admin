// Package timer 定时器任务
package timer

import (
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/pkg/schedule/timer"
)

// Init 定时器任务
func Init() {
	cfg := global.Instance().Config().Schedule
	// 添加任务
	timer.Add(timer.New("demo", "*/5 * * * * ?", cfg.IsEnableTimer("enable_demo"), RegisterDemoPrintln))

	// 开始执行任务
	timer.Start()
}
