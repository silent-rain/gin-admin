/*验证码
 */
package system

// Captcha 验证码
type Captcha struct {
	CaptchaId string `json:"captcha_id"` // 验证 Key
	B64s      string `json:"b64s"`       // 验证图片 base64 值
}
