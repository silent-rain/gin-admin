/*Http协议接口管理表*/
package apiauth

import (
	apiAuthDAO "gin-admin/internal/dao/api_auth"
	apiAuthDTO "gin-admin/internal/dto/api_auth"
	apiAuthModel "gin-admin/internal/model/api_auth"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// ApiHttpService 角色与Http协议接口关系接口
type ApiHttpService interface {
	All(ctx *gin.Context) ([]apiAuthModel.ApiHttp, int64, error)
	List(ctx *gin.Context, req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error)
	Add(ctx *gin.Context, bean apiAuthModel.ApiHttp) (uint, error)
	Update(ctx *gin.Context, bean apiAuthModel.ApiHttp) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 角色与Http协议接口关系
type apiHttpService struct {
	dao apiAuthDAO.ApiHttp
}

// NewApiHttpService 创建服务对象
func NewApiHttpService() *apiHttpService {
	return &apiHttpService{
		dao: apiAuthDAO.NewApiHttpDao(),
	}
}

// All 获取所有列表
func (s *apiHttpService) All(ctx *gin.Context) ([]apiAuthModel.ApiHttp, int64, error) {
	results, total, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return results, total, nil
}

// List 获取列表
func (s *apiHttpService) List(ctx *gin.Context, req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return results, total, nil
}

// Add 添加
func (h *apiHttpService) Add(ctx *gin.Context, bean apiAuthModel.ApiHttp) (uint, error) {
	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}

// Update 更新
func (h *apiHttpService) Update(ctx *gin.Context, bean apiAuthModel.ApiHttp) (int64, error) {
	row, err := h.dao.Update(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateError)
	}
	return row, nil
}

// Delete 删除
func (h *apiHttpService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBDeleteError)
	}
	return row, nil
}

// BatchDelete 批量删除
func (h *apiHttpService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBBatchDeleteError)
	}
	return row, nil
}

// Status 更新状态
func (h *apiHttpService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateStatusError)
	}
	return row, nil
}
