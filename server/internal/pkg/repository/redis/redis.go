/*Redis 数据库*/
package redis

import (
	"fmt"
	"time"

	"gin-admin/internal/pkg/conf"
	"gin-admin/pkg/errcode"

	"gopkg.in/redis.v5"
)

// DBRepo 数据库接口
type DBRepo interface {
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
	TTL(key string) (time.Duration, error)
	Expire(key string, ttl time.Duration) bool
	ExpireAt(key string, ttl time.Time) bool
	Exists(keys ...string) bool
	Del(key string) bool
	Incr(key string) int64
	Decr(key string) int64
	Close() error
}

var (
	dbInstance DBRepo
)

// New 创建 Redis 客户端
func New() (DBRepo, error) {
	cfg := conf.Instance().Redis

	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:   cfg.Password,
		DB:         cfg.Db,
		MaxRetries: cfg.MaxRetries,
		PoolSize:   cfg.PoolSize,
	})
	// 测试链接
	if err := client.Ping().Err(); err != nil {
		return nil, errcode.New(errcode.RedisPingError)
	}
	return &dbRepo{client: client}, nil
}

// 数据库
type dbRepo struct {
	client *redis.Client
}

// Set set some <key,value,ttl> into redis
func (d *dbRepo) Set(key, value string, ttl time.Duration) error {
	if err := d.client.Set(key, value, ttl).Err(); err != nil {
		return errcode.New(errcode.RedisSetKeyError).WithMsg(err.Error())
	}
	return nil
}

// Get 获取 KEY 的值
func (d *dbRepo) Get(key string) (string, error) {
	value, err := d.client.Get(key).Result()
	if err != nil {
		return "", errcode.New(errcode.RedisGetKeyError).WithMsg(err.Error())
	}
	return value, nil
}

// TTL 查看 Key 剩余的过期时间，以秒为单位。
func (d *dbRepo) TTL(key string) (time.Duration, error) {
	ttl, err := d.client.TTL(key).Result()
	if err != nil {
		return -1, errcode.New(errcode.RedisTTLGetKeyError).WithMsg(err.Error())
	}
	return ttl, nil
}

// Expire 设置 key 的过期时间，以秒为单位
func (d *dbRepo) Expire(key string, ttl time.Duration) bool {
	ok, _ := d.client.Expire(key, ttl).Result()
	return ok
}

// ExpireAt 用于为 key 设置过期时间，不同在于，它的时间参数值采用的是时间戳格式。
func (d *dbRepo) ExpireAt(key string, ttl time.Time) bool {
	ok, _ := d.client.ExpireAt(key, ttl).Result()
	return ok
}

// Exists 用于检查指定的一个 key 或者多个 key 是否存在。
// 若存在则返回 1，否则返回 0
func (d *dbRepo) Exists(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := d.client.ExistsMulti(keys...).Result()
	return value > 0
}

// Del 若键存在的情况下，该命令用于删除键
func (d *dbRepo) Del(key string) bool {
	if key == "" {
		return true
	}
	value, _ := d.client.Del(key).Result()
	return value > 0
}

// Incr 将 key 中储存的数字值增一
func (d *dbRepo) Incr(key string) int64 {
	value, _ := d.client.Incr(key).Result()
	return value
}

// Decr 将 key 中储存的数字值减一
func (d *dbRepo) Decr(key string) int64 {
	value, _ := d.client.Decr(key).Result()
	return value
}

// Close 关闭客户端
func (d *dbRepo) Close() error {
	return d.client.Close()
}

// Init 初始化数据库
func Init() error {
	db, err := New()
	if err != nil {
		panic(fmt.Sprintf("初始化数据库失败! err: %v", err))
	}
	dbInstance = db
	return err
}

// Instance 获取数据库实例
func Instance() DBRepo {
	return dbInstance
}
