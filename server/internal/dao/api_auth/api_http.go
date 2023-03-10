/*Http协议接口管理表*/
package apiauth

import (
	apiAuthDTO "gin-admin/internal/dto/api_auth"
	apiAuthModel "gin-admin/internal/model/api_auth"
	"gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// ApiHttp Http协议接口信息接口
type ApiHttp interface {
	All() ([]apiAuthModel.ApiHttp, int64, error)
	List(req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error)
	Add(bean apiAuthModel.ApiHttp) (uint, error)
	Update(bean apiAuthModel.ApiHttp) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// Http协议接口信息
type apiAuth struct {
	db mysql.DBRepo
}

// NewApiHttpDao 创建Http协议接口 Dao 对象
func NewApiHttpDao() *apiAuth {
	return &apiAuth{
		db: mysql.Instance(),
	}
}

// All 获取所有Http协议接口列表
func (d *apiAuth) All() ([]apiAuthModel.ApiHttp, int64, error) {
	tx := d.db.GetDbR().Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&apiAuthModel.ApiHttp{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]apiAuthModel.ApiHttp, 0)
	if result := tx.Order("updated_at DESC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}

	return bean, total, nil
}

// List 查询Http协议接口列表
func (d *apiAuth) List(req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error) {
	tx := d.db.GetDbR()
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
	if result := tx.Model(&apiAuthModel.ApiHttp{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]apiAuthModel.ApiHttp, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).Order("updated_at DESC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// Add 添加Http协议接口
func (d *apiAuth) Add(bean apiAuthModel.ApiHttp) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新Http协议接口
func (d *apiAuth) Update(bean apiAuthModel.ApiHttp) (int64, error) {
	result := d.db.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除Http协议接口
func (d *apiAuth) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&apiAuthModel.ApiHttp{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除Http协议接口
func (d *apiAuth) BatchDelete(ids []uint) (int64, error) {
	beans := make([]apiAuthModel.ApiHttp, len(ids))
	for _, id := range ids {
		beans = append(beans, apiAuthModel.ApiHttp{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *apiAuth) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&apiAuthModel.ApiHttp{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
