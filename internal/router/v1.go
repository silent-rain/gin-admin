/*
 * @Author: silent-rain
 * @Date: 2023-01-06 00:26:00
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 01:18:14
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

// NewApiV1 API V1 路由
func NewApiV1(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	// 接口测试
	v1.GET("/sayHello/:name", handler.SayHello)

	// 注册/登录/登出/验证码
	userLogin := v1.Group("/")
	{
		// 注册
		userLogin.POST("/register", system.UserRegisterHandlerImpl.Add)
		// 验证码
		userLogin.GET("/captcha", system.UserLoginImpl.Captcha)
		// 验证码验证
		userLogin.GET("/captcha/verify", system.UserLoginImpl.CaptchaVerify)
		// 登录
		userLogin.POST("/login", system.UserLoginImpl.Login)
		// 登出
		userLogin.POST("/logout", system.UserLoginImpl.Logout)
	}

	// 用户管理
	user := v1.Group("/user")
	{
		// 获取用户信息
		user.GET("/userInfo", system.UserManageImpl.UserInfo)
		// 添加用户
		user.POST("/add", system.UserRegisterHandlerImpl.Add)
	}

	// 角色管理
	role := v1.Group("/role")
	{
		// 获取列表
		role.GET("/list", system.RoleHandlerImpl.List)
		// 添加
		role.POST("/add", system.RoleHandlerImpl.Add)
		// 更新
		role.PUT("/update", system.RoleHandlerImpl.Update)
		// 删除
		role.DELETE("/delete", system.RoleHandlerImpl.Delete)
		// 更新状态
		role.POST("/status", system.RoleHandlerImpl.Status)
	}
}
