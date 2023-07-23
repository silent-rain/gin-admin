// Package redis Redis 数据库
package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/silent-rain/gin-admin/pkg/conf"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/go-redis/redis/v8"
)

// DBRepo 数据库接口
type DBRepo interface {
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	Expire(ctx context.Context, key string, ttl time.Duration) (bool, error)
	ExpireAt(ctx context.Context, key string, ttl time.Time) (bool, error)
	Exists(ctx context.Context, keys ...string) (bool, error)
	Del(ctx context.Context, key string) (bool, error)
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	Close() error
}

// DBName Rides 数据库名称
type DBName = int

const (
	// Default 默认表
	Default DBName = iota
	// UserLogin 用户登录表
	UserLogin
	// ApiTokenLogin API Token 登录信息表
	ApiTokenLogin
)

// New 创建 Redis 客户端
func New(cfg *conf.RedisConfig, dbName DBName) (*Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	db, err := dbConnect(*cfg, dbName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(ctx).Err(); err != nil {
		return nil, errcode.RedisPingError
	}
	pool := &Pool{
		db,
	}
	return pool, nil
}

// 连接数据库
func dbConnect(cfg conf.RedisConfig, dbName DBName) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:   cfg.Password,
		DB:         dbName,
		MaxRetries: cfg.MaxRetries,
		PoolSize:   cfg.PoolSize,
	})

	// 测试链接
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errcode.RedisPingError
	}
	return client, nil
}

// 数据库连接池
type Pool struct {
	client *redis.Client
}

// Set set some <key,value,ttl> into redis
func (d *Pool) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	if err := d.client.Set(ctx, key, value, ttl).Err(); err != nil {
		return errcode.RedisSetKeyError
	}
	return nil
}

// Get 获取 KEY 的值
func (d *Pool) Get(ctx context.Context, key string) (string, error) {
	value, err := d.client.Get(ctx, key).Result()
	if err != nil {
		return "", errcode.RedisGetKeyError
	}
	return value, nil
}

// TTL 查看 Key 剩余的过期时间，以秒为单位。
func (d *Pool) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := d.client.TTL(ctx, key).Result()
	if err != nil {
		return -1, errcode.RedisTTLGetKeyError
	}
	return ttl, nil
}

// Expire 设置 key 的过期时间，以秒为单位
func (d *Pool) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	ok, err := d.client.Expire(ctx, key, ttl).Result()
	if err != nil {
		return false, errcode.RedisSetKeyExpireError
	}
	return ok, nil
}

// ExpireAt 用于为 key 设置过期时间，不同在于，它的时间参数值采用的是时间戳格式。
func (d *Pool) ExpireAt(ctx context.Context, key string, ttl time.Time) (bool, error) {
	ok, err := d.client.ExpireAt(ctx, key, ttl).Result()
	if err != nil {
		return false, errcode.RedisSetKeyExpireAtError
	}
	return ok, nil
}

// Exists 用于检查指定的一个 key 或者多个 key 是否存在。
// 若存在则返回 1，否则返回 0
func (d *Pool) Exists(ctx context.Context, keys ...string) (bool, error) {
	if len(keys) == 0 {
		return true, nil
	}
	value, err := d.client.Exists(ctx, keys...).Result()
	if err != nil {
		return false, errcode.RedisGetKeyExistsError
	}
	return value > 0, nil
}

// Del 若键存在的情况下，该命令用于删除键
func (d *Pool) Del(ctx context.Context, key string) (bool, error) {
	if key == "" {
		return true, nil
	}
	value, err := d.client.Del(ctx, key).Result()
	if err != nil {
		return false, errcode.RedisDelKeyError
	}
	return value > 0, nil
}

// Incr 将 key 中储存的数字值增一
func (d *Pool) Incr(ctx context.Context, key string) (int64, error) {
	value, err := d.client.Incr(ctx, key).Result()
	if err != nil {
		return -1, errcode.RedisIncrKeyError
	}
	return value, nil
}

// Decr 将 key 中储存的数字值减一
func (d *Pool) Decr(ctx context.Context, key string) (int64, error) {
	value, err := d.client.Decr(ctx, key).Result()
	if err != nil {
		return -1, errcode.RedisDecrKeyError
	}
	return value, nil
}

// Close 关闭客户端
func (d *Pool) Close() error {
	if err := d.client.Close(); err != nil {
		return err
	}
	return nil
}
