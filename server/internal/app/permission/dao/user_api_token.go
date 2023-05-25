// Package dao 用户API接口Token令牌表
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// UserApiToken Token令牌接口
type UserApiToken interface {
	List(req dto.QueryUserApiTokenReq) ([]dto.UserApiTokenResp, int64, error)
	Info(token string) (model.UserApiToken, bool, error)
	Add(bean model.UserApiToken) (uint, error)
	Update(bean model.UserApiToken) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// Token 令牌
type userApiToken struct {
	mysql.DBRepo
}

// NewUserApiTokenDao 创建 Token 令牌对象
func NewUserApiTokenDao() *userApiToken {
	return &userApiToken{
		DBRepo: mysql.Instance(),
	}
}

// List 查询 Token 令牌列表
func (d *userApiToken) List(req dto.QueryUserApiTokenReq) ([]dto.UserApiTokenResp, int64, error) {
	tx := d.GetDbR().Model(&model.UserApiToken{}).
		Select("perm_user_api_token.*, perm_user.nickname").
		Joins("left join perm_user on perm_user.id = perm_user_api_token.user_id")
	if req.UserId != nil {
		tx = tx.Where("perm_user_api_token.user_id = ?", *req.UserId)
	}
	if req.Nickname != "" {
		tx = tx.Where("perm_user.nickname like ?", req.Nickname+"%")
	}
	if req.Status != nil {
		tx = tx.Where("perm_user_api_token.status = ?", *req.Status)
	}

	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&model.UserApiToken{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]dto.UserApiTokenResp, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).
		Order("sort DESC").Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// Info 获取 Token 令牌信息
func (d *userApiToken) Info(token string) (model.UserApiToken, bool, error) {
	bean := model.UserApiToken{}
	result := d.GetDbR().Where("token = ?", token).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加 Token 令牌
func (d *userApiToken) Add(bean model.UserApiToken) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新 Token 令牌
func (d *userApiToken) Update(bean model.UserApiToken) (int64, error) {
	result := d.GetDbW().Select("permission", "passphrase", "note", "status").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除 Token 令牌
func (d *userApiToken) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.UserApiToken{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除 Token 令牌
func (d *userApiToken) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.UserApiToken, len(ids))
	for _, id := range ids {
		beans = append(beans, model.UserApiToken{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *userApiToken) Status(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.UserApiToken{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
