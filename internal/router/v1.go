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
	"gin-admin/internal/controller"
	"gin-admin/internal/controller/system"

	"github.com/gin-gonic/gin"
)

// NewApiV1 API V1 路由
func NewApiV1(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	// 接口测试
	v1.GET("/sayHello", controller.SayHello)

	// 注册/登录/登出/验证码
	userLogin := v1.Group("/")
	{
		// 注册
		userLogin.POST("/register", system.NewUserController().Add)
		// 验证码
		userLogin.GET("/captcha", system.NewUserLoginController().Captcha)
		// 验证码验证
		userLogin.GET("/captcha/verify", system.NewUserLoginController().CaptchaVerify)
		// 登录
		userLogin.POST("/login", system.NewUserLoginController().Login)
		// 登出
		userLogin.POST("/logout", system.NewUserLoginController().Logout)
	}

	// 文件上传
	upload := v1.Group("/upload")
	{
		// 上传用户头像
		upload.POST("/avatar", system.NewUploadController().Avatar)
	}

	// 用户管理
	user := v1.Group("/user")
	{
		// 获取用户信息
		user.GET("/info", system.NewUserController().Info)
		// 获取所有用户列表
		user.GET("/all", system.NewUserController().All)
		// 获取用户列表
		user.GET("/list", system.NewUserController().List)
		// 添加用户
		user.POST("/add", system.NewUserController().Add)
		// 更新用户详情信息
		user.PUT("/update", system.NewUserController().Update)
		// 删除用户
		user.DELETE("/delete", system.NewUserController().Delete)
		// 批量删除用户
		user.DELETE("/batchDelete", system.NewUserController().BatchDelete)
		// 更新用户状态
		user.PUT("/status", system.NewUserController().Status)
		// 重置用户密码
		user.PUT("/resetPwd", system.NewUserController().ResetPassword)
		// 更新用户密码
		user.PUT("/updatePwd", system.NewUserController().UpdatePassword)
		// 更新用户手机号码
		user.PUT("/updatePhone", system.NewUserController().UpdatePhone)
		// 更新用户邮箱
		user.PUT("/updateEmail", system.NewUserController().UpdateEmail)
	}

	// 角色管理
	role := v1.Group("/role")
	{
		// 获取所有角色列表
		role.GET("/all", system.NewRoleController().All)
		// 获取角色列表
		role.GET("/list", system.NewRoleController().List)
		// 添加角色
		role.POST("/add", system.NewRoleController().Add)
		// 更新角色
		role.PUT("/update", system.NewRoleController().Update)
		// 删除角色
		role.DELETE("/delete", system.NewRoleController().Delete)
		// 批量删除角色
		role.DELETE("/batchDelete", system.NewRoleController().BatchDelete)
		// 更新角色状态
		role.PUT("/status", system.NewRoleController().Status)
	}

	// 菜单管理
	menu := v1.Group("/menu")
	{
		// 获取所有菜单树
		menu.GET("/allTree", system.NewMenuController().AllTree)
		// 获取菜单树
		menu.GET("/tree", system.NewMenuController().Tree)
		// 添加菜单
		menu.POST("/add", system.NewMenuController().Add)
		// 更新菜单
		menu.PUT("/update", system.NewMenuController().Update)
		// 删除菜单
		menu.DELETE("/delete", system.NewMenuController().Delete)
		// 批量删除菜单
		menu.DELETE("/batchDelete", system.NewMenuController().BatchDelete)
		// 更新菜单状态
		menu.PUT("/status", system.NewMenuController().Status)

	}

	// 角色菜单关系管理
	roleMenuRel := v1.Group("/roleMenuRel")
	{
		// 获取角色关联的菜单列表
		roleMenuRel.GET("/list", system.NewMenuRelController().List)
		// 更新角色菜单关联关系
		roleMenuRel.PUT("/update", system.NewMenuRelController().Update)
	}
}
