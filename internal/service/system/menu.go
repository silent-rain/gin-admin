/*菜单*/
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// MenuService 菜单接口
type MenuService interface {
	AllTree(ctx *gin.Context) ([]systemModel.Menu, int64, error)
	Tree(ctx *gin.Context, req systemDTO.QueryMenuReq) ([]systemModel.Menu, int64, error)
	Add(ctx *gin.Context, menu systemModel.Menu) (uint, error)
	Update(ctx *gin.Context, menu systemModel.Menu) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
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
func (s *menuService) AllTree(ctx *gin.Context) ([]systemModel.Menu, int64, error) {
	menus, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	// 菜单列表数据转为树结构
	tree := menuListToTree(menus, nil)
	return tree, int64(len(tree)), nil
}

// Tree 获取菜单树
func (s *menuService) Tree(ctx *gin.Context, req systemDTO.QueryMenuReq) ([]systemModel.Menu, int64, error) {
	menuList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}
	menuAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}

	// 菜单列表数据转为树结构
	tree := menuListToTree(menuAll, nil)

	// 过滤
	treeFilter := make([]systemModel.Menu, 0)
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
func (s *menuService) Add(ctx *gin.Context, menu systemModel.Menu) (uint, error) {
	id, err := s.dao.Add(menu)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}

// Update 更新菜单
func (s *menuService) Update(ctx *gin.Context, menu systemModel.Menu) (int64, error) {
	row, err := s.dao.Update(menu)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateError)
	}
	return row, nil
}

// Delete 删除菜单
func (s *menuService) Delete(ctx *gin.Context, id uint) (int64, error) {
	childrenMenu, err := s.dao.ChildrenMenu(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBQueryError)
	}
	if len(childrenMenu) > 0 {
		log.New(ctx).WithCode(errcode.DBDataExistChildrenError).Errorf("删除失败, 存在子菜单, %v", err)
		return 0, errcode.New(errcode.DBDataExistChildrenError).WithMsg("删除失败, 存在子菜单")
	}

	row, err := s.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBDeleteError)
	}
	return row, nil
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (s *menuService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := s.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBBatchDeleteError)
	}
	return row, nil
}

// Status 更新菜单状态
func (s *menuService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateStatusError)
	}
	return row, nil
}

// 菜单列表数据转为树结构
func menuListToTree(src []systemModel.Menu, parentId *uint) []systemModel.Menu {
	tree := make([]systemModel.Menu, 0)
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
