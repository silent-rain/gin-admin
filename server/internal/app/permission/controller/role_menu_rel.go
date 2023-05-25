/* 角色菜单
 */
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/service"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// 角色菜单关系
type roleMenuRelController struct {
	service service.RoleMenuRelService
}

// NewRoleMenuRelController 创建角色菜单关系对象
func NewRoleMenuRelController() *roleMenuRelController {
	return &roleMenuRelController{
		service: service.NewRoleMenuRelService(),
	}
}

// List 获取角色关联的菜单列表
func (c *roleMenuRelController) List(ctx *gin.Context) {
	req := dto.QueryRoleMenuRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	if req.RoleId == 0 && req.MenuId == 0 {
		log.New(ctx).WithField("data", req).Errorf("role_id/menu_id 不能同时为空")
		response.New(ctx).WithCode(errcode.ReqParameterParsingError).WithMsg("role_id/menu_id 不能同时为空")
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Update 更新角色菜单关联关系
func (c *roleMenuRelController) Update(ctx *gin.Context) {
	req := dto.UpdateRoleMenuRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if err := c.service.Update(ctx, req.RoleId, req.MenuIds); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
