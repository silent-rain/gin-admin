// Package conf 配置
package conf

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

var (
	once   sync.Once
	config *Config
	// ConfigFile 配置文件路径
	ConfigFile = "./conf.toml"
)

// Config 定义配置信息
type Config struct {
	Environment *EnvironmentConfig `toml:"environment"` // 系统环境
	Server      *ServerConfig      `toml:"server"`      // 系统服务配置
	JWT         *JWTConfig         `toml:"jwt"`         // jwt 鉴权
	MySQL       *MySQLConfig       `toml:"mysql"`       // mysql 数据库配置
	Redis       *RedisConfig       `toml:"redis"`       // redis 数据库配置
	Sqlite      *SqliteConfig      `toml:"sqlite"`      // sqlite 数据库配置
	Logger      *LoggerConfig      `toml:"logger"`      // 日志配置
	Tasks       *TasksConfig       `toml:"tasks"`       // 定时任务
}

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

// JWTConfig jwt 鉴权
type JWTConfig struct {
	Secret string        `toml:"secret"` // 加密密匙
	Expire time.Duration `toml:"expire"` // 过期时间(h)
	Issuer string        `toml:"issuer"` // 签发人
	Prefix string        `toml:"prefix"` // 前缀
	Header string        `toml:"header"` // 请求标识
}

// GetExpire 获取过期时间(h)
func (r *JWTConfig) GetExpire() time.Duration {
	return r.Expire * time.Hour
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Read  MySQLAuthConfig `toml:"read"`
	Write MySQLAuthConfig `toml:"write"`
	Base  struct {
		MaxOpenConn     int           `toml:"max_open_conn"`     // 最大打开的连接数
		MaxIdleConn     int           `toml:"max_idle_conn"`     // 闲置的连接数
		ConnMaxLifeTime time.Duration `toml:"conn_max_lifetime"` // 设置最大连接超时(min)
	} `toml:"base"`
}

// MySQLAuthConfig Mysql 配置信息
type MySQLAuthConfig struct {
	Key      string `toml:"key"`      // db信息唯一标识
	Host     string `toml:"host"`     // db连接实例IP或域名
	Port     int    `toml:"port"`     // db连接实例端口
	DbName   string `toml:"db_name"`  // db库名
	Username string `toml:"username"` // db连接账号
	Password string `toml:"password"` // db连接密码
}

// RedisConfig 数据库配置
type RedisConfig struct {
	Host       string `toml:"host"`        // IP或域名
	Port       int    `toml:"port"`        // 端口
	Password   string `toml:"password"`    // 连接密码
	MaxRetries int    `toml:"max_retries"` // 最大重试次数
	PoolSize   int    `toml:"pool_size"`   // 连接池大小
}

// SqliteConfig sqlite3 数据库配置
type SqliteConfig struct {
	FilePath string `toml:"filepath"` // sqlite3 文件路径
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Filename   string `toml:"filename"`    // 日志文件路径
	Level      string `toml:"level"`       // 日志级别: debug/info/warn/error/panic
	MaxSize    int    `toml:"max_size"`    // 日志文件旋转之前的最大大小
	MaxBackups int    `toml:"max_backups"` // 保留的旧日志文件的最大数量
	MaxAge     int    `toml:"max_age"`     // 保留旧日志文件的最大天数
	Color      bool   `toml:"color"`
}

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

// TasksConfig 定时任务
type TasksConfig struct {
	Ticker map[string]bool `toml:"ticker"` // 即时器
	Timer  map[string]bool `toml:"timer"`  // 定时器
}

// IsEnableTicker 是否启用即时器
func (t *TasksConfig) IsEnableTicker(taskName string) bool {
	falg, ok := t.Ticker[taskName]
	if !ok {
		return false
	}
	return falg
}

// IsEnableTicker 是否启用定时器
func (t *TasksConfig) IsEnableTimer(taskName string) bool {
	falg, ok := t.Timer[taskName]
	if !ok {
		return false
	}
	return falg
}

// Init 加载配置文件
func Init(filepath string) {
	once.Do(func() {
		// 读取配置文件
		buf, err := os.ReadFile(filepath)
		if err != nil {
			panic(fmt.Sprintf("配置文件读取失败! err: %v", err))
		}
		// 解析配置信息至配置结构体
		if err := toml.Unmarshal(buf, &config); err != nil {
			panic(fmt.Sprintf("配置文件解析失败! err: %v", err))
		}
	})
}

// Instance 获取配置实例
func Instance() Config {
	return *config
}
