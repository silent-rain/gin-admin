// Package log 日志
package log

import (
	"testing"

	"github.com/silent-rain/gin-admin/global"

	"go.uber.org/zap"
)

func TestInit(t *testing.T) {
	global.Init()
	Init()
	zap.L().Debug("this is debug")
	zap.L().Info("this is info")
	zap.L().Warn("this is warn")
	zap.L().Error("this is error")
}
