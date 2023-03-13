/*角色与Http协议接口关联表*/
package apiauth

import (
	apiAuthDAO "gin-admin/internal/dao/api_auth"
	apiAuthDTO "gin-admin/internal/dto/api_auth"
	apiAuthModel "gin-admin/internal/model/api_auth"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// ApiRoleHttpRelService 角色Http协议接口关系接口
type ApiRoleHttpRelService interface {
	List(ctx *gin.Context, req apiAuthDTO.QueryApiRoleHttpRelReq) ([]apiAuthModel.ApiRoleHttpRel, int64, error)
	Update(ctx *gin.Context, roleId uint, menuIds []uint) error
}

// 角色Http协议接口关系
type roleMenuRelService struct {
	dao apiAuthDAO.ApiRoleHttpRel
}

// NewApiRoleHttpRelService 创建角色Http协议接口关系对象
func NewApiRoleHttpRelService() *roleMenuRelService {
	return &roleMenuRelService{
		dao: apiAuthDAO.NewApiRoleHttpRelDao(),
	}
}

// List 获取角色Http协议接口关系列表
func (s *roleMenuRelService) List(ctx *gin.Context, req apiAuthDTO.QueryApiRoleHttpRelReq) ([]apiAuthModel.ApiRoleHttpRel, int64, error) {
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
