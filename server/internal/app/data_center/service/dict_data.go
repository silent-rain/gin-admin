// Package service 字典数据管理
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/data_center/dao"
	"github.com/silent-rain/gin-admin/internal/app/data_center/dto"
	"github.com/silent-rain/gin-admin/internal/app/data_center/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// DictDataService 字典数据信息
type DictDataService struct {
	dao *dao.DictData
}

// NewDictDataService 创建字典数据信息服务对象
func NewDictDataService() *DictDataService {
	return &DictDataService{
		dao: dao.NewDictDataDao(),
	}
}

// List 获取字典数据信息列表
func (s *DictDataService) List(ctx *gin.Context, req dto.QueryDictDataReq) ([]model.DictData, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}

// Add 添加字典数据信息
func (h *DictDataService) Add(ctx *gin.Context, bean model.DictData) (uint, error) {
	_, ok, err := h.dao.InfoByValue(bean.DictId, bean.Value)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("典数据已存在")
		return 0, errcode.DBDataExistError.WithMsg("典数据已存在")
	}

	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Update 更新字典数据信息
func (h *DictDataService) Update(ctx *gin.Context, bean model.DictData) (int64, error) {
	row, err := h.dao.Update(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// Delete 删除字典数据信息
func (h *DictDataService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}
	return row, nil
}

// BatchDelete 批量删除字典数据信息
func (h *DictDataService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}
	return row, nil
}

// UpdateStatus 更新字典数据信息状态
func (h *DictDataService) UpdateStatus(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.UpdateStatus(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}
	return row, nil
}
