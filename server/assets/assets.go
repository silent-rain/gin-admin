/*Package assets 内嵌外部资源
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

// 正则:go:embed dist/static dist/static/js/_*.js
//
//go:embed dist/static
var WebStaticAssets embed.FS

//go:embed docs
var WebDocsAssets embed.FS
