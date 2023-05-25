/*角色与Http协议接口关联表*/
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/api_auth/dao"
	"github.com/silent-rain/gin-admin/internal/app/api_auth/dto"
	"github.com/silent-rain/gin-admin/internal/app/api_auth/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// ApiRoleHttpRelService 角色Http协议接口关系接口
type ApiRoleHttpRelService interface {
	List(ctx *gin.Context, req dto.QueryApiRoleHttpRelReq) ([]model.ApiRoleHttpRel, int64, error)
	Update(ctx *gin.Context, roleId uint, menuIds []uint) error
}

// 角色Http协议接口关系
type roleMenuRelService struct {
	dao dao.ApiRoleHttpRel
}

// NewApiRoleHttpRelService 创建角色Http协议接口关系对象
func NewApiRoleHttpRelService() *roleMenuRelService {
	return &roleMenuRelService{
		dao: dao.NewApiRoleHttpRelDao(),
	}
}

// List 获取角色Http协议接口关系列表
func (s *roleMenuRelService) List(ctx *gin.Context, req dto.QueryApiRoleHttpRelReq) ([]model.ApiRoleHttpRel, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	return results, total, nil
}

// Update 更新角色Http协议接口关系
func (h *roleMenuRelService) Update(ctx *gin.Context, roleId uint, apiIds []uint) error {
	if err := h.dao.Update(roleId, apiIds); err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return errcode.DBUpdateError
	}
	return nil
}
