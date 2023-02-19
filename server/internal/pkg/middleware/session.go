/*session 中间件
 */
package middleware

import (
	"gin-admin/internal/pkg/conf"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 中间件
func Session() gin.HandlerFunc {
	store := sessionConfig()
	return sessions.Sessions(conf.SessionKeyPairs, store)
}

// session 配置
func sessionConfig() sessions.Store {
	store := cookie.NewStore([]byte(conf.Secret))
	store.Options(sessions.Options{
		MaxAge: int(conf.SessionMaxAge.Milliseconds()), // seconds
		Path:   "/",
	})
	return store
}
