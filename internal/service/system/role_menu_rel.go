/* 角色菜单
 */
package service

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

type RoleMenuRelService interface {
	List(ctx *gin.Context, req systemDTO.QueryRoleMenuRelReq) *response.ResponseAPI
	Update(ctx *gin.Context, roleId uint, menuIds []uint) *response.ResponseAPI
}

// 角色菜单关系
type roleMenuRelService struct {
	dao systemDAO.RoleMenuRel
}

// NewRoleMenuRelService 创建角色菜单关系 Handler 对象
func NewRoleMenuRelService() *roleMenuRelService {
	return &roleMenuRelService{
		dao: systemDAO.NewRoleMenuRelDao(),
	}
}

// List 获取角色关联的菜单列表
func (s *roleMenuRelService) List(ctx *gin.Context, req systemDTO.QueryRoleMenuRelReq) *response.ResponseAPI {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	return response.New().WithDataList(results, total)
}

// Update 更新角色菜单关联关系
func (h *roleMenuRelService) Update(ctx *gin.Context, roleId uint, menuIds []uint) *response.ResponseAPI {
	if err := h.dao.Update(roleId, menuIds); err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBUpdateError)
	}
	return response.New()
}
