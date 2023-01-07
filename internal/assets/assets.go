/**内嵌外部资源
 */
package assets

import (
	"embed"
)

// Assets 静态文件
// 用法: static/** favicon.ico index.html
//go:embed .gitignore
var Assets embed.FS

//go:embed dist
var WebAssets embed.FS

//go:embed dist/static
var WebStaticAssets embed.FS
