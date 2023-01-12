/*
 * @Author: silent-rain
 * @Date: 2023-01-05 23:33:10
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-12 22:53:39
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
