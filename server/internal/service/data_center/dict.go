/*字典维度管理*/
package datacenter

import (
	dictCenterDAO "gin-admin/internal/dao/data_center"
	dictCenterDTO "gin-admin/internal/dto/data_center"
	dictCenterModel "gin-admin/internal/model/data_center"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// DictService 字典维度信息接口
type DictService interface {
	List(ctx *gin.Context, req dictCenterDTO.QueryDictReq) ([]dictCenterModel.Dict, int64, error)
	Add(ctx *gin.Context, bean dictCenterModel.Dict) (uint, error)
	Update(ctx *gin.Context, bean dictCenterModel.Dict) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 字典维度信息
type dictService struct {
	dao dictCenterDAO.Dict
}

// NewDictService 创建字典维度信息服务对象
func NewDictService() *dictService {
	return &dictService{
		dao: dictCenterDAO.NewDictDao(),
	}
}

// List 获取字典维度信息列表
func (s *dictService) List(ctx *gin.Context, req dictCenterDTO.QueryDictReq) ([]dictCenterModel.Dict, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return results, total, nil
}

// Add 添加字典维度信息
func (h *dictService) Add(ctx *gin.Context, bean dictCenterModel.Dict) (uint, error) {
	_, ok, err := h.dao.InfoByCode(bean.Code)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBQueryError)
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("字典维度已存在")
		return 0, errcode.New(errcode.DBDataExistError).WithMsg("字典维度已存在")
	}

	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}

// Update 更新字典维度信息
func (h *dictService) Update(ctx *gin.Context, bean dictCenterModel.Dict) (int64, error) {
	row, err := h.dao.Update(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateError)
	}
	return row, nil
}

// Delete 删除字典维度信息
func (h *dictService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBDeleteError)
	}
	return row, nil
}

// BatchDelete 批量删除字典维度信息
func (h *dictService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBBatchDeleteError)
	}
	return row, nil
}

// Status 更新字典维度信息状态
func (h *dictService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateStatusError)
	}
	return row, nil
}
