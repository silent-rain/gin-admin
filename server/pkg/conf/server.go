// Package conf 系统服务配置
package conf

import "fmt"

// ServerConfig 系统服务配置
type ServerConfig struct {
	// 服务配置
	Base struct {
		Address string `toml:"address"` // 服务地址
		Port    int    `toml:"port"`    // 服务端口
	} `toml:"base"`
	// 插件配置
	Plugin struct {
		EnableLogo           bool   `toml:"enable_logo"`             // 是否启用启动后显示 logo
		EnableSingleLogin    bool   `toml:"enable_single_login"`     // 是否启用单点登录
		EnableRateLimiter    bool   `toml:"enable_rate_limiter"`     // 是否启用限速
		MaxRequestsPerSecond int    `toml:"max_requests_per_second"` // 每秒最大请求量
		EnablePprof          bool   `toml:"enable_pprof"`            // 是否启用 pprof 性能剖析工具
		EnablePrometheus     bool   `toml:"enable_prometheus"`       // 是否启用 Prometheus 监控指标工具
		EnableRecordMetrics  bool   `toml:"enable_record_metrics"`   // 是否启用 记录指标
		EnableSwagger        bool   `toml:"enable_swagger"`          // 是否启用 swagger API 文档
		EnableOpenBrowser    bool   `toml:"enable_open_browser"`     // 是否启用服务启动后打开浏览器
		OpenBrowserUrl       string `toml:"open_browser_url"`        // 启动后在浏览器中打开的 URL
	} `toml:"plugin"`
	// 上传路径配置
	Upload struct {
		FilePath string `toml:"filepath"` // 上传路径
	} `toml:"upload"`
}

// ServerAddress 获取服务地址
func (s ServerConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%d", s.Base.Address, s.Base.Port)
}
