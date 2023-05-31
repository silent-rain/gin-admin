// Package initialize 数据初始化
package initialize

// Init 初始化数据
func Init() {
	// 初始化站点配置缓存
	initWebSiteConfigCache()
	// 数据库表自动迁移
	initTableAutoMigrate()
}
