/*
 * @Author: silent-rain
 * @Date: 2023-01-08 21:24:21
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-11 22:20:54
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/user.go
 * @Descripttion: 用户管理
 */
package system

import (
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// UserManageImpl 用户管理对象
var UserManageImpl = new(userManageHandler)

// 用户管理
type userManageHandler struct {
}

// UserInfo 获取用户信息
func (h *userManageHandler) UserInfo(ctx *gin.Context) {
	// zap.S().Error("===================", "xxxxxxxxxxxxxxxx")
	// log.Debug(ctx, "xxxxxxxx", zap.String("method", ctx.Request.Method))
	// log.New(ctx).
	// 	WithCode(statuscode.CaptchaNotFoundError).
	// 	Debug("xxxxxxxxdebug", zap.String("method", ctx.Request.Method))

	log.New(ctx).
		WithCode(statuscode.DbQueryEmptyError).WithAny("aaaaaaa", "aaaaaaaaa").
		Debugf("xxxxxxxxdebug:  %v", "xxxxxxxxxxxAAA")
	response.New(ctx).Json()
}
