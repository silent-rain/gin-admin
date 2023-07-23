// Package conf 配置
package conf

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	// ConfigFile 配置文件路径
	ConfigFile = "./config.toml"
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
	Schedule    *ScheduleConfig    `toml:"schedule"`    // 任务调度配置
}

// New 创建配置对象
func New(filepath string) *Config {
	// 读取配置文件
	buf, err := os.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintf("配置文件读取失败! err: %v", err))
	}
	config := &Config{}
	// 解析配置信息至配置结构体
	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(fmt.Sprintf("配置文件解析失败! err: %v", err))
	}
	return config
}
