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
		userLogin.POST("/register", system.NewRegisterUserHandler().Add)
		// 验证码
		userLogin.GET("/captcha", system.NewUserLoginHandler().Captcha)
		// 验证码验证
		userLogin.GET("/captcha/verify", system.NewUserLoginHandler().CaptchaVerify)
		// 登录
		userLogin.POST("/login", system.NewUserLoginHandler().Login)
		// 登出
		userLogin.POST("/logout", system.NewUserLoginHandler().Logout)
	}

	// 用户管理
	user := v1.Group("/user")
	{
		// 获取用户信息
		user.GET("/info", system.NewUserHandler().Info)
		// 获取所有用户列表
		user.GET("/all", system.NewUserHandler().All)
		// 获取用户列表
		user.GET("/list", system.NewUserHandler().List)
		// 添加用户
		user.POST("/add", system.NewRegisterUserHandler().Add)
		// 更新用户详情信息
		user.PUT("/update", system.NewUserHandler().Update)
		// 删除用户
		user.DELETE("/delete", system.NewUserHandler().Delete)
		// 批量删除用户
		user.DELETE("/batchDelete", system.NewUserHandler().BatchDelete)
		// 更新用户状态
		user.PUT("/status", system.NewUserHandler().Status)
		// 重置用户密码
		user.PUT("/resetPwd", system.NewUserHandler().ResetPassword)
		// 更新用户密码
		user.PUT("/updatePwd", system.NewUserHandler().UpdatePassword)
		// 更新用户手机号码
		user.PUT("/updatePhone", system.NewUserHandler().UpdatePhone)
		// 更新用户邮箱
		user.PUT("/updateEmail", system.NewUserHandler().UpdateEmail)
	}

	// 角色管理
	role := v1.Group("/role")
	{
		// 获取所有角色列表
		role.GET("/all", system.NewRoleHandler().All)
		// 获取角色列表
		role.GET("/list", system.NewRoleHandler().List)
		// 添加角色
		role.POST("/add", system.NewRoleHandler().Add)
		// 更新角色
		role.PUT("/update", system.NewRoleHandler().Update)
		// 删除角色
		role.DELETE("/delete", system.NewRoleHandler().Delete)
		// 批量删除角色
		role.DELETE("/batchDelete", system.NewRoleHandler().BatchDelete)
		// 更新角色状态
		role.PUT("/status", system.NewRoleHandler().Status)
	}

	// 菜单管理
	menu := v1.Group("/menu")
	{
		// 获取所有菜单树
		menu.GET("/allTree", system.NewMenuHandler().AllTree)
		// 获取菜单树
		menu.GET("/tree", system.NewMenuHandler().Tree)
		// 添加菜单
		menu.POST("/add", system.NewMenuHandler().Add)
		// 更新菜单
		menu.PUT("/update", system.NewMenuHandler().Update)
		// 删除菜单
		menu.DELETE("/delete", system.NewMenuHandler().Delete)
		// 批量删除菜单
		menu.DELETE("/batchDelete", system.NewMenuHandler().BatchDelete)
		// 更新菜单状态
		menu.PUT("/status", system.NewMenuHandler().Status)

	}

	// 角色菜单关系管理
	roleMenuRel := v1.Group("/roleMenuRel")
	{
		// 获取角色关联的菜单列表
		roleMenuRel.GET("/list", system.NewMenuRelHandler().List)
		// 更新角色菜单关联关系
		roleMenuRel.PUT("/update", system.NewMenuRelHandler().Update)
	}
}
