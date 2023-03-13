/*缓存初始化*/
package initialize

import (
	systemDAO "github.com/silent-rain/gin-admin/internal/dao/system"
	"github.com/silent-rain/gin-admin/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// 初始化站点配置缓存
func initWebSiteConfigCache() {
	if err := systemDAO.NewWebSiteConfigCache().Set(); err != nil {
		log.New(&gin.Context{}).WithError(err).Error("")
		return
	}
}
