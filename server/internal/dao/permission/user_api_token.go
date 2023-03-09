/*用户API接口Token令牌表*/
package permissionDAO

import (
	permissionDTO "gin-admin/internal/dto/permission"
	permissionModel "gin-admin/internal/model/permission"
	"gin-admin/internal/pkg/repository/mysql"
	permissionVO "gin-admin/internal/vo/permission"

	"gorm.io/gorm"
)

// UserApiToken Token令牌接口
type UserApiToken interface {
	List(req permissionDTO.QueryUserApiTokenReq) ([]permissionVO.UserApiToken, int64, error)
	Add(bean permissionModel.UserApiToken) (uint, error)
	Update(bean permissionModel.UserApiToken) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// Token 令牌
type userApiToken struct {
	db mysql.DBRepo
}

// NewUserApiTokenDao 创建 Token 令牌对象
func NewUserApiTokenDao() *userApiToken {
	return &userApiToken{
		db: mysql.Instance(),
	}
}

// List 查询 Token 令牌列表
func (d *userApiToken) List(req permissionDTO.QueryUserApiTokenReq) ([]permissionVO.UserApiToken, int64, error) {
	tx := d.db.GetDbR().Model(&permissionModel.UserApiToken{}).
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
	if result := tx.Model(&permissionModel.UserApiToken{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]permissionVO.UserApiToken, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).
		Order("sort DESC").Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// Add 添加 Token 令牌
func (d *userApiToken) Add(bean permissionModel.UserApiToken) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新 Token 令牌
func (d *userApiToken) Update(bean permissionModel.UserApiToken) (int64, error) {
	result := d.db.GetDbW().Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除 Token 令牌
func (d *userApiToken) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&permissionModel.UserApiToken{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除 Token 令牌
func (d *userApiToken) BatchDelete(ids []uint) (int64, error) {
	beans := make([]permissionModel.UserApiToken, len(ids))
	for _, id := range ids {
		beans = append(beans, permissionModel.UserApiToken{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *userApiToken) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&permissionModel.UserApiToken{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
