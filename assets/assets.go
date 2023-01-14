/*
 * @Author: silent-rain
 * @Date: 2023-01-05 23:33:10
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 18:42:18
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/assets/assets.go
 * @Descripttion:
 */
/**内嵌外部资源
 */
package assets

import (
	"embed"
)

// Assets 静态文件
// 用法: static/** favicon.ico index.html
//
//go:embed .gitignore
var Assets embed.FS

//go:embed dist
var WebAssets embed.FS

//go:embed dist/static dist/static/js/_plugin-vue_export-helper-*.js
var WebStaticAssets embed.FS

//go:embed docs
var WebDocsAssets embed.FS
