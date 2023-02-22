/*即时器, 周期性的定时器*/
package ticker

import (
	"time"

	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"
)

var tickerTasks []*tickerTask

// TickerFn 定义即时器执行函数类型
type TickerFn func() error

// 即时器中的成员
type tickerTask struct {
	*time.Ticker          // 即时器
	name         string   // 任务名称
	interval     int      // 间隔周期
	runner       TickerFn // 任务函数
	enable       bool     // 是否启用
}

// Add 将即时器添加到任务列表中
func Add(ticke ...*tickerTask) {
	tickerTasks = append(tickerTasks, ticke...)
}

// Start 开始执行即时器任务列表
func Start() {
	for _, task := range tickerTasks {
		task.Start()
	}
}

// Stop 关闭即时器, 无法关闭已经开始的任务
func Stop() {
	for _, task := range tickerTasks {
		task.Stop()
	}
}

// New 创建一个即时器
// name: 任务名称
// interval: 间隔周期(秒)
// fn: 回调函数
func New(name string, interval int, enable bool, fn TickerFn) *tickerTask {
	return &tickerTask{
		Ticker:   time.NewTicker(time.Duration(interval) * time.Second),
		name:     name,
		interval: interval,
		runner:   fn,
		enable:   enable,
	}
}

// 任务包装, 异常捕获
func (t *tickerTask) runTask() {
	defer func() {
		if err := recover(); err != nil {
			log.New(nil).
				WithCode(errcode.TickerPanicError).
				WithStack().
				Errorf("%v", err)
		}
		log.New(nil).WithCode(errcode.Ok).Debugf("ticker end: %s, interval: %d", t.name, t.interval)
	}()
	log.New(nil).WithCode(errcode.Ok).Debugf("ticker start: %s, interval: %d", t.name, t.interval)
	if err := t.runner(); err != nil {
		log.New(nil).WithCode(errcode.TickerRunnerError).Errorf("%v", err)
	}
}

// Start 启动即时器需要执行的任务
func (t *tickerTask) Start() {
	if !t.enable {
		return
	}
	go func() {
		for range t.C {
			t.runTask()
		}
	}()
}
