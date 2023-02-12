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
	AllTree(ctx *gin.Context)
	Tree(ctx *gin.Context, req systemDTO.QueryMenuReq)
	Add(ctx *gin.Context, menu systemModel.Menu)
	Update(ctx *gin.Context, menu systemModel.Menu)
	Delete(ctx *gin.Context, id uint)
	BatchDelete(ctx *gin.Context, ids []uint)
	Status(ctx *gin.Context, id uint, status uint)
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
func (s *menuService) AllTree(ctx *gin.Context) {
	menus, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	// 菜单列表数据转为树结构
	tree := MenuListToTree(menus, nil)
	response.New(ctx).WithDataList(tree, int64(len(tree))).Json()
}

// Tree 获取菜单树
func (s *menuService) Tree(ctx *gin.Context, req systemDTO.QueryMenuReq) {
	menuList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	menuAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
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
	response.New(ctx).WithDataList(treeFilter, int64(len(tree))).Json()
}

// Add 添加菜单
func (s *menuService) Add(ctx *gin.Context, menu systemModel.Menu) {
	if _, err := s.dao.Add(menu); err != nil {
		log.New(ctx).WithCode(statuscode.DBAddError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBAddError).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新菜单
func (s *menuService) Update(ctx *gin.Context, menu systemModel.Menu) {
	row, err := s.dao.Update(menu)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Delete 删除菜单
func (s *menuService) Delete(ctx *gin.Context, id uint) {
	childrenMenu, err := s.dao.ChildrenMenu(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	if len(childrenMenu) > 0 {
		log.New(ctx).WithCode(statuscode.DBDataExistChildrenError).Errorf("删除失败, 存在子菜单, %v", err)
		response.New(ctx).WithCode(statuscode.DBDataExistChildrenError).WithMsg("删除失败, 存在子菜单").Json()
		return
	}

	row, err := s.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (s *menuService) BatchDelete(ctx *gin.Context, ids []uint) {
	row, err := s.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBBatchDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBBatchDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新菜单状态
func (s *menuService) Status(ctx *gin.Context, id uint, status uint) {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
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
