/*session 中间件
 */
package middleware

import (
	"github.com/silent-rain/gin-admin/internal/pkg/constant"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 中间件
func Session() gin.HandlerFunc {
	store := sessionConfig()
	return sessions.Sessions(constant.SessionKeyPairs, store)
}

// session 配置
func sessionConfig() sessions.Store {
	store := cookie.NewStore([]byte(constant.Secret))
	store.Options(sessions.Options{
		MaxAge: int(constant.SessionMaxAge.Milliseconds()), // seconds
		Path:   "/",
	})
	return store
}
