// Package conf 系统环境配置
package conf

import "github.com/gin-gonic/gin"

// EnvironmentConfig 系统环境配置
type EnvironmentConfig struct {
	Env string `toml:"env"` // 系统环境配置: prod/test/dev
}

// Active 当前配置的环境
func (r EnvironmentConfig) Active() string {
	var mode = gin.DebugMode
	switch r.Env {
	case "prod":
		mode = gin.ReleaseMode
	case "test":
		mode = gin.TestMode
	case "debug":
		mode = gin.DebugMode
	}
	return mode
}
