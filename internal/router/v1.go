/*
 * @Author: silent-rain
 * @Date: 2023-01-06 00:26:00
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 17:38:49
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
	v1.GET("/sayHello", handler.SayHello)

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
		user.GET("/info", system.UserHandlerImpl.Info)
		// 获取所有用户列表
		user.GET("/all", system.UserHandlerImpl.All)
		// 获取用户列表
		user.GET("/list", system.UserHandlerImpl.List)
		// 添加用户
		user.POST("/add", system.UserRegisterHandlerImpl.Add)
		// 更新用户详情信息
		user.PUT("/updateDetails", system.UserHandlerImpl.UpdateDetails)
		// 删除用户
		user.DELETE("/delete", system.UserHandlerImpl.Delete)
		// 更新用户状态
		user.PUT("/status", system.UserHandlerImpl.Status)
		// 更新用户密码
		user.PUT("/updatePwd", system.UserHandlerImpl.UpdatePassword)
		// 重置用户密码
		user.PUT("/resetPwd", system.UserHandlerImpl.ResetPassword)
		// 更新用户手机号码
		user.PUT("/updatePhone", system.UserHandlerImpl.UpdatePhone)
		// 更新用户邮箱
		user.PUT("/updateEmail", system.UserHandlerImpl.UpdateEmail)
	}

	// 角色管理
	role := v1.Group("/role")
	{
		// 获取所有角色列表
		role.GET("/all", system.RoleHandlerImpl.All)
		// 获取角色列表
		role.GET("/list", system.RoleHandlerImpl.List)
		// 添加角色
		role.POST("/add", system.RoleHandlerImpl.Add)
		// 更新角色
		role.PUT("/update", system.RoleHandlerImpl.Update)
		// 删除角色
		role.DELETE("/delete", system.RoleHandlerImpl.Delete)
		// 更新角色状态
		role.PUT("/status", system.RoleHandlerImpl.Status)
	}
}
