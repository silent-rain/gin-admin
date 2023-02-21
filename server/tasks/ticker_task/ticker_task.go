/*即时器*/
package tickerTask

import "gin-admin/pkg/cron/ticker"

// Init 即时器任务
func Init() {
	// 添加任务
	ticker.Add(ticker.New("demo", 10, RegisterDemoPrintln))

	// 开始执行任务
	ticker.Start()
}
