// Package conf Redis 配置
package conf

// RedisConfig 数据库配置
type RedisConfig struct {
	Host       string `toml:"host"`        // IP或域名
	Port       int    `toml:"port"`        // 端口
	Password   string `toml:"password"`    // 连接密码
	MaxRetries int    `toml:"max_retries"` // 最大重试次数
	PoolSize   int    `toml:"pool_size"`   // 连接池大小
	StoreType  string `toml:"store_type"`  // 数据库类型, redis/mem_sqlite
}
