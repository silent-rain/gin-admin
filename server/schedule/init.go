// Package schedule 定时任务
package schedule

import (
	"github.com/silent-rain/gin-admin/schedule/ticker"
	"github.com/silent-rain/gin-admin/schedule/timer"
)

// Init 初始化定时任务
func Init() {
	// 初始化即时器任务
	ticker.Init()
	// 初始化即时器任务
	timer.Init()
}
