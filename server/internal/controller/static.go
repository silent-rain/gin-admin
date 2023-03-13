/*静态资源
 */
package controller

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/silent-rain/gin-admin/assets"
	systemDAO "github.com/silent-rain/gin-admin/internal/dao/system"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// FaviconIco 网站 favicon.ico
func FaviconIco(ctx *gin.Context) {
	ctx.Header("content-type", "image/x-icon")
	ctx.FileFromFS("dist/favicon.ico", http.FS(assets.WebAssets))
	ctx.Status(200)
}

// Index 网站首页
func Index(ctx *gin.Context) {
	ctx.Header("content-type", "text/html;charset=utf-8")
	buf, err := assets.WebAssets.ReadFile("dist/index.html")
	if err != nil {
		log.New(ctx).WithCode(errcode.LoadIndexHtmlError).Errorf("index.html 文件加载失败 %#v", err)
		// 正常加载首页
		ctx.HTML(http.StatusOK, "index.html", nil)
		return
	}
	text := string(buf)

	// SEO 优化
	reg := regexp.MustCompile("(<title>.*</title>)")
	find := reg.FindString(text)

	if find == "" {
		log.New(ctx).WithCode(errcode.LoadIndexHtmlError).Errorf("index.html 未找到 title 标签  %#v", err)
		// 正常加载首页
		ctx.HTML(http.StatusOK, "index.html", nil)
		return
	}

	// 加载 SEO 内容
	results, err := systemDAO.NewWebSiteConfigCache().Get()
	if err != nil {
		log.New(ctx).WithCode(errcode.LoadIndexHtmlError).Errorf("站点配置缓存查询失败  %#v", err)
		ctx.HTML(http.StatusOK, "index.html", nil)
		return
	}
	if len(results) == 0 {
		log.New(ctx).WithCode(errcode.LoadIndexHtmlError).Errorf("站点配置缓存为空")
		ctx.HTML(http.StatusOK, "index.html", nil)
		return
	}
	configMap := make(map[string]interface{}, 0)
	for _, item := range results {
		configMap[item.Key] = item.Value
	}

	// title/description/keywords/anthor/robots
	// title，description，keywords他们的权重逐渐减小。
	// title 就是我们看到的网页标题
	// description 为对该网页的简要描述
	// keywords 的作用就是告诉搜索引擎，本网页中主要围绕着哪些方面的关键词展开
	seoText := fmt.Sprintf(`
	<title>%s</title>
	<meta name="description" content='%v'>
	<meta name="keywords" content="%v">
	<meta name="anthor" content="%v">
	`,
		configMap["website_title"],
		configMap["website_seo_desc"],
		configMap["website_keywords"],
		configMap["website_anthor"],
	)
	text = reg.ReplaceAllString(text, seoText)
	ctx.String(http.StatusOK, text)
}
