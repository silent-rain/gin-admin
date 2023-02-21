/*即时器, 周期性的定时器*/
package ticker

import (
	"runtime/debug"
	"time"

	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"
)

var tickerTasks []*ticker

// TickerFn 定义执行函数类型
type TickerFn func() error

// 定时器中的成员
type ticker struct {
	*time.Ticker          // 即时器
	name         string   // 任务名称
	interval     int      // 间隔周期
	runner       TickerFn // 任务函数
	enable       bool     // 是否启用
}

// Add 将即时器添加到任务列表中
func Add(ticke ...*ticker) {
	tickerTasks = append(tickerTasks, ticke...)
}

// Start 开始执行即时器任务列表
func Start() {
	for _, task := range tickerTasks {
		task.Start()
	}
}

// New 创建一个即时器
// name: 任务名称
// interval: 间隔周期(秒)
// fn: 回调函数
func New(name string, interval int, fn TickerFn, enable bool) *ticker {
	return &ticker{
		Ticker:   time.NewTicker(time.Duration(interval) * time.Second),
		name:     name,
		interval: interval,
		runner:   fn,
		enable:   enable,
	}
}

// 任务包装, 异常捕获
func (t *ticker) runTask() {
	defer func() {
		if err := recover(); err != nil {
			log.New(nil).
				WithCode(errcode.TickerPanicError).
				WithStack(debug.Stack()).
				Errorf("%v", err)
			return
		}

	}()

	if err := t.runner(); err != nil {
		log.New(nil).WithCode(errcode.TickerRunnerError).Errorf("%v", err)
	}
}

// Start 启动即时器需要执行的任务
func (t *ticker) Start() {
	if !t.enable {
		return
	}
	go func() {
		for range t.C {
			t.runTask()
		}
	}()
}
