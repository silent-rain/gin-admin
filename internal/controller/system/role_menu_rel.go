/* 角色菜单
 */
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	service "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 角色菜单关系
type roleMenuRelController struct {
	service service.RoleMenuRelService
}

// NewMenuRelController 创建角色菜单关系对象
func NewMenuRelController() *roleMenuRelController {
	return &roleMenuRelController{
		service: service.NewRoleMenuRelService(),
	}
}

// List 获取角色关联的菜单列表
func (c *roleMenuRelController) List(ctx *gin.Context) {
	req := systemDTO.QueryRoleMenuRelReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}
	if req.RoleId == 0 && req.MenuId == 0 {
		log.New(ctx).WithField("data", req).Errorf("role_id/menu_id 不能同时为空")
		response.New().WithCode(statuscode.ReqParameterParsingError).WithMsg("role_id/menu_id 不能同时为空").Json(ctx)
		return
	}

	c.service.List(ctx, req).Json(ctx)
}

// Update 更新角色菜单关联关系
func (c *roleMenuRelController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateRoleMenuRelReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Update(ctx, req.RoleId, req.MenuIds).Json(ctx)
}
