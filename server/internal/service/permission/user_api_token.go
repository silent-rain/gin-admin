/*用户API接口Token令牌表*/
package permission

import (
	permissionDAO "github.com/silent-rain/gin-admin/internal/dao/permission"
	permissionDTO "github.com/silent-rain/gin-admin/internal/dto/permission"
	permissionModel "github.com/silent-rain/gin-admin/internal/model/permission"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	permissionVO "github.com/silent-rain/gin-admin/internal/vo/permission"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// UserApiTokenService Token 令牌接口
type UserApiTokenService interface {
	List(ctx *gin.Context, req permissionDTO.QueryUserApiTokenReq) ([]permissionVO.UserApiToken, int64, error)
	Add(ctx *gin.Context, role permissionModel.UserApiToken) (uint, error)
	Update(ctx *gin.Context, role permissionModel.UserApiToken) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// Token 令牌
type userApiTokenService struct {
	dao permissionDAO.UserApiToken
}

// NewUserApiTokenService 创建 Token 令牌对象
func NewUserApiTokenService() *userApiTokenService {
	return &userApiTokenService{
		dao: permissionDAO.NewUserApiTokenDao(),
	}
}

// List 获取所有 Token 令牌列表
func (s *userApiTokenService) List(ctx *gin.Context, req permissionDTO.QueryUserApiTokenReq) ([]permissionVO.UserApiToken, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}

// Add 添加 Token 令牌
func (h *userApiTokenService) Add(ctx *gin.Context, role permissionModel.UserApiToken) (uint, error) {
	id, err := h.dao.Add(role)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Update 更新 Token 令牌
func (h *userApiTokenService) Update(ctx *gin.Context, role permissionModel.UserApiToken) (int64, error) {
	row, err := h.dao.Update(role)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// Delete 删除 Token 令牌
func (h *userApiTokenService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}
	return row, nil
}

// BatchDelete 批量删除 Token 令牌
func (h *userApiTokenService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}
	return row, nil
}

// Status 更新 Token 令牌状态
func (h *userApiTokenService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}
	return row, nil
}
