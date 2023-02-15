/*角色
 */
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/code_errors"
	"gin-admin/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// RoleService 角色接口
type RoleService interface {
	All(ctx *gin.Context) ([]systemModel.Role, int64, error)
	List(ctx *gin.Context, req systemDTO.QueryRoleReq) ([]systemModel.Role, int64, error)
	Add(ctx *gin.Context, role systemModel.Role) (uint, error)
	Update(ctx *gin.Context, role systemModel.Role) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 角色
type roleService struct {
	dao systemDAO.Role
}

// NewRoleService 创建角色对象
func NewRoleService() *roleService {
	return &roleService{
		dao: systemDAO.NewRoleDao(),
	}
}

// All 获取所有角色列表
func (s *roleService) All(ctx *gin.Context) ([]systemModel.Role, int64, error) {
	roles, total, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBQueryError).Errorf("%v", err)
		return nil, 0, code_errors.New(code_errors.DBQueryError)

	}
	return roles, total, nil
}

// List 获取所有角色列表
func (s *roleService) List(ctx *gin.Context, req systemDTO.QueryRoleReq) ([]systemModel.Role, int64, error) {
	roles, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBQueryError).Errorf("%v", err)
		return nil, 0, code_errors.New(code_errors.DBQueryError)

	}
	return roles, total, nil
}

// Add 添加角色
func (h *roleService) Add(ctx *gin.Context, role systemModel.Role) (uint, error) {
	_, ok, err := h.dao.InfoByName(role.Name)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBQueryError).Errorf("%v", err)
		return 0, code_errors.New(code_errors.DBQueryError)
	}
	if ok {
		log.New(ctx).WithCode(code_errors.DBDataExistError).Errorf("%v", err)
		return 0, code_errors.New(code_errors.DBDataExistError).WithMsg("角色已存在")
	}

	id, err := h.dao.Add(role)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBAddError).Errorf("%v", err)
		return 0, code_errors.New(code_errors.DBAddError)
	}
	return id, nil
}

// Update 更新角色
func (h *roleService) Update(ctx *gin.Context, role systemModel.Role) (int64, error) {
	row, err := h.dao.Update(role)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBUpdateError).Errorf("%v", err)
		return 0, code_errors.New(code_errors.DBUpdateError)
	}
	return row, nil
}

// Delete 删除角色
func (h *roleService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBDeleteError).Errorf("%v", err)
		return 0, code_errors.New(code_errors.DBDeleteError)
	}
	return row, nil
}

// BatchDelete 批量删除角色
func (h *roleService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBBatchDeleteError).Errorf("%v", err)
		return 0, code_errors.New(code_errors.DBBatchDeleteError)
	}
	return row, nil
}

// Status 更新角色状态
func (h *roleService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBUpdateStatusError).Errorf("%v", err)
		return 0, code_errors.New(code_errors.DBUpdateStatusError)
	}
	return row, nil
}
