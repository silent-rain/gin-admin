// Package conf 系统环境配置
package conf

import "github.com/gin-gonic/gin"

// EnvironmentConfig 系统环境配置
type EnvironmentConfig struct {
	Env string `toml:"env"` // 系统环境配置: prod/test/dev/embed
}

// GinMode 将当前配置的环境转换为 gin 的环境
func (r EnvironmentConfig) GinMode() string {
	var mode = gin.DebugMode
	switch r.Env {
	case "prod":
		mode = gin.ReleaseMode
	case "test":
		mode = gin.TestMode
	case "dev":
		mode = gin.DebugMode
	case "embed":
		mode = gin.DebugMode
	default:
		mode = gin.DebugMode
	}
	return mode
}
