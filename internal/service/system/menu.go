/*菜单*/
package service

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

type MenuService interface {
	AllTree(ctx *gin.Context) *response.ResponseAPI
	Tree(ctx *gin.Context, req systemDTO.QueryMenuReq) *response.ResponseAPI
	Add(ctx *gin.Context, menu systemModel.Menu) *response.ResponseAPI
	Update(ctx *gin.Context, menu systemModel.Menu) *response.ResponseAPI
	Delete(ctx *gin.Context, id uint) *response.ResponseAPI
	BatchDelete(ctx *gin.Context, ids []uint) *response.ResponseAPI
	Status(ctx *gin.Context, id uint, status uint) *response.ResponseAPI
}

// 菜单
type menuService struct {
	dao systemDAO.Menu
}

// NewMenuService 创建菜单对象
func NewMenuService() *menuService {
	return &menuService{
		dao: systemDAO.NewMenuDao(),
	}
}

// AllTree 获取所有菜单树
func (s *menuService) AllTree(ctx *gin.Context) *response.ResponseAPI {
	menus, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	// 菜单列表数据转为树结构
	tree := MenuListToTree(menus, nil)
	return response.New().WithDataList(tree, int64(len(tree)))
}

// Tree 获取菜单树
func (s *menuService) Tree(ctx *gin.Context, req systemDTO.QueryMenuReq) *response.ResponseAPI {
	menuList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	menuAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}

	// 菜单列表数据转为树结构
	tree := MenuListToTree(menuAll, nil)

	// 过滤
	treeFilter := make([]systemModel.Menu, 0)
	for _, itemA := range tree {
		for _, item := range menuList {
			if itemA.Title == item.Title {
				treeFilter = append(treeFilter, itemA)
			}
		}
	}
	return response.New().WithDataList(treeFilter, int64(len(tree)))
}

// Add 添加菜单
func (s *menuService) Add(ctx *gin.Context, menu systemModel.Menu) *response.ResponseAPI {
	if _, err := s.dao.Add(menu); err != nil {
		log.New(ctx).WithCode(statuscode.DBAddError).Errorf("%v", err)

		return response.New().WithCode(statuscode.DBAddError)
	}
	return response.New()
}

// Update 更新菜单
func (s *menuService) Update(ctx *gin.Context, menu systemModel.Menu) *response.ResponseAPI {
	row, err := s.dao.Update(menu)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBUpdateError)
	}
	return response.New().WithData(row)
}

// Delete 删除菜单
func (s *menuService) Delete(ctx *gin.Context, id uint) *response.ResponseAPI {
	childrenMenu, err := s.dao.ChildrenMenu(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	if len(childrenMenu) > 0 {
		log.New(ctx).WithCode(statuscode.DBDataExistChildrenError).Errorf("删除失败, 存在子菜单, %v", err)
		return response.New().WithCode(statuscode.DBDataExistChildrenError).WithMsg("删除失败, 存在子菜单")
	}

	row, err := s.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBDeleteError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBDeleteError)
	}
	return response.New().WithData(row)
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (s *menuService) BatchDelete(ctx *gin.Context, ids []uint) *response.ResponseAPI {
	row, err := s.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBBatchDeleteError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBBatchDeleteError)
	}
	return response.New().WithData(row)
}

// Status 更新菜单状态
func (s *menuService) Status(ctx *gin.Context, id uint, status uint) *response.ResponseAPI {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateStatusError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBUpdateStatusError)
	}
	return response.New().WithData(row)
}

// MenuListToTree 菜单列表数据转为树结构
func MenuListToTree(src []systemModel.Menu, parentId *uint) []systemModel.Menu {
	tree := make([]systemModel.Menu, 0)
	for _, item := range src {
		if (item.ParentId == nil && parentId == nil) ||
			(item.ParentId != nil && parentId != nil && *item.ParentId == *parentId) {
			tree = append(tree, item)
		}
	}

	for i := range tree {
		children := MenuListToTree(src, &tree[i].ID)
		if tree[i].Children == nil {
			tree[i].Children = children
		} else {
			tree[i].Children = append(tree[i].Children, children...)
		}
	}
	return tree
}
