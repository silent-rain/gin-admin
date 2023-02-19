/*
 * @Author: silent-rain
 * @Date: 2023-01-07 22:02:42
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 22:30:59
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/log/log_test.go
 * @Descripttion: 日志
 */
package log

import (
	"net"
	"testing"

	"golang.org/x/exp/slog"
)

func TestInit(t *testing.T) {
	Init()
	slog.Debug("this is debug")
	slog.Info("this is info")
	slog.Warn("this is warn")
	slog.Error("this is error", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.LevelInfo, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}
