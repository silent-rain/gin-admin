/*
 * @Author: silent-rain
 * @Date: 2023-01-08 21:24:21
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 17:20:59
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/user.go
 * @Descripttion: 用户管理
 */
package system

import (
	systemDao "gin-admin/internal/dao/system"
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/conf"
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

// All 获取所有用户列表
func (h *userHandler) All(ctx *gin.Context) {
	results, total, err := systemDao.UserImpl.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// List 获取用户列表
func (h *userHandler) List(ctx *gin.Context) {
	req := new(systemDto.UserQueryReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	results, total, err := systemDao.UserImpl.List(*req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// UpdateDetails 更新用户详情信息
func (h *userHandler) UpdateDetails(ctx *gin.Context) {
	req := new(systemDto.UserUpdateDetailsReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 数据转换
	user := new(systemModel.User)
	if err := utils.ApiJsonConvertJson(ctx, req, user); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	roleIds := req.RoleIds
	if err := systemDao.UserImpl.UpdateDetails(*user, roleIds); err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).Json()
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
		log.New(ctx).WithCode(statuscode.DbSetStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbSetStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdatePassword 更新密码
func (h *userHandler) UpdatePassword(ctx *gin.Context) {
	req := new(systemDto.UserUpdatePasswordReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 密码加密
	req.OldPassword = utils.Md5(req.OldPassword)
	req.NewPassword = utils.Md5(req.NewPassword)

	// 用户密码验证
	ok, err := systemDao.UserImpl.ExistUserPassword(req.ID, req.OldPassword)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.UserOldPasswordError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.UserOldPasswordError).Json()
		return
	}

	row, err := systemDao.UserImpl.UpdatePassword(req.ID, req.NewPassword)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// ResetPassword 重置密码
func (h *userHandler) ResetPassword(ctx *gin.Context) {
	req := new(systemDto.UserResetPasswordReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 默认密码加密
	password := utils.Md5(conf.ServerUserDefaultPwd)
	row, err := systemDao.UserImpl.ResetPassword(req.ID, password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbResetError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbResetError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdatePhone 更新手机号码
func (h *userHandler) UpdatePhone(ctx *gin.Context) {
	req := new(systemDto.UserUpdatePhoneReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := systemDao.UserImpl.UpdatePhone(req.ID, req.Phone)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdateEmail 更新邮箱
func (h *userHandler) UpdateEmail(ctx *gin.Context) {
	req := new(systemDto.UserUpdateEmailReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := systemDao.UserImpl.UpdateEmail(req.ID, req.Email)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
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
