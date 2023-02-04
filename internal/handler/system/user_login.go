/*
 * @Author: silent-rain
 * @Date: 2023-01-08 16:47:40
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-12 00:43:18
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
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginImpl 用户登录注册对象
var UserLoginImpl = new(userLoginHandler)

// 用户登录/登出
type userLoginHandler struct {
}

// Login 登录
func (h *userLoginHandler) Login(ctx *gin.Context) {
	req := new(systemDto.UserLoginReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 查询用户
	user, ok, err := systemDao.NewDaoUser().GetUsername(req.Username, req.Password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.DbQueryEmptyError).Error("用户名或者密码不正确")
		response.New(ctx).WithCode(statuscode.DbQueryEmptyError).WithMsg("用户名或者密码不正确").Json()
		return
	}
	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(statuscode.UserDisableError).Error("")
		response.New(ctx).WithCode(statuscode.UserDisableError).Json()
		return
	}

	// 生成 Token
	token, err := utils.GenerateToken(user.ID, user.Phone, user.Email, user.Password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.TokenGenerateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.TokenGenerateError).Json()
		return
	}

	// 返回 Token
	result := systemDto.UserLoginRsp{
		Token: token,
	}
	response.New(ctx).WithMsg("登录成功").WithData(result).Json()
}

// Logout 注销系统
func (h *userLoginHandler) Logout(ctx *gin.Context) {
	result := systemDto.UserLoginRsp{}
	response.New(ctx).WithMsg("注销成功").WithData(result).Json()
}

// Captcha 验证码
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

// CaptchaVerify 验证码验证
func (h *userLoginHandler) CaptchaVerify(ctx *gin.Context) {
	value := ctx.DefaultQuery("captchaId", "")
	if value == "" {
		log.New(ctx).WithCode(statuscode.SessionGetCaptchaEmptyError).Error("")
		response.New(ctx).WithCode(statuscode.SessionGetCaptchaEmptyError).Json()
		return
	}

	session := sessions.Default(ctx)
	captchaId := session.Get("captcha")
	if captchaId == nil {
		log.New(ctx).WithCode(statuscode.CaptchaNotFoundError).Error("")
		response.New(ctx).WithCode(statuscode.CaptchaNotFoundError).Json()
		return
	}
	session.Delete("captcha")
	_ = session.Save()
	if !captcha.VerifyString(captchaId.(string), value) {
		log.New(ctx).WithCode(statuscode.CaptchaVerifyError).Error("")
		response.New(ctx).WithCode(statuscode.CaptchaVerifyError).Json()
		return
	}
	response.New(ctx).WithMsg("验证成功").Json()
}
