// Package service Http协议接口管理表
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/api_auth/dao"
	"github.com/silent-rain/gin-admin/internal/app/api_auth/dto"
	"github.com/silent-rain/gin-admin/internal/app/api_auth/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// ApiHttpService 角色与Http协议接口关系
type ApiHttpService struct {
	dao *dao.ApiHttp
}

// NewApiHttpService 创建服务对象
func NewApiHttpService() *ApiHttpService {
	return &ApiHttpService{
		dao: dao.NewApiHttpDao(),
	}
}

// AllTree 获取所有接口树
func (s *ApiHttpService) AllTree(ctx *gin.Context) ([]model.ApiHttp, int64, error) {
	results, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	// 菜单列表数据转为树结构
	tree := apiHttpListToTree(results, nil)
	return tree, int64(len(tree)), nil
}

// Tree 获取接口树
func (s *ApiHttpService) Tree(ctx *gin.Context, req dto.QueryApiHttpReq) ([]model.ApiHttp, int64, error) {
	resultList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	resultAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}

	// 列表数据转为树结构
	tree := apiHttpListToTree(resultAll, nil)

	// 过滤
	treeFilter := make([]model.ApiHttp, 0)
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
func apiHttpListToTree(src []model.ApiHttp, parentId *uint) []model.ApiHttp {
	tree := make([]model.ApiHttp, 0)
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
func (h *ApiHttpService) Add(ctx *gin.Context, bean model.ApiHttp) (uint, error) {
	_, ok, err := h.dao.InfoByUri(bean.Uri)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("接口已存在")
		return 0, errcode.DBDataExistError.WithMsg("接口已存在")
	}

	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Update 更新
func (h *ApiHttpService) Update(ctx *gin.Context, bean model.ApiHttp) (int64, error) {
	row, err := h.dao.Update(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// Delete 删除
func (h *ApiHttpService) Delete(ctx *gin.Context, id uint) (int64, error) {
	childrenConfig, err := h.dao.Children(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if len(childrenConfig) > 0 {
		log.New(ctx).WithCode(errcode.DBDataExistChildrenError).Errorf("删除失败, 存在子接口, %v", err)
		return 0, errcode.DBDataExistChildrenError.WithMsg("删除失败, 存在子接口")
	}

	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}
	return row, nil
}

// BatchDelete 批量删除
func (h *ApiHttpService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}
	return row, nil
}

// UpdateStatus 更新状态
func (h *ApiHttpService) UpdateStatus(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.UpdateStatus(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}
	return row, nil
}
