/*验证码*/
package utils

import (
	"image/color"
	"time"

	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/mojocn/base64Captcha"
)

// 更改验证码存储上限，以下代码设置存储的验证码为 20240个，过期时间为 3分钟
var CaptchaStore = base64Captcha.NewMemoryStore(20240, 3*time.Minute)

// 验证码
type captcha struct {
}

// 新建验证码对象
func NewCaptcha() captcha {
	return captcha{}
}

// 生成图形化算术验证码配置
func (r captcha) mathConfig() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return mathType
}

// 生成图形化数字验证码配置
func (r captcha) digitConfig() *base64Captcha.DriverDigit {
	digitType := &base64Captcha.DriverDigit{
		Height:   40,
		Width:    120,
		Length:   5,
		MaxSkew:  0.45,
		DotCount: 80,
	}
	return digitType
}

// 生成图形化字符串验证码配置
func (r captcha) stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          100,
		Width:           50,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          5,
		Source:          "123456789qwertyuiopasdfghjklzxcvb",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return stringType
}

// 生成图形化汉字验证码配置
func (r captcha) chineseConfig() *base64Captcha.DriverChinese {
	chineseType := &base64Captcha.DriverChinese{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
		Length:          2,
		Source:          "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,不想要,的值",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return chineseType
}

// 生成图形化数字音频验证码配置
func (r captcha) audioConfig() *base64Captcha.DriverAudio {
	chineseType := &base64Captcha.DriverAudio{
		Length:   4,
		Language: "zh",
	}
	return chineseType
}

// MekeCaptcha 生成验证码
func (r captcha) MekeCaptcha(captchaType string) (string, string, error) {
	var driver base64Captcha.Driver
	switch captchaType {
	case "audio":
		driver = r.audioConfig()
	case "string":
		driver = r.stringConfig()
	case "math":
		driver = r.mathConfig()
	case "chinese":
		driver = r.chineseConfig()
	case "digit":
		driver = r.digitConfig()
	default:
		return "", "", errcode.CaptchaTypeError
	}

	// 创建验证码并传入创建的类型的配置，以及存储的对象
	captcha := base64Captcha.NewCaptcha(driver, CaptchaStore)
	captchaId, base64string, err := captcha.Generate()
	return captchaId, base64string, err
}
