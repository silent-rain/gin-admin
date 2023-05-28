// Package service 用户登录信息表
package service

import (
	"errors"

	"github.com/silent-rain/gin-admin/internal/app/system/cache"
	"github.com/silent-rain/gin-admin/internal/app/system/dao"
	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserLoginService 用户登录信息接口
type UserLoginService interface {
	List(ctx *gin.Context, req dto.QueryUserLoginReq) ([]model.UserLogin, int64, error)
	Add(ctx *gin.Context, bean model.UserLogin) (uint, error)
	Status(ctx *gin.Context, id, userId uint, status uint) (int64, error)
}

// 用户登录信息
type userLoginService struct {
	dao   dao.UserLogin
	cache cache.UserLoginCache
}

// NewUserLoginService 创建用户登录信息对象
func NewUserLoginService() *userLoginService {
	return &userLoginService{
		dao:   dao.NewUserLoginDao(),
		cache: cache.NewUserLoginCache(),
	}
}

// List 获取用户登录信息列表
func (s *userLoginService) List(ctx *gin.Context, req dto.QueryUserLoginReq) ([]model.UserLogin, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	return results, total, nil
}

// Add 添加用户登录信息
func (s *userLoginService) Add(ctx *gin.Context, bean model.UserLogin) (uint, error) {
	id, err := s.dao.Add(bean)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Status 更新用户登录信息状态
func (s *userLoginService) Status(ctx *gin.Context, id, userId uint, status uint) (int64, error) {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}

	// 禁用登录
	if status == 0 {
		s.cache.Set(userId, "")
	}
	return row, nil
}