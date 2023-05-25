// Package log 日志
package log

import (
	"context"
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

	ctx := context.Background()
	slog.LogAttrs(ctx, slog.LevelInfo, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}
