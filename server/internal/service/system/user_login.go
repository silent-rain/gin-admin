/*用户登录信息表*/
package system

import (
	"errors"

	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserLoginService 用户登录信息接口
type UserLoginService interface {
	List(ctx *gin.Context, req systemDTO.QueryUserLoginReq) ([]systemModel.UserLogin, int64, error)
	Add(ctx *gin.Context, bean systemModel.UserLogin) (uint, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 用户登录信息
type userLoginService struct {
	dao systemDAO.UserLogin
}

// NewUserLoginService 创建用户登录信息对象
func NewUserLoginService() *userLoginService {
	return &userLoginService{
		dao: systemDAO.NewUserLoginDao(),
	}
}

// List 获取用户登录信息列表
func (s *userLoginService) List(ctx *gin.Context, req systemDTO.QueryUserLoginReq) ([]systemModel.UserLogin, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}
	return results, total, nil
}

// Add 添加用户登录信息
func (s *userLoginService) Add(ctx *gin.Context, bean systemModel.UserLogin) (uint, error) {
	id, err := s.dao.Add(bean)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}

// Status 更新用户登录信息状态
func (s *userLoginService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateStatusError)
	}
	return row, nil
}
