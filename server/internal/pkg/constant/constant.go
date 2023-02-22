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

	// Token 过期时间
	TokenExpireDuration = time.Hour * 24
	// Token 签发人
	TokenIssuer = "silent-rain"
	// Token 前缀
	TokenPrefix = "Bearer "
	TokenHeader = "Authorization"

	// Session 最大过期时间
	SessionMaxAge = time.Hour * 24
	// Session 密匙对
	SessionKeyPairs = "silent-rain"

	// ServerUserDefaultPwd 用户默认密码
	ServerUserDefaultPwd = "888888"

	// 验证码类型
	CaptchaType = "digit"
)
