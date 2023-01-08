/*
 * @Author: silent-rain
 * @Date: 2023-01-06 00:26:00
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:04:45
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/router/v1.go
 * @Descripttion:
 */
/**Api version 1 路由
 */
package router

import (
	"gin-admin/internal/handler"
	"gin-admin/internal/handler/system"

	"github.com/gin-gonic/gin"
)

func NewApiV1(engine *gin.Engine) {
	v1 := engine.Group("api/v1")
	// 接口测试
	v1.GET("/sayHello/:name", handler.SayHello)

	// 注册/登录/登出
	userLogin := v1.Group("/")
	{
		// 注册
		userLogin.POST("/register", system.UserRegisterHandlerImpl.Add)
		// 验证码
		userLogin.GET("/captcha", system.UserLoginImpl.Captcha)
		// 验证码验证
		userLogin.GET("/captcha/verify/:value", system.UserLoginImpl.CaptchaVerify)
		// 登录
		userLogin.POST("/login", system.UserLoginImpl.Login)
		// 登出
		userLogin.POST("/logout", system.UserLoginImpl.Logout)
	}
}
