// Package conf MySQL 配置
package conf

import (
	"time"

	"gorm.io/gorm/logger"
)

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Read    MySQLAuthConfig    `toml:"read"`
	Write   MySQLAuthConfig    `toml:"write"`
	Options MySQLOptionsConfig `toml:"options"`
}

// MySQLConfig Mysql 认证信息
type MySQLAuthConfig struct {
	Key      string `toml:"key"`      // db信息唯一标识
	Host     string `toml:"host"`     // db连接实例IP或域名
	Port     int    `toml:"port"`     // db连接实例端口
	DbName   string `toml:"db_name"`  // db库名
	Username string `toml:"username"` // db连接账号
	Password string `toml:"password"` // db连接密码
}

// MySQLOptionsConfig Mysql 配置参数
type MySQLOptionsConfig struct {
	MaxOpenConn     int           `toml:"max_open_conn"`     // 最大打开的连接数
	MaxIdleConn     int           `toml:"max_idle_conn"`     // 闲置的连接数
	ConnMaxLifeTime time.Duration `toml:"conn_max_lifetime"` // 设置最大连接超时(min)
	LogLevel        string        `toml:"log_level"`         // 日志级别: info/warn/error/silent
}

// GetLogLevel 获取数据库日志级别
func (m MySQLOptionsConfig) GetLogLevel() logger.LogLevel {
	dict := map[string]logger.LogLevel{
		"info":   logger.Info,
		"warn":   logger.Info,
		"error":  logger.Info,
		"silent": logger.Info,
	}
	level, ok := dict[m.LogLevel]
	if !ok {
		return logger.Warn
	}
	return level
}

// SqliteConfig sqlite3 数据库配置
type SqliteConfig struct {
	FilePath string `toml:"filepath"` // sqlite3 文件路径
}
