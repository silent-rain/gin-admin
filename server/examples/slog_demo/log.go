/*
 * @Author: silent-rain
 * @Date: 2023-01-07 22:02:42
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 23:13:58
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/log/log.go
 * @Descripttion: 日志
 */
package log

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

// 初始化日志输出配置
func Init() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	slog.SetDefault(slog.New(opts.NewTextHandler(os.Stderr)))
	slog.SetDefault(slog.New(NewDbHandler()))
}

// DbHandler 数据库日志结构
type DbHandler struct {
	slog.Handler
	buf *bytes.Buffer
}

// Enabled 设置日志级别
func (h *DbHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

// Handle 句柄处理记录
func (h *DbHandler) Handle(ctx context.Context, r slog.Record) error {
	err := h.Handler.Handle(ctx, r)
	if err != nil {
		return err
	}
	var nb = make([]byte, h.buf.Len())
	copy(nb, h.buf.Bytes())
	fmt.Printf("========== %#v\n", string(nb))
	h.buf.Reset()
	return nil
}

// WithAttrs 添加属性
func (h *DbHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.Handler = h.Handler.WithAttrs(attrs)
	return h
}

// WithGroup 添加分组名称
func (h *DbHandler) WithGroup(name string) slog.Handler {
	h.Handler = h.Handler.WithGroup(name)
	return h
}

// 新建一个 Handler 对象
func NewDbHandler() *DbHandler {
	var b = make([]byte, 256)
	h := &DbHandler{
		buf: bytes.NewBuffer(b),
	}

	h.Handler = slog.NewJSONHandler(h.buf)
	return h
}

// slog.SetDefault(slog.New(NewDbHandler(ch).WithAttrs(attrs)))
