/*字典数据管理*/
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/internal/app/data_center/dto"
	"github.com/silent-rain/gin-admin/internal/app/data_center/model"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// DictData 字典数据信息接口
type DictData interface {
	List(req dto.QueryDictDataReq) ([]model.DictData, int64, error)
	InfoByValue(dictId uint, value string) (model.DictData, bool, error)
	Add(bean model.DictData) (uint, error)
	Update(bean model.DictData) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// 字典数据信息
type dictDataCenter struct {
	mysql.DBRepo
}

// NewDictDataDao 创建字典数据信息对象
func NewDictDataDao() *dictDataCenter {
	return &dictDataCenter{
		DBRepo: mysql.Instance(),
	}
}

// List 查询字典数据信息列表
func (d *dictDataCenter) List(req dto.QueryDictDataReq) ([]model.DictData, int64, error) {
	tx := d.GetDbR()
	if req.DictId != 0 {
		tx = tx.Where("dict_id = ?", req.DictId)
	}
	if req.Name != "" {
		tx = tx.Where("name like ?", req.Name+"%")
	}
	if req.Value != "" {
		tx = tx.Where("value like ?", req.Value+"%")
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&model.DictData{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]model.DictData, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).Order("id ASC").Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// InfoByValue 获取指定字典的字典项数据信息
func (d *dictDataCenter) InfoByValue(dictId uint, value string) (model.DictData, bool, error) {
	bean := model.DictData{}
	result := d.GetDbR().Where("dict_id = ? and value = ?", dictId, value).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加字典数据信息
func (d *dictDataCenter) Add(bean model.DictData) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新字典数据信息
func (d *dictDataCenter) Update(bean model.DictData) (int64, error) {
	result := d.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除字典数据信息
func (d *dictDataCenter) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.DictData{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除字典数据信息
func (d *dictDataCenter) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.DictData, len(ids))
	for _, id := range ids {
		beans = append(beans, model.DictData{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新字典数据信息状态
func (d *dictDataCenter) Status(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.DictData{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
