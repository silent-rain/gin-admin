/*字典数据管理*/
package datacenter

import (
	dictCenterDAO "gin-admin/internal/dao/data_center"
	dictCenterDTO "gin-admin/internal/dto/data_center"
	dictCenterModel "gin-admin/internal/model/data_center"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// DictDataService 字典数据信息接口
type DictDataService interface {
	List(ctx *gin.Context, req dictCenterDTO.QueryDictDataReq) ([]dictCenterModel.DictData, int64, error)
	Add(ctx *gin.Context, bean dictCenterModel.DictData) (uint, error)
	Update(ctx *gin.Context, bean dictCenterModel.DictData) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 字典数据信息
type dictDataService struct {
	dao dictCenterDAO.DictData
}

// NewDictDataService 创建字典数据信息服务对象
func NewDictDataService() *dictDataService {
	return &dictDataService{
		dao: dictCenterDAO.NewDictDataDao(),
	}
}

// List 获取字典数据信息列表
func (s *dictDataService) List(ctx *gin.Context, req dictCenterDTO.QueryDictDataReq) ([]dictCenterModel.DictData, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return results, total, nil
}

// Add 添加字典数据信息
func (h *dictDataService) Add(ctx *gin.Context, bean dictCenterModel.DictData) (uint, error) {
	_, ok, err := h.dao.InfoByValue(bean.Value)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBQueryError)
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("典数据已存在")
		return 0, errcode.New(errcode.DBDataExistError).WithMsg("典数据已存在")
	}

	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}

// Update 更新字典数据信息
func (h *dictDataService) Update(ctx *gin.Context, bean dictCenterModel.DictData) (int64, error) {
	row, err := h.dao.Update(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateError)
	}
	return row, nil
}

// Delete 删除字典数据信息
func (h *dictDataService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBDeleteError)
	}
	return row, nil
}

// BatchDelete 批量删除字典数据信息
func (h *dictDataService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBBatchDeleteError)
	}
	return row, nil
}

// Status 更新字典数据信息状态
func (h *dictDataService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateStatusError)
	}
	return row, nil
}
