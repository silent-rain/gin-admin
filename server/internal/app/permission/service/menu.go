// Package service 菜单
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dao"
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// MenuService 菜单
type MenuService struct {
	dao *dao.Menu
}

// NewMenuService 创建菜单对象
func NewMenuService() *MenuService {
	return &MenuService{
		dao: dao.NewMenuDao(),
	}
}

// AllTree 获取所有菜单树
func (s *MenuService) AllTree(ctx *gin.Context) ([]model.Menu, int64, error) {
	menus, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	// 菜单列表数据转为树结构
	tree := menuListToTree(menus, nil)
	return tree, int64(len(tree)), nil
}

// Tree 获取菜单树
func (s *MenuService) Tree(ctx *gin.Context, req dto.QueryMenuReq) ([]model.Menu, int64, error) {
	menuList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	menuAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}

	// 菜单列表数据转为树结构
	tree := menuListToTree(menuAll, nil)

	// 过滤
	treeFilter := make([]model.Menu, 0)
	for _, itemA := range tree {
		for _, item := range menuList {
			if itemA.Title == item.Title {
				treeFilter = append(treeFilter, itemA)
			}
		}
	}
	return treeFilter, int64(len(tree)), nil
}

// Add 添加菜单
func (s *MenuService) Add(ctx *gin.Context, menu model.Menu) (uint, error) {
	id, err := s.dao.Add(menu)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}

// Update 更新菜单
func (s *MenuService) Update(ctx *gin.Context, menu model.Menu) (int64, error) {
	row, err := s.dao.Update(menu)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// Delete 删除菜单
func (s *MenuService) Delete(ctx *gin.Context, id uint) (int64, error) {
	childrenMenu, err := s.dao.ChildrenMenus(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if len(childrenMenu) > 0 {
		log.New(ctx).WithCode(errcode.DBDataExistChildrenError).Errorf("删除失败, 存在子菜单, %v", err)
		return 0, errcode.DBDataExistChildrenError.WithMsg("删除失败, 存在子菜单")
	}

	row, err := s.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}
	return row, nil
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (s *MenuService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := s.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}
	return row, nil
}

// UpdateStatus 更新菜单状态
func (s *MenuService) UpdateStatus(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := s.dao.UpdateStatus(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}
	return row, nil
}

// 菜单列表数据转为树结构
func menuListToTree(src []model.Menu, parentId *uint) []model.Menu {
	tree := make([]model.Menu, 0)
	for _, item := range src {
		if (item.ParentId == nil && parentId == nil) ||
			(item.ParentId != nil && parentId != nil && *item.ParentId == *parentId) {
			tree = append(tree, item)
		}
	}

	for i := range tree {
		children := menuListToTree(src, &tree[i].ID)
		if tree[i].Children == nil {
			tree[i].Children = children
		} else {
			tree[i].Children = append(tree[i].Children, children...)
		}
	}
	return tree
}

// ChildrenMenus 通过父 ID 获取子配置列表
func (s *MenuService) ChildrenMenus(ctx *gin.Context, parentId uint) ([]model.Menu, error) {
	results, err := s.dao.ChildrenMenus(parentId)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, errcode.DBQueryError
	}
	return results, nil
}
