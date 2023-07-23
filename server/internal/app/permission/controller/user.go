// Package controller 用户管理
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/app/permission/service"
	DTO "github.com/silent-rain/gin-admin/internal/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/pkg/constant"
	"github.com/silent-rain/gin-admin/pkg/md5"
	"github.com/silent-rain/gin-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 用户管理
type userController struct {
	service *service.UserService
}

// 创建用户对象
func NewUserController() *userController {
	return &userController{
		service: service.NewUserService(),
	}
}

// All 获取所有用户列表
func (c *userController) All(ctx *gin.Context) {
	results, total, err := c.service.All(ctx)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// List 获取用户列表
func (c *userController) List(ctx *gin.Context) {
	req := dto.QueryUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Add 添加用户
func (c *userController) Add(ctx *gin.Context) {
	req := dto.AddUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if err := c.service.Add(ctx, req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新用户详情信息
func (c *userController) Update(ctx *gin.Context) {
	req := dto.UpdateUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	// 数据转换
	user := model.User{}
	if err := http.ApiJsonConvertJson(ctx, req, &user); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	roleIds := req.RoleIds

	if err := c.service.Update(ctx, user, roleIds); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除用户
func (c *userController) Delete(ctx *gin.Context) {
	req := DTO.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if _, err := c.service.Delete(ctx, req.ID); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// BatchDelete 批量删除用户
func (c *userController) BatchDelete(ctx *gin.Context) {
	req := DTO.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if _, err := c.service.BatchDelete(ctx, req.Ids); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdateStatus 更新用户状态
func (c *userController) UpdateStatus(ctx *gin.Context) {
	req := DTO.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if _, err := c.service.UpdateStatus(ctx, req.ID, req.Status); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdatePassword 更新密码
func (c *userController) UpdatePassword(ctx *gin.Context) {
	req := dto.UpdateUserPasswordReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	// 密码加密
	req.OldPassword = md5.EncryptMd5(req.OldPassword)
	req.NewPassword = md5.EncryptMd5(req.NewPassword)

	if _, err := c.service.UpdatePassword(ctx, req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// ResetPassword 重置密码
func (c *userController) ResetPassword(ctx *gin.Context) {
	req := dto.ResetUserPasswordReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	// 默认密码加密
	password := md5.EncryptMd5(constant.ServerUserDefaultPwd)

	if _, err := c.service.ResetPassword(ctx, req.ID, password); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdatePhone 更新手机号码
func (c *userController) UpdatePhone(ctx *gin.Context) {
	req := dto.UpdateUserPhoneReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if _, err := c.service.UpdatePhone(ctx, req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdateEmail 更新邮箱
func (c *userController) UpdateEmail(ctx *gin.Context) {
	req := dto.UpdateUserEmailReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if _, err := c.service.UpdateEmail(ctx, req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Info 获取用户信息
func (c *userController) Info(ctx *gin.Context) {
	userId := core.Context(ctx).UserId

	result, err := c.service.Info(ctx, userId)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}
