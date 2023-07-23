// Package controller 用户登录表
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/app/system/service"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 用户登录信息
type userLoginController struct {
	service *service.UserLoginService
}

// NewUserLoginController 创建用户登录信息对象
func NewUserLoginController() *userLoginController {
	return &userLoginController{
		service: service.NewUserLoginService(),
	}
}

// List 获取用户登录信息列表
func (c *userLoginController) List(ctx *gin.Context) {
	req := dto.QueryUserLoginReq{}
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

// UpdateStatus 更新用户登录信息状态
func (c *userLoginController) UpdateStatus(ctx *gin.Context) {
	req := dto.UpdateUserLoginStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	row, err := c.service.UpdateStatus(ctx, req.ID, req.UserId, req.Status)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}
