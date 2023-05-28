// Package initialize 系统缓存初始化
package initialize

import (
	"github.com/silent-rain/gin-admin/internal/app/system/cache"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
)

// 初始化站点配置缓存
func initWebSiteConfigCache() {
	if err := cache.NewWebSiteConfigCache().Set(); err != nil {
		log.New(nil).WithError(err).Error("")
		return
	}
}
