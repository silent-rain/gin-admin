// Package cache Http协议接口管理
package cache

import (
	"context"
	"time"

	"github.com/silent-rain/gin-admin/internal/app/api_auth/dto"
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/redis"
)

// ApiTokenLoginCache API Token 登录信息缓存接口
type ApiTokenLoginCache interface {
	Set(userId uint, token string) error
	Get(userId uint) (string, error)
}

// API Token 登录信息缓存
type redisApiTokenLogin struct {
	db redis.DBRepo
}

// NewApiTokenLoginCache 创建 API Token 登录信息缓存对象
func NewApiTokenLoginCache() *redisApiTokenLogin {
	return &redisApiTokenLogin{
		db: global.Instance().Redis(redis.ApiTokenLogin),
	}
}

// Set 设置缓存
func (d *redisApiTokenLogin) Set(tokenUri string, userId uint, Nickname string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	user := dto.ApiHttpUserCache{
		UserId:   userId,
		Nickname: Nickname,
	}
	value, err := user.String()
	if err != nil {
		return err
	}
	return d.db.Set(ctx, tokenUri, value, constant.ApiHttpTokenExpire)
}

// Get 获取缓存
func (d *redisApiTokenLogin) Get(tokenUri string) (dto.ApiHttpUserCache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	user := dto.ApiHttpUserCache{}
	value, err := d.db.Get(ctx, tokenUri)
	if err != nil {
		return dto.ApiHttpUserCache{}, err
	}
	if err = user.Unmarshal(value); err != nil {
		return user, err
	}
	return user, nil
}
