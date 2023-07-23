// Package dao Http协议接口管理
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/internal/app/api_auth/dto"
	"github.com/silent-rain/gin-admin/internal/app/api_auth/model"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// ApiHttp Http协议接口信息
type ApiHttp struct {
	mysql.DBRepo
}

// NewApiHttpDao 创建Http协议接口 Dao 对象
func NewApiHttpDao() *ApiHttp {
	return &ApiHttp{
		DBRepo: mysql.Instance(),
	}
}

// All 获取所有Http协议接口列表
func (d *ApiHttp) All() ([]model.ApiHttp, int64, error) {
	tx := d.GetDbR().Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&model.ApiHttp{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]model.ApiHttp, 0)
	if result := tx.Order("updated_at ASC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}

	return bean, total, nil
}

// List 查询Http协议接口列表
func (d *ApiHttp) List(req dto.QueryApiHttpReq) ([]model.ApiHttp, int64, error) {
	tx := d.GetDbR()
	if req.Method != "" {
		tx = tx.Where("method = ?", req.Method)
	}
	if req.Status != nil {
		tx = tx.Where("status = ?", *req.Status)
	}
	if req.Name != "" {
		tx = tx.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Uri != "" {
		tx = tx.Where("uri like ?", req.Uri+"%")
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&model.ApiHttp{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]model.ApiHttp, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// InfoByUri 获取Http协议接口信息
func (d *ApiHttp) InfoByUri(uri string) (model.ApiHttp, bool, error) {
	bean := model.ApiHttp{}
	result := d.GetDbR().Where("uri = ?", uri).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加Http协议接口
func (d *ApiHttp) Add(bean model.ApiHttp) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新Http协议接口
func (d *ApiHttp) Update(bean model.ApiHttp) (int64, error) {
	result := d.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除Http协议接口
func (d *ApiHttp) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.ApiHttp{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除Http协议接口
func (d *ApiHttp) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.ApiHttp, len(ids))
	for _, id := range ids {
		beans = append(beans, model.ApiHttp{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// UpdateStatus 更新状态
func (d *ApiHttp) UpdateStatus(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.ApiHttp{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// Children 通过父 ID 获取子配置列表
func (d *ApiHttp) Children(parentId uint) ([]model.ApiHttp, error) {
	beans := make([]model.ApiHttp, 0)
	result := d.GetDbR().Where("parent_id = ?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}

// GetUriListByToken 获取 Token 令牌对应的 URI 资源列表
func (d *ApiHttp) GetUriListByToken(token, uri string) (model.ApiHttp, bool, error) {
	bean := model.ApiHttp{}
	result := d.GetDbR().Model(&model.ApiHttp{}).
		Joins("left join api_role_http_rel arhr on arhr.api_id = api_http.id").
		Joins("left join perm_user_role_rel purr on purr.role_id = arhr.role_id").
		Joins("left join perm_user_api_token puat on puat.user_id = purr.user_id").
		Where("puat.token = ?", token).
		Where("api_http.uri = ?", uri).
		Group("api_http.id").
		First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}
