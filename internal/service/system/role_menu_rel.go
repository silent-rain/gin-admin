/* 角色菜单
 */
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// RoleMenuRelService 角色菜单接口
type RoleMenuRelService interface {
	List(ctx *gin.Context, req systemDTO.QueryRoleMenuRelReq) ([]systemModel.RoleMenuRel, int64, error)
	Update(ctx *gin.Context, roleId uint, menuIds []uint) error
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
func (s *roleMenuRelService) List(ctx *gin.Context, req systemDTO.QueryRoleMenuRelReq) ([]systemModel.RoleMenuRel, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}
	return results, total, nil
}

// Update 更新角色菜单关联关系
func (h *roleMenuRelService) Update(ctx *gin.Context, roleId uint, menuIds []uint) error {
	if err := h.dao.Update(roleId, menuIds); err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return errcode.New(errcode.DBUpdateError)
	}
	return nil
}
