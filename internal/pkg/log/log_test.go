/*
 * @Author: silent-rain
 * @Date: 2023-01-07 22:02:42
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:16:19
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/log/log_test.go
 * @Descripttion: 日志
 */
package log

import (
	"testing"

	"gin-admin/internal/pkg/conf"

	"go.uber.org/zap"
)

func TestInit(t *testing.T) {
	conf.InitLoadConfig("../../../cmd/conf.yaml")
	Init()
	zap.L().Debug("this is debug")
	zap.L().Info("this is info")
	zap.L().Warn("this is warn")
	zap.L().Error("this is error")
}
