/*内嵌外部资源
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

//go:embed dist/static dist/static/js/_*.js
var WebStaticAssets embed.FS

//go:embed docs
var WebDocsAssets embed.FS
