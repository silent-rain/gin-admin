/*菜单*/
package system

import (
	systemDao "gin-admin/internal/dao/system"
	"gin-admin/internal/dto"
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 菜单
type menuHandler struct {
	dao systemDao.Menu
}

// 创建菜单 Handler 对象
func NewMenuHandler() *menuHandler {
	return &menuHandler{
		dao: systemDao.NewMenuDao(),
	}
}

// AllTree 获取所有菜单树
func (h *menuHandler) AllTree(ctx *gin.Context) {
	menus, _, err := h.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	// 菜单列表数据转为树结构
	tree := MenuListToTree(menus, nil)
	response.New(ctx).WithDataList(tree, int64(len(tree))).Json()
}

// Tree 获取菜单树
func (h *menuHandler) Tree(ctx *gin.Context) {
	req := systemDto.QueryMenuReq{}
	if err := utils.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	menuList, _, err := h.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	menuAll, _, err := h.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
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

// Add 添加菜单
func (h *menuHandler) Add(ctx *gin.Context) {
	req := new(systemDto.AddMenuReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	menu := systemModel.Menu{}
	if err := utils.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	userId := utils.GetUserId(ctx)
	menu.CreateUserId = userId
	menu.UpdateUserId = userId
	if _, err := h.dao.Add(menu); err != nil {
		log.New(ctx).WithCode(statuscode.DbAddError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbAddError).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新菜单
func (h *menuHandler) Update(ctx *gin.Context) {
	req := new(systemDto.UpdateMenuReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	menu := systemModel.Menu{}
	if err := utils.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	userId := utils.GetUserId(ctx)
	menu.UpdateUserId = userId
	row, err := h.dao.Update(menu)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Delete 删除菜单
func (h *menuHandler) Delete(ctx *gin.Context) {
	req := new(dto.DeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	childrenMenu, err := h.dao.ChildrenMenu(req.ID)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	if len(childrenMenu) > 0 {
		log.New(ctx).WithCode(statuscode.DbDataExistChildrenError).Errorf("删除失败, 存在子菜单, %v", err)
		response.New(ctx).WithCode(statuscode.DbDataExistChildrenError).WithMsg("删除失败, 存在子菜单").Json()
		return
	}

	row, err := h.dao.Delete(req.ID)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (h *menuHandler) BatchDelete(ctx *gin.Context) {
	req := new(dto.BatchDeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.BatchDelete(req.Ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbBatchDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbBatchDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新菜单状态
func (h *menuHandler) Status(ctx *gin.Context) {
	req := new(dto.UpdateStatusReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.Status(req.ID, req.Status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}
