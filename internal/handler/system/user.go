/*
 * @Author: silent-rain
 * @Date: 2023-01-08 21:24:21
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 23:52:51
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/user.go
 * @Descripttion: 用户管理
 */
package system

import (
	systemDao "gin-admin/internal/dao/system"
	systemDto "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserHandlerImpl 用户管理对象
var UserHandlerImpl = new(userHandler)

// 用户管理
type userHandler struct {
}

// List 获取用户列表
func (h *userHandler) List(ctx *gin.Context) {
	req := new(systemDto.UserQueryReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	roles, total, err := systemDao.UserImpl.List(*req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(roles, total).Json()
}

// Delete 删除用户
func (h *userHandler) Delete(ctx *gin.Context) {
	req := new(systemDto.UserDeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := systemDao.UserImpl.Delete(req.ID)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新用户状态
func (h *userHandler) Status(ctx *gin.Context) {
	req := new(systemDto.UserStatusReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := systemDao.UserImpl.Status(req.ID, req.Status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UserInfo 获取用户信息
func (h *userHandler) UserInfo(ctx *gin.Context) {
	// zap.S().Error("===================", "xxxxxxxxxxxxxxxx")
	// log.Debug(ctx, "xxxxxxxx", zap.String("method", ctx.Request.Method))
	// log.New(ctx).
	// 	WithCode(statuscode.CaptchaNotFoundError).
	// 	Debug("xxxxxxxxdebug", zap.String("method", ctx.Request.Method))

	log.New(ctx).
		WithCode(statuscode.DbQueryEmptyError).
		WithField("aaa", "AAAAAA").
		WithField("xxx", 111).
		Debugf("xxxxxxxxdebug:  %v", "xxxxxxxxxxxAAA")
	response.New(ctx).Json()
}
