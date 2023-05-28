// Package cache 用户登录信息表
package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/silent-rain/gin-admin/internal/pkg/conf"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/redis"
)

// UserLoginCache 用户登录信息缓存接口
type UserLoginCache interface {
	Set(userId uint, token string) error
	Get(userId uint) (string, error)
}

// 用户登录信息缓存
type redisUserLogin struct {
	db redis.DBRepo
}

// NewUserLoginCache 创建用户登录信息对象
func NewUserLoginCache() *redisUserLogin {
	return &redisUserLogin{
		db: redis.Instance().DB(redis.UserLogin),
	}
}

// Set 设置缓存
func (d *redisUserLogin) Set(userId uint, token string) error {
	expire := conf.Instance().JWT.GetExpire()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	return d.db.Set(ctx, strconv.Itoa(int(userId)), token, expire)
}

// Get 获取缓存
func (d *redisUserLogin) Get(userId uint) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return d.db.Get(ctx, strconv.Itoa(int(userId)))
}
