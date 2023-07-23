// Package controller 用户API接口Token令牌表
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/app/permission/service"
	DTO "github.com/silent-rain/gin-admin/internal/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	"github.com/silent-rain/gin-admin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Token 令牌
type userApiTokenController struct {
	service *service.UserApiTokenService
}

// NewUserApiTokenController 创建 Token 令牌对象
func NewUserApiTokenController() *userApiTokenController {
	return &userApiTokenController{
		service: service.NewUserApiTokenService(),
	}
}

// List 获取用 Token 令牌列表
func (c *userApiTokenController) List(ctx *gin.Context) {
	req := dto.QueryUserApiTokenReq{}
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

// Add 添加 Token 令牌
func (c *userApiTokenController) Add(ctx *gin.Context) {
	req := dto.AddUserApiTokenReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := model.UserApiToken{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	// 生成用户API接口Token
	bean.Token = utils.GenerateTUserApiToken()
	_, err := c.service.Add(ctx, bean)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新 Token 令牌
func (c *userApiTokenController) Update(ctx *gin.Context) {
	req := dto.UpdateUserApiTokenReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := model.UserApiToken{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Update(ctx, bean)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除 Token 令牌
func (c *userApiTokenController) Delete(ctx *gin.Context) {
	req := DTO.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Delete(ctx, req.ID)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// BatchDelete 批量删除 Token 令牌
func (c *userApiTokenController) BatchDelete(ctx *gin.Context) {
	req := DTO.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.BatchDelete(ctx, req.Ids)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdateStatus 更新 Token 令牌状态
func (c *userApiTokenController) UpdateStatus(ctx *gin.Context) {
	req := DTO.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.UpdateStatus(ctx, req.ID, req.Status)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
