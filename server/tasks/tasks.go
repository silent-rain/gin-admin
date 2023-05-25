// Package tasks 定时任务
package tasks

import (
	tickerTask "github.com/silent-rain/gin-admin/tasks/ticker_task"
	timerTask "github.com/silent-rain/gin-admin/tasks/timer_task"
)

// Init 初始化定时任务
func Init() {
	// 初始化即时器任务
	tickerTask.Init()
	// 初始化即时器任务
	timerTask.Init()
}
