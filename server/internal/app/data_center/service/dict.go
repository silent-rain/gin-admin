// Package service 字典维度管理
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/data_center/dao"
	"github.com/silent-rain/gin-admin/internal/app/data_center/dto"
	"github.com/silent-rain/gin-admin/internal/app/data_center/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// DictService 字典维度信息
type DictService struct {
	dao *dao.Dict
}

// NewDictService 创建字典维度信息服务对象
func NewDictService() *DictService {
	return &DictService{
		dao: dao.NewDictDao(),
	}
}

// List 获取字典维度信息列表
func (s *DictService) List(ctx *gin.Context, req dto.QueryDictReq) ([]model.Dict, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}

// Add 添加字典维度信息
func (h *DictService) Add(ctx *gin.Context, bean model.Dict) (uint, error) {
	_, ok, err := h.dao.InfoByCode(bean.Code)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("字典已存在")
		return 0, errcode.DBDataExistError.WithMsg("字典已存在")
	}

	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Update 更新字典维度信息
func (h *DictService) Update(ctx *gin.Context, bean model.Dict) (int64, error) {
	row, err := h.dao.Update(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// Delete 删除字典维度信息
func (h *DictService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}
	return row, nil
}

// BatchDelete 批量删除字典维度信息
func (h *DictService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}
	return row, nil
}

// UpdateStatus 更新字典维度信息状态
func (h *DictService) UpdateStatus(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.UpdateStatus(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}
	return row, nil
}
