/*关闭资源*/
package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/repository/mysql"
	"gin-admin/pkg/color"
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

// WithCloseMysql 关闭 Mysql 服务
func WithCloseMysql() {
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
}

// WithCloseInfo 服务关闭后的消息提示
func WithCloseInfo() {
	fmt.Println(color.Blue("see you again~"))
}
