// Package cache 应用配置表
package cache

import (
	"github.com/silent-rain/gin-admin/internal/app/system/dao"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/cache"
	"github.com/silent-rain/gin-admin/pkg/errcode"
)

// Config 站点配置缓存接口
type WebSiteConfigCache interface {
	Set() error
	Get() ([]model.Config, error)
}

// 站点配置缓存
type webSiteConfigCache struct {
	cache cache.CacheRepo
	dao   dao.Config
}

// NewWebSiteConfigCache 创建站点配置缓存对象
func NewWebSiteConfigCache() *webSiteConfigCache {
	return &webSiteConfigCache{
		cache: cache.Instance(),
	}
}

// Set 设置站点配置缓存
func (c *webSiteConfigCache) Set() error {
	results, err := dao.NewConfigDao().ChildrenByKey(constant.WebsiteConfigKey)
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return nil
	}
	c.cache.Set(constant.CacheWebSiteConfig, results)
	return nil
}

// Get 获取站点配置缓存
func (c *webSiteConfigCache) Get() ([]model.Config, error) {
	value, ok := c.cache.Get(constant.CacheWebSiteConfig)
	if !ok {
		return nil, errcode.CacheGetError
	}

	vs, ok := value.([]model.Config)
	if !ok {
		return nil, errcode.DataTypeConversionError
	}
	return vs, nil
}
