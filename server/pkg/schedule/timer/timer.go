// Package timer 定时器
package timer

import (
	"time"

	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/robfig/cron/v3"
)

// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
var crontab = cron.New(cron.WithSeconds()) // 精确到秒

// Add 添加定时器
func Add(jobs ...*timerTask) {
	for _, job := range jobs {
		if !job.enable {
			continue
		}
		crontab.AddJob(job.spec, job)
	}
}

// Start 启动定时器
func Start() {
	crontab.Start()
}

// Stop 关闭定时器, 无法关闭已经开始的任务
func Stop() {
	crontab.Stop()
}

// TimerFn 定义定时器执行函数类型
type TimerFn func() error

// 定时器中的成员
type timerTask struct {
	*time.Ticker         // 即时器
	name         string  // 任务名称
	spec         string  // cron 表达式
	enable       bool    // 是否启用
	runner       TimerFn // 任务函数
}

// New 创建定时器
func New(name, spec string, enable bool, fn TimerFn) *timerTask {
	return &timerTask{
		name:   name,
		spec:   spec,
		enable: enable,
		runner: fn,
	}
}

// Run 开始执行任务
// 定时器是协程执行的, 不用再 go 出去
func (t *timerTask) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.New(nil).
				WithCode(errcode.TimerPanicError).
				WithStack().
				Errorf("%v", err)
		}
		log.New(nil).WithCode(errcode.Ok).Debugf("ticker end: %s, spec: %s", t.name, t.spec)
	}()
	log.New(nil).WithCode(errcode.Ok).Debugf("ticker start: %s, spec: %s", t.name, t.spec)
	if err := t.runner(); err != nil {
		log.New(nil).WithCode(errcode.TimerRunnerError).Errorf("%v", err)
	}
}
