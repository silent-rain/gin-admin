/*关闭资源*/
package shutdown

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/repository/mysql"
	"gin-admin/pkg/color"
	"gin-admin/pkg/cron/ticker"
	"gin-admin/pkg/cron/timer"
	"gin-admin/pkg/errcode"
)

// Hook a graceful shutdown hook, default with signals of SIGINT and SIGTERM
type Hook interface {
	// WithSignals add more signals into hook
	WithSignals(signals ...syscall.Signal) Hook

	// Close register shutdown handles
	Close(funcs ...func())
}

type hook struct {
	ctx chan os.Signal
}

// NewHook create a Hook instance
func NewHook() Hook {
	hook := &hook{
		ctx: make(chan os.Signal, 1),
	}

	return hook.WithSignals(syscall.SIGINT, syscall.SIGTERM)
}

func (h *hook) WithSignals(signals ...syscall.Signal) Hook {
	for _, s := range signals {
		signal.Notify(h.ctx, s)
	}

	return h
}

func (h *hook) Close(funcs ...func()) {
	<-h.ctx
	signal.Stop(h.ctx)

	for _, f := range funcs {
		f()
	}
}

// WithCloseHttpServer 关闭 Http 服务
func WithCloseHttpServer(server *http.Server) func() {
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.New(nil).WithCode(errcode.HttpServerCloseError).Error("")
		}
		fmt.Println(color.Blue("server closed ..."))
	}
}

// WithCloseCron 关闭定时任务
func WithCloseCron() func() {
	return func() {
		ticker.Stop()
		timer.Stop()
		fmt.Println(color.Blue("cron closed ..."))
	}
}

// WithCloseMysql 关闭 Mysql 服务
func WithCloseMysql() func() {
	return func() {
		db := mysql.Instance()
		if db == nil {
			return
		}

		if err := db.DbWClose(); err != nil {
			log.New(nil).WithCode(errcode.DBWriteCloseError).Error("")
		}
		if err := db.DbRClose(); err != nil {
			log.New(nil).WithCode(errcode.DBReadCloseError).Error("")
		}
		fmt.Println(color.Blue("mysql closed ..."))
	}
}

// WithCloseInfo 服务关闭后的消息提示
func WithCloseInfo() func() {
	return func() {
		fmt.Println(color.Blue("see you again~"))
	}
}