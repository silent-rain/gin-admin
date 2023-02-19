/* 角色菜单
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// 角色菜单关系
type roleMenuRelController struct {
	service systemService.RoleMenuRelService
}

// NewRoleMenuRelController 创建角色菜单关系对象
func NewRoleMenuRelController() *roleMenuRelController {
	return &roleMenuRelController{
		service: systemService.NewRoleMenuRelService(),
	}
}

// List 获取角色关联的菜单列表
func (c *roleMenuRelController) List(ctx *gin.Context) {
	req := systemDTO.QueryRoleMenuRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	if req.RoleId == 0 && req.MenuId == 0 {
		log.New(ctx).WithField("data", req).Errorf("role_id/menu_id 不能同时为空")
		response.New(ctx).WithCode(errcode.ReqParameterParsingError).WithMsg("role_id/menu_id 不能同时为空")
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Update 更新角色菜单关联关系
func (c *roleMenuRelController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateRoleMenuRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	if err := c.service.Update(ctx, req.RoleId, req.MenuIds); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}