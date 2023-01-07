/**配置
 */
package conf

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var (
	once       sync.Once
	config     *Config
	ConfigFile = "./conf.yaml"
)

// Config 定义配置信息
type Config struct {
	*DBConfig     `yaml:"db"`     // 数据库配置
	*LoggerConfig `yaml:"logger"` // 日志配置
	*EnvConfig    `yaml:"env"`    // 系统环境
}

// DBConfig 数据库配置
type DBConfig struct {
	Key      string `yaml:"key"`      // db信息唯一标识
	Host     string `yaml:"host"`     // db连接实例IP或域名
	Port     int    `yaml:"port"`     // db连接实例端口
	DbName   string `yaml:"db_name"`  // db库名
	Username string `yaml:"username"` // db连接账号
	Password string `yaml:"password"` // db连接密码
	Url      string `yaml:"url"`      // url 链接
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	FileName string `yaml:"filename"`
	Level    int    `yaml:"level"`
	MaxLines int    `yaml:"max_lines"`
	MaxSize  int    `yaml:"max_size"`
	MaxDays  int    `yaml:"max_days"`
	Daily    bool   `yaml:"daily"`
	Color    bool   `yaml:"color"`
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

// 将日志配置结构体转换为字符串
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
		// 读取配置文件
		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			panic("配置文件读取失败! err: %v" + err.Error())
		}
		// 解析配置信息至配置结构体
		if err := yaml.Unmarshal(buf, &config); err != nil {
			panic("配置文件解析失败! err: %v" + err.Error())
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
