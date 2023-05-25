/*用户API接口Token令牌表*/
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dao"
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// UserApiTokenService Token 令牌接口
type UserApiTokenService interface {
	List(ctx *gin.Context, req dto.QueryUserApiTokenReq) ([]dto.UserApiTokenResp, int64, error)
	Add(ctx *gin.Context, role model.UserApiToken) (uint, error)
	Update(ctx *gin.Context, role model.UserApiToken) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// Token 令牌
type userApiTokenService struct {
	dao dao.UserApiToken
}

// NewUserApiTokenService 创建 Token 令牌对象
func NewUserApiTokenService() *userApiTokenService {
	return &userApiTokenService{
		dao: dao.NewUserApiTokenDao(),
	}
}

// List 获取所有 Token 令牌列表
func (s *userApiTokenService) List(ctx *gin.Context, req dto.QueryUserApiTokenReq) ([]dto.UserApiTokenResp, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}

// Add 添加 Token 令牌
func (h *userApiTokenService) Add(ctx *gin.Context, role model.UserApiToken) (uint, error) {
	id, err := h.dao.Add(role)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Update 更新 Token 令牌
func (h *userApiTokenService) Update(ctx *gin.Context, role model.UserApiToken) (int64, error) {
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
