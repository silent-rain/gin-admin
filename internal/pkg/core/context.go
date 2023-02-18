/*扩展 gin context*/
package core

import "github.com/gin-gonic/gin"

var contextKey = "__context__"

// 扩展 Context
type etxContext struct {
	// 用户信息
	UserId uint // 用户 ID

	// 日志
	TraceId string
	SpanId  string
}

// GetContext 获取扩展 Context
func GetContext(ctx *gin.Context) *etxContext {
	c, ok := ctx.Get(contextKey)
	if ok {
		return c.(*etxContext)
	}
	newC := &etxContext{}
	ctx.Set(contextKey, newC)
	return newC
}
