// Package service 角色
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dao"
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// RoleService 角色接口
type RoleService interface {
	All(ctx *gin.Context) ([]model.Role, int64, error)
	List(ctx *gin.Context, req dto.QueryRoleReq) ([]model.Role, int64, error)
	Add(ctx *gin.Context, role model.Role) (uint, error)
	Update(ctx *gin.Context, role model.Role) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 角色
type roleService struct {
	dao dao.Role
}

// NewRoleService 创建角色对象
func NewRoleService() *roleService {
	return &roleService{
		dao: dao.NewRoleDao(),
	}
}

// All 获取所有角色列表
func (s *roleService) All(ctx *gin.Context) ([]model.Role, int64, error) {
	roles, total, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return roles, total, nil
}

// List 获取所有角色列表
func (s *roleService) List(ctx *gin.Context, req dto.QueryRoleReq) ([]model.Role, int64, error) {
	roles, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return roles, total, nil
}

// Add 添加角色
func (h *roleService) Add(ctx *gin.Context, role model.Role) (uint, error) {
	_, ok, err := h.dao.InfoByName(role.Name)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("角色已存在")
		return 0, errcode.DBDataExistError.WithMsg("角色已存在")
	}

	id, err := h.dao.Add(role)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Update 更新角色
func (h *roleService) Update(ctx *gin.Context, role model.Role) (int64, error) {
	row, err := h.dao.Update(role)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// Delete 删除角色
func (h *roleService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}
	return row, nil
}

// BatchDelete 批量删除角色
func (h *roleService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}
	return row, nil
}

// Status 更新角色状态
func (h *roleService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}
	return row, nil
}
