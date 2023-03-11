/*角色
 */
package permission

import (
	permissionDAO "gin-admin/internal/dao/permission"
	permissionDTO "gin-admin/internal/dto/permission"
	permissionModel "gin-admin/internal/model/permission"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// RoleService 角色接口
type RoleService interface {
	All(ctx *gin.Context) ([]permissionModel.Role, int64, error)
	List(ctx *gin.Context, req permissionDTO.QueryRoleReq) ([]permissionModel.Role, int64, error)
	Add(ctx *gin.Context, role permissionModel.Role) (uint, error)
	Update(ctx *gin.Context, role permissionModel.Role) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 角色
type roleService struct {
	dao permissionDAO.Role
}

// NewRoleService 创建角色对象
func NewRoleService() *roleService {
	return &roleService{
		dao: permissionDAO.NewRoleDao(),
	}
}

// All 获取所有角色列表
func (s *roleService) All(ctx *gin.Context) ([]permissionModel.Role, int64, error) {
	roles, total, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return roles, total, nil
}

// List 获取所有角色列表
func (s *roleService) List(ctx *gin.Context, req permissionDTO.QueryRoleReq) ([]permissionModel.Role, int64, error) {
	roles, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return roles, total, nil
}

// Add 添加角色
func (h *roleService) Add(ctx *gin.Context, role permissionModel.Role) (uint, error) {
	_, ok, err := h.dao.InfoByName(role.Name)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBQueryError)
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("角色已存在")
		return 0, errcode.New(errcode.DBDataExistError).WithMsg("角色已存在")
	}

	id, err := h.dao.Add(role)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}

// Update 更新角色
func (h *roleService) Update(ctx *gin.Context, role permissionModel.Role) (int64, error) {
	row, err := h.dao.Update(role)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateError)
	}
	return row, nil
}

// Delete 删除角色
func (h *roleService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBDeleteError)
	}
	return row, nil
}

// BatchDelete 批量删除角色
func (h *roleService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBBatchDeleteError)
	}
	return row, nil
}

// Status 更新角色状态
func (h *roleService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateStatusError)
	}
	return row, nil
}
