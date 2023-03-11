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
	AllTree(ctx *gin.Context) ([]apiAuthModel.ApiHttp, int64, error)
	Tree(ctx *gin.Context, req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error)
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

// AllTree 获取所有接口树
func (s *apiHttpService) AllTree(ctx *gin.Context) ([]apiAuthModel.ApiHttp, int64, error) {
	results, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	// 菜单列表数据转为树结构
	tree := apiHttpListToTree(results, nil)
	return tree, int64(len(tree)), nil
}

// Tree 获取接口树
func (s *apiHttpService) Tree(ctx *gin.Context, req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error) {
	resultList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}
	resultAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}

	// 列表数据转为树结构
	tree := apiHttpListToTree(resultAll, nil)

	// 过滤
	treeFilter := make([]apiAuthModel.ApiHttp, 0)
	for _, itemA := range tree {
		for _, item := range resultList {
			if itemA.Name == item.Name {
				treeFilter = append(treeFilter, itemA)
			}
		}
	}
	return treeFilter, int64(len(tree)), nil
}

// 列表数据转为树结构
func apiHttpListToTree(src []apiAuthModel.ApiHttp, parentId *uint) []apiAuthModel.ApiHttp {
	tree := make([]apiAuthModel.ApiHttp, 0)
	for _, item := range src {
		if (item.ParentId == nil && parentId == nil) ||
			(item.ParentId != nil && parentId != nil && *item.ParentId == *parentId) {
			tree = append(tree, item)
		}
	}

	for i := range tree {
		children := apiHttpListToTree(src, &tree[i].ID)
		if tree[i].Children == nil {
			tree[i].Children = children
		} else {
			tree[i].Children = append(tree[i].Children, children...)
		}
	}
	return tree
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
