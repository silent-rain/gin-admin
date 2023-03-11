/*常量*/
package constant

import "time"

const (
	// 项目名称
	ProjectName = "gin-admin"

	// HeaderTraceTd 请求头 traceTd
	HeaderTraceTd = "trace_id"

	// Secret 加密密匙
	Secret = "8Xui8SN4mI+7egV/9dlfYYLGQJeEx4+DwmSQLwDVXJg="

	// Session 最大过期时间
	SessionMaxAge = time.Hour * 24
	// Session 密匙对
	SessionKeyPairs = "silent-rain"

	// ServerUserDefaultPwd 用户默认密码
	ServerUserDefaultPwd = "888888"

	// 验证码类型
	CaptchaType = "digit"

	// 站点配置 KEY
	WebsiteConfigKey = "website"

	// CacheWebSiteConfig 缓存 Key
	CacheWebSiteConfig = "webSiteConfig"

	// API 访问鉴权
	// ApiHttpToken API 请求密匙
	ApiHttpToken = "X-API-Token"
	// ApiHttpTokenPassphrase API 请求密匙口令
	ApiHttpTokenPassphrase = "X-Passphrase"
)
