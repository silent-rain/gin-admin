// Package log 数据库日志
package log

import (
	"encoding/json"

	"github.com/silent-rain/gin-admin/internal/app/log/dao"
	"github.com/silent-rain/gin-admin/internal/app/log/model"

	"go.uber.org/zap"
)

// 数据库异步日志
type dbAsyncer struct {
	*zap.Logger
}

// 创建数据库异步日志对象
func newDbAsyncer() *dbAsyncer {
	base := &dbAsyncer{
		new(zap.Logger),
	}
	return base
}

// Write 定义Write方法以实现Sink接口
func (d dbAsyncer) Write(p []byte) (int, error) {
	sysLog := model.SystemLog{}
	if err := json.Unmarshal(p, &sysLog); err != nil {
		return len(p), err
	}
	go func() {
		dao.NewSystemLogDao().Add(sysLog)
	}()
	// 返回写入日志的长度,以及错误
	return len(p), nil
}

// Close 定义Close方法以实现Sink接口
func (d dbAsyncer) Close() error {
	// 涉及不到关闭对象的问题, 所以return就可以
	return nil
}

// Sync 定义Sync方法以实现Sink接口
func (d *dbAsyncer) Sync() error {
	// 涉及不到缓存同步问题, 所以return就可以
	return nil
}
