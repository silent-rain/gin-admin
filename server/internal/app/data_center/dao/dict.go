// Package dao 字典维度管理
package dao

import (
	"errors"

	"github.com/silent-rain/gin-admin/global"
	"github.com/silent-rain/gin-admin/internal/app/data_center/dto"
	"github.com/silent-rain/gin-admin/internal/app/data_center/model"
	"github.com/silent-rain/gin-admin/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Dict 字典维度信息
type Dict struct {
	mysql.DBRepo
}

// NewDictDao 创建字典维度信息对象
func NewDictDao() *Dict {
	return &Dict{
		DBRepo: global.Instance().Mysql(),
	}
}

// List 查询字典维度信息列表
func (d *Dict) List(req dto.QueryDictReq) ([]model.Dict, int64, error) {
	tx := d.GetDbR()
	if req.Name != "" {
		tx = tx.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Code != "" {
		tx = tx.Where("code like ?", "%"+req.Code+"%")
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&model.Dict{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]model.Dict, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).Order("id ASC").Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// InfoByCode 获取字典维度信息
func (d *Dict) InfoByCode(code string) (model.Dict, bool, error) {
	bean := model.Dict{}
	result := d.GetDbR().Where("code = ?", code).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加字典维度信息
func (d *Dict) Add(bean model.Dict) (uint, error) {
	result := d.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新字典维度信息
func (d *Dict) Update(bean model.Dict) (int64, error) {
	result := d.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除字典维度信息
func (d *Dict) Delete(id uint) (int64, error) {
	result := d.GetDbW().Delete(&model.Dict{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除字典维度信息
func (d *Dict) BatchDelete(ids []uint) (int64, error) {
	beans := make([]model.Dict, len(ids))
	for _, id := range ids {
		beans = append(beans, model.Dict{
			ID: id,
		})
	}
	result := d.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// UpdateStatus 更新字典维度信息状态
func (d *Dict) UpdateStatus(id uint, status uint) (int64, error) {
	result := d.GetDbW().Select("status").Updates(&model.Dict{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}
