/*内存缓存*/
package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	// 创建一个默认过期时间为 5 分钟的缓存，每 10 分钟清除一次过期项目
	cacheInstance = &cacheRepo{cache: cache.New(5*time.Minute, 10*time.Minute)}
)

// CacheRepo 缓存接口
type CacheRepo interface {
	Set(key string, value interface{})
	SetAt(key string, value interface{}, expire time.Duration)
	Get(key string) (interface{}, bool)
}

// 缓存
type cacheRepo struct {
	cache *cache.Cache
}

// Set 设置缓存值，无过期时间
func (c *cacheRepo) Set(key string, value interface{}) {
	c.cache.Set(key, value, cache.NoExpiration)
}

// SetAt 设置缓存值，设置过期时间
func (c *cacheRepo) SetAt(key string, value interface{}, expire time.Duration) {
	c.cache.Set(key, value, expire)
}

// Get 获取缓存值
func (c *cacheRepo) Get(key string) (interface{}, bool) {
	value, ok := c.cache.Get(key)
	if !ok {
		return nil, false
	}
	return value, true
}

// Instance 获取缓存实例
func Instance() CacheRepo {
	return cacheInstance
}
