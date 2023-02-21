/*定时任务*/
package tasks

import tickerTask "gin-admin/tasks/ticker_task"

// Init 初始化定时任务
func Init() {
	// 初始化即时器任务
	tickerTask.Init()
}
