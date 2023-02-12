/*配置
 */
package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var (
	once   sync.Once
	config *Config
	// ConfigFile 配置文件路径
	ConfigFile = "./conf.yaml"
)

const (
	// Secret 加密密匙
	Secret = "8Xui8SN4mI+7egV/9dlfYYLGQJeEx4+DwmSQLwDVXJg="

	// Token 过期时间
	TokenExpireDuration = time.Hour * 24
	// Token 签发人
	TokenIssuer = "silent-rain"
	// Token 前缀
	TokenPrefix = "Bearer "
	TokenHeader = "Authorization"

	// Session 最大过期时间
	SessionMaxAge = time.Hour * 24
	// Session 密匙对
	SessionKeyPairs = "silent-rain"

	// ServerUserDefaultPwd 用户默认密码
	ServerUserDefaultPwd = "888888"

	// 验证码类型
	CaptchaType = "digit"
)

// Config 定义配置信息
type Config struct {
	ServerConfig *ServerConfig `yaml:"server"` // 系统服务配置
	DBConfig     *DBConfig     `yaml:"db"`     // mysql 数据库配置
	SqliteConfig *SqliteConfig `yaml:"sqlite"` // sqlite 数据库配置
	UploadConfig *UploadConfig `yaml:"upload"` // 上传文件配置
	LoggerConfig *LoggerConfig `yaml:"logger"` // 日志配置
	EnvConfig    *EnvConfig    `yaml:"env"`    // 系统环境
}

// ServerConfig 系统服务配置
type ServerConfig struct {
	Address string
	Port    int
}

// ServerAddress 获取服务地址
func (s ServerConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%d", s.Address, s.Port)
}

// DBConfig mysql 数据库配置;
// 当 mysql 配置为空时，使用 sqlite3 数据库
type DBConfig struct {
	Key      string `yaml:"key"`      // db信息唯一标识
	Host     string `yaml:"host"`     // db连接实例IP或域名
	Port     int    `yaml:"port"`     // db连接实例端口
	DbName   string `yaml:"db_name"`  // db库名
	Username string `yaml:"username"` // db连接账号
	Password string `yaml:"password"` // db连接密码
}

// Dsn 拼接 mysql 数据库 DSN 地址
func (r DBConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		r.Username,
		r.Password,
		r.Host,
		r.Port,
		r.DbName,
	)
}

// SqliteConfig sqlite3 数据库配置
type SqliteConfig struct {
	FilePath string `yaml:"filepath"` // sqlite3 文件路径
}

// UploadConfig 上传文件配置
type UploadConfig struct {
	FilePath string `yaml:"filepath"` // 文件文件路径
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Filename   string `yaml:"filename"`    // 日志文件路径
	Level      string `yaml:"level"`       // 日志级别: debug/info/warn/error/panic
	MaxSize    int    `yaml:"max_size"`    // 日志文件旋转之前的最大大小
	MaxBackups int    `yaml:"max_backups"` // 保留的旧日志文件的最大数量
	MaxAge     int    `yaml:"max_age"`     // 保留旧日志文件的最大天数
	Color      bool   `yaml:"color"`
}

// EnvConfig 系统环境配置 prod/test/dev
type EnvConfig string

// Env 获取环境名称
func (r EnvConfig) Env() string {
	var mode = gin.DebugMode
	switch r {
	case "prod":
		mode = gin.ReleaseMode
	case "test":
		mode = gin.TestMode
	case "debug":
		mode = gin.DebugMode
	}
	return mode
}

// String 将日志配置结构体转换为字符串
func (r *LoggerConfig) String() string {
	buf, err := json.Marshal(r)
	if err != nil {
		panic("日志配置信息编码失败! err: %v" + err.Error())
	}
	return string(buf)
}

// InitLoadConfig 加载配置文件
func InitLoadConfig(filepath string) {
	once.Do(func() {
		// 环境变量获取配置文件
		envConfigPath := os.Getenv("config")
		if envConfigPath != "" {
			filepath = envConfigPath
		}
		// 读取配置文件
		buf, err := os.ReadFile(filepath)
		if err != nil {
			panic(fmt.Sprintf("配置文件读取失败! err: %v", err))
		}
		// 解析配置信息至配置结构体
		if err := yaml.Unmarshal(buf, &config); err != nil {
			panic(fmt.Sprintf("配置文件解析失败! err: %v", err))
		}
	})
}

// Instance 获取配置实例
func Instance() *Config {
	if config == nil {
		InitLoadConfig(ConfigFile)
	}
	return config
}
