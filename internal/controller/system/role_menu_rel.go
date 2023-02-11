/* 角色菜单
 */
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 角色菜单关系
type roleMenuRelHandler struct {
	dao systemDAO.RoleMenuRel
}

// 创建角色菜单关系 Handler 对象
func NewMenuRelHandler() *roleMenuRelHandler {
	return &roleMenuRelHandler{
		dao: systemDAO.NewRoleMenuRelDao(),
	}
}

// List 获取角色关联的菜单列表
func (h *roleMenuRelHandler) List(ctx *gin.Context) {
	req := new(systemDTO.QueryRoleMenuRelReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	if req.RoleId == 0 && req.MenuId == 0 {
		log.New(ctx).WithField("data", req).Errorf("role_id/menu_id 不能同时为空")
		return
	}
	results, total, err := h.dao.List(*req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Update 更新角色菜单关联关系
func (h *roleMenuRelHandler) Update(ctx *gin.Context) {
	req := new(systemDTO.UpdateRoleMenuRelReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	if err := h.dao.Update(req.RoleId, req.MenuIds); err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).Json()
}
