/*日志
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
