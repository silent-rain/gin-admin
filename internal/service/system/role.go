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

type RoleService interface {
	All(ctx *gin.Context)
	List(ctx *gin.Context, req systemDTO.QueryRoleReq)
	Add(ctx *gin.Context, role systemModel.Role)
	Update(ctx *gin.Context, role systemModel.Role)
	Delete(ctx *gin.Context, id uint)
	BatchDelete(ctx *gin.Context, ids []uint)
	Status(ctx *gin.Context, id uint, status uint)
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
func (s *roleService) All(ctx *gin.Context) {
	roles, total, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(roles, total).Json()
}

// List 获取用角色列表
func (s *roleService) List(ctx *gin.Context, req systemDTO.QueryRoleReq) {
	roles, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(roles, total).Json()
}

// Add 添加角色
func (h *roleService) Add(ctx *gin.Context, role systemModel.Role) {
	_, ok, err := h.dao.InfoByName(role.Name)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	if ok {
		log.New(ctx).WithCode(statuscode.DBDataExistError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBDataExistError).WithMsg("角色已存在").Json()
		return
	}

	if _, err := h.dao.Add(role); err != nil {
		log.New(ctx).WithCode(statuscode.DBAddError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBAddError).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新角色
func (h *roleService) Update(ctx *gin.Context, role systemModel.Role) {
	row, err := h.dao.Update(role)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Delete 删除角色
func (h *roleService) Delete(ctx *gin.Context, id uint) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// BatchDelete 批量删除角色
func (h *roleService) BatchDelete(ctx *gin.Context, ids []uint) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBBatchDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBBatchDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新角色状态
func (h *roleService) Status(ctx *gin.Context, id uint, status uint) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}
