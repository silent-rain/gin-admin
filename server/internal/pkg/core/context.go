// Package core 扩展 gin context
package core

import (
	"github.com/gin-gonic/gin"
)

var contextKey = "__context__"

// 扩展 Context
type etxContext struct {
	// 用户信息
	UserId   uint   // 用户 ID
	Nickname string // 用户昵称

	// 日志链路 ID
	TraceId string

	// 中间件处理
	DisableCheckLogin  bool // 禁用登录检查
	DisableRateLimiter bool // 禁用接口限流
}

// GetContext 获取扩展 Context
func GetContext(ctx *gin.Context) *etxContext {
	newC := &etxContext{
		DisableCheckLogin:  false,
		DisableRateLimiter: false,
	}
	if ctx == nil {
		return newC
	}
	c, ok := ctx.Get(contextKey)
	if !ok {
		ctx.Set(contextKey, newC)
		return newC
	}
	return c.(*etxContext)
}
