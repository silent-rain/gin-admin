// Package controller 角色与Http协议接口关联管理
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/api_auth/dto"
	"github.com/silent-rain/gin-admin/internal/app/api_auth/service"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// 角色与Http协议接口关系
type apiRoleHttpRelController struct {
	service *service.ApiRoleHttpRelService
}

// NewApiRoleHttpRelController 创建角色与Http协议接口关系对象
func NewApiRoleHttpRelController() *apiRoleHttpRelController {
	return &apiRoleHttpRelController{
		service: service.NewApiRoleHttpRelService(),
	}
}

// List 获取角色与Http协议接口关系列表
func (c *apiRoleHttpRelController) List(ctx *gin.Context) {
	req := dto.QueryApiRoleHttpRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	if req.RoleId == 0 && req.ApiId == 0 {
		log.New(ctx).WithField("data", req).Errorf("role_id/api_id 不能同时为空")
		response.New(ctx).WithCode(errcode.ReqParameterParsingError).WithMsg("role_id/api_id 不能同时为空")
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Update 更新角色与Http协议接口关系
func (c *apiRoleHttpRelController) Update(ctx *gin.Context) {
	req := dto.UpdateApiRoleHttpRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if err := c.service.Update(ctx, req.RoleId, req.ApiIds); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
