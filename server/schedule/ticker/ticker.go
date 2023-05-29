// Package ticker 即时器任务
package ticker

import (
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/pkg/schedule/ticker"
)

// Init 即时器任务
func Init() {
	cfg := global.Instance().Config().Schedule
	// 添加任务
	ticker.Add(ticker.New("demo", 10, cfg.IsEnableTicker("enable_demo"), RegisterDemoPrintln))

	// 开始执行任务
	ticker.Start()
}
