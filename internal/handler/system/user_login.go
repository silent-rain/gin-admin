/*
 * @Author: silent-rain
 * @Date: 2023-01-08 16:47:40
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:00:40
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/user_login.go
 * @Descripttion: 用户登录/登出
 */
package system

import (
	"bytes"

	systemDao "gin-admin/internal/dao/system"
	systemDto "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户登录注册对象
var UserLoginImpl = new(userLoginHandler)

// 用户登录/登出
type userLoginHandler struct {
}

// 登录
func (h *userLoginHandler) Login(ctx *gin.Context) {
	req := new(systemDto.UserLoginReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		zap.S().Errorf("data: %v, err: %v", req, err)
		return
	}

	// 验证码验证

	// 查询用户
	user, ok, err := systemDao.UserImpl.GetUsername(req.Username, req.Password)
	if err != nil {
		zap.S().Errorf("code: %v, err: %v", statuscode.DbQueryError, err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	if !ok {
		zap.S().Infof("code: %v, err: %v", statuscode.DbQueryError, "用户名或者密码不正确")
		response.New(ctx).WithCode(statuscode.DbQueryEmptyError).WithMsg("用户名或者密码不正确").Json()
		return
	}
	// 判断当前用户状态
	if user.Status != 1 {
		zap.S().Infof("code: %v, err: %v", statuscode.UserDisableError, statuscode.UserDisableError.Error())
		response.New(ctx).WithCode(statuscode.UserDisableError).Json()
		return
	}

	// 生成 Token
	token, err := utils.GenerateToken(user.ID, user.Phone, user.Email, user.Password)
	if err != nil {
		zap.S().Errorf("code: %v, err: %v", statuscode.TokenGenerateError, err)
		response.New(ctx).WithCode(statuscode.TokenGenerateError).Json()
		return
	}

	// 返回 Token
	result := systemDto.UserLoginRsp{
		Token: token,
	}
	response.New(ctx).WithMsg("登录成功").WithData(result).Json()
}

/*
   userInfo: {
     headImgUrl: "logo_W2xEX_2x.jpg",
     phone: "13302254696",
     roleId: "[13]",
     name: "panda",
     id: 5,
   },
*/

// 注销系统
func (h *userLoginHandler) Logout(ctx *gin.Context) {
	result := systemDto.UserLoginRsp{}
	response.New(ctx).WithMsg("注销成功").WithData(result).Json()
}

// 验证码
func (h *userLoginHandler) Captcha(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")

	captchaId := captcha.NewLen(5)

	var content bytes.Buffer
	ext := ".png"
	switch ext {
	case ".png":
		ctx.Header("Content-Type", "image/png")
		captcha.WriteImage(&content, captchaId, captcha.StdWidth, captcha.StdHeight)
	case ".wav":
		ctx.Header("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, captchaId, "zh")
	default:
		response.New(ctx).WithCode(statuscode.CaptchaEtxNotFoundError).Json()
		return
	}

	download := false
	if download {
		ctx.Header("Content-Type", "application/octet-stream")
	}

	session := sessions.Default(ctx)
	session.Set("captcha", captchaId)
	_ = session.Save()

	ctx.Writer.Write(content.Bytes())
}

// 验证码验证
func (h *userLoginHandler) CaptchaVerify(ctx *gin.Context) {
	value := ctx.Param("value")
	session := sessions.Default(ctx)

	captchaId := session.Get("captcha")
	if captchaId == nil {
		zap.S().Errorf("code: %v, data: %v, err: %v", statuscode.SessionGetCaptchaEmptyError, captchaId, statuscode.SessionGetCaptchaEmptyError.Error())
		response.New(ctx).WithCode(statuscode.SessionGetCaptchaEmptyError).Json()
		return
	}
	session.Delete("captcha")
	_ = session.Save()
	if !captcha.VerifyString(captchaId.(string), value) {
		zap.S().Errorf("code: %v, err: %v", statuscode.CaptchaVerifyError, statuscode.CaptchaVerifyError.Error())
		response.New(ctx).WithCode(statuscode.CaptchaVerifyError).Json()
		return
	}
	response.New(ctx).WithMsg("验证成功").Json()
}
