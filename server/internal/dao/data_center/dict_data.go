/*字典数据管理*/
package datacenter

import (
	"errors"

	dictDataCenterDTO "gin-admin/internal/dto/data_center"
	dictDataCenterModel "gin-admin/internal/model/data_center"
	"gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// DictData 字典数据信息接口
type DictData interface {
	List(req dictDataCenterDTO.QueryDictDataReq) ([]dictDataCenterModel.DictData, int64, error)
	InfoByValue(dictId uint, value string) (dictDataCenterModel.DictData, bool, error)
	Add(bean dictDataCenterModel.DictData) (uint, error)
	Update(bean dictDataCenterModel.DictData) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
}

// 字典数据信息
type dictDataCenter struct {
	db mysql.DBRepo
}

// NewDictDataDao 创建字典数据信息对象
func NewDictDataDao() *dictDataCenter {
	return &dictDataCenter{
		db: mysql.Instance(),
	}
}

// List 查询字典数据信息列表
func (d *dictDataCenter) List(req dictDataCenterDTO.QueryDictDataReq) ([]dictDataCenterModel.DictData, int64, error) {
	tx := d.db.GetDbR()
	if req.Name != "" {
		tx = tx.Where("name = ?", req.Name)
	}
	if req.Value != "" {
		tx = tx.Where("value = ?", req.Value)
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&dictDataCenterModel.DictData{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]dictDataCenterModel.DictData, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).Order("id ASC").Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// InfoByValue 获取指定字典的字典项数据信息
func (d *dictDataCenter) InfoByValue(dictId uint, value string) (dictDataCenterModel.DictData, bool, error) {
	bean := dictDataCenterModel.DictData{}
	result := d.db.GetDbR().Where("dict_id = ? and value = ?", dictId, value).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加字典数据信息
func (d *dictDataCenter) Add(bean dictDataCenterModel.DictData) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新字典数据信息
func (d *dictDataCenter) Update(bean dictDataCenterModel.DictData) (int64, error) {
	result := d.db.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除字典数据信息
func (d *dictDataCenter) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&dictDataCenterModel.DictData{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除字典数据信息
func (d *dictDataCenter) BatchDelete(ids []uint) (int64, error) {
	beans := make([]dictDataCenterModel.DictData, len(ids))
	for _, id := range ids {
		beans = append(beans, dictDataCenterModel.DictData{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新字典数据信息状态
func (d *dictDataCenter) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&dictDataCenterModel.DictData{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
