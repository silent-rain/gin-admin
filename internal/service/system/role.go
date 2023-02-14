/*角色
 */
package service

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// RoleService 角色接口
type RoleService interface {
	All(ctx *gin.Context) *response.ResponseAPI
	List(ctx *gin.Context, req systemDTO.QueryRoleReq) *response.ResponseAPI
	Add(ctx *gin.Context, role systemModel.Role) *response.ResponseAPI
	Update(ctx *gin.Context, role systemModel.Role) *response.ResponseAPI
	Delete(ctx *gin.Context, id uint) *response.ResponseAPI
	BatchDelete(ctx *gin.Context, ids []uint) *response.ResponseAPI
	Status(ctx *gin.Context, id uint, status uint) *response.ResponseAPI
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
func (s *roleService) All(ctx *gin.Context) *response.ResponseAPI {
	roles, total, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)

	}
	return response.New().WithDataList(roles, total)
}

// List 获取所有角色列表
func (s *roleService) List(ctx *gin.Context, req systemDTO.QueryRoleReq) *response.ResponseAPI {
	roles, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)

	}
	return response.New().WithDataList(roles, total)
}

// Add 添加角色
func (h *roleService) Add(ctx *gin.Context, role systemModel.Role) *response.ResponseAPI {
	_, ok, err := h.dao.InfoByName(role.Name)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	if ok {
		log.New(ctx).WithCode(statuscode.DBDataExistError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBDataExistError).WithMsg("角色已存在")
	}

	if _, err := h.dao.Add(role); err != nil {
		log.New(ctx).WithCode(statuscode.DBAddError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBAddError)
	}
	return response.New()
}

// Update 更新角色
func (h *roleService) Update(ctx *gin.Context, role systemModel.Role) *response.ResponseAPI {
	row, err := h.dao.Update(role)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBUpdateError)
	}
	return response.New().WithData(row)
}

// Delete 删除角色
func (h *roleService) Delete(ctx *gin.Context, id uint) *response.ResponseAPI {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBDeleteError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBDeleteError)
	}
	return response.New().WithData(row)
}

// BatchDelete 批量删除角色
func (h *roleService) BatchDelete(ctx *gin.Context, ids []uint) *response.ResponseAPI {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBBatchDeleteError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBBatchDeleteError)
	}
	return response.New().WithData(row)
}

// Status 更新角色状态
func (h *roleService) Status(ctx *gin.Context, id uint, status uint) *response.ResponseAPI {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateStatusError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBUpdateStatusError)
	}
	return response.New().WithData(row)
}
