/*用户登录/登出
 */
package system

type Captcha struct {
	CaptchaId string `json:"captcha_id"` // 验证 Key
	B64s      string `json:"b64s"`       // 验证图片 base64 值
}
