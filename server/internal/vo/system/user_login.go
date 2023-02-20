/*用户登录/登出
 */
package system

// UserLogin 登录响应
type UserLogin struct {
	Token string `json:"token"` // 令牌
}
