// Package conf JWT 配置
package conf

import "time"

// JWTConfig jwt 鉴权
type JWTConfig struct {
	Secret string        `toml:"secret"` // 加密密匙
	Expire time.Duration `toml:"expire"` // 过期时间(h)
	Issuer string        `toml:"issuer"` // 签发人
	Prefix string        `toml:"prefix"` // 前缀
	Header string        `toml:"header"` // 请求标识
}

// GetExpire 获取过期时间(h)
func (r *JWTConfig) GetExpire() time.Duration {
	return r.Expire * time.Hour
}
