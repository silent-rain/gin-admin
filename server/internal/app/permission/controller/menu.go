// Package controller 菜单
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/app/permission/service"
	DTO "github.com/silent-rain/gin-admin/internal/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// 菜单
type menuController struct {
	service service.MenuService
}

// NewMenuController 创建菜单对象
func NewMenuController() *menuController {
	return &menuController{
		service: service.NewMenuService(),
	}
}

// AllTree 获取所有菜单树
func (c *menuController) AllTree(ctx *gin.Context) {
	results, total, err := c.service.AllTree(ctx)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Tree 获取菜单树
func (c *menuController) Tree(ctx *gin.Context) {
	req := dto.QueryMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	results, total, err := c.service.Tree(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Add 添加菜单
func (c *menuController) Add(ctx *gin.Context) {
	req := dto.AddMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	menu := model.Menu{}
	if err := http.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	// 设置菜单类型为按钮的参数
	menu = c.setMenuTypeByButtonParams(menu)
	// 设置菜单打开类型为组件的参数
	menu = c.setMenuOpenTypeByComponentParams(menu)
	// 设置菜单打开类型为外链接的参数
	menu = c.setMenuOpenTypeByOutLinkParams(menu)
	// 设置菜单打开类型为内链接的参数
	menu = c.setMenuOpenTypeByInnerLinkParams(menu)

	userId := core.Context(ctx).UserId
	menu.CreateUserId = userId
	menu.UpdateUserId = userId

	_, err := c.service.Add(ctx, menu)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// 设置菜单类型为按钮的参数
func (c *menuController) setMenuTypeByButtonParams(menu model.Menu) model.Menu {
	if menu.MenuType != uint(model.MenuTypeByButton) {
		return menu
	}
	menu.ElSvgIcon = ""
	menu.Icon = ""
	menu.Name = ""
	menu.Path = ""
	menu.Component = ""
	menu.Redirect = ""
	menu.Link = ""
	return menu
}

// 设置菜单打开类型为组件的参数
func (c *menuController) setMenuOpenTypeByComponentParams(menu model.Menu) model.Menu {
	if !(menu.MenuType == uint(model.MenuTypeByMenu) &&
		menu.OpenType == uint(model.MenuOpenTypeByComponent)) {
		return menu
	}
	menu.Permission = ""
	return menu
}

// 设置菜单打开类型为外链接的参数
func (c *menuController) setMenuOpenTypeByOutLinkParams(menu model.Menu) model.Menu {
	if !(menu.MenuType == uint(model.MenuTypeByMenu) &&
		menu.OpenType == uint(model.MenuOpenTypeByOutLink)) {
		return menu
	}
	menu.Name = ""
	menu.Permission = ""
	menu.Component = ""
	menu.Redirect = ""
	return menu
}

// 设置菜单打开类型为内链接的参数
func (c *menuController) setMenuOpenTypeByInnerLinkParams(menu model.Menu) model.Menu {
	if !(menu.MenuType == uint(model.MenuTypeByMenu) &&
		menu.OpenType == uint(model.MenuOpenTypeByInnerLink)) {
		return menu
	}
	menu.Permission = ""
	menu.Redirect = ""
	return menu
}

// Update 更新菜单
func (c *menuController) Update(ctx *gin.Context) {
	req := dto.UpdateMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	menu := model.Menu{}
	if err := http.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	// 设置菜单类型为按钮的参数
	menu = c.setMenuTypeByButtonParams(menu)
	// 设置菜单打开类型为组件的参数
	menu = c.setMenuOpenTypeByComponentParams(menu)
	// 设置菜单打开类型为外链接的参数
	menu = c.setMenuOpenTypeByOutLinkParams(menu)
	// 设置菜单打开类型为内链接的参数
	menu = c.setMenuOpenTypeByInnerLinkParams(menu)

	userId := core.Context(ctx).UserId
	menu.UpdateUserId = userId

	_, err := c.service.Update(ctx, menu)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除菜单
func (c *menuController) Delete(ctx *gin.Context) {
	req := DTO.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Delete(ctx, req.ID)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (c *menuController) BatchDelete(ctx *gin.Context) {
	req := DTO.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.BatchDelete(ctx, req.Ids)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Status 更新菜单状态
func (c *menuController) Status(ctx *gin.Context) {
	req := DTO.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Status(ctx, req.ID, req.Status)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// ChildrenMenu 通过父 ID 获取子配置列表
func (c *menuController) ChildrenMenu(ctx *gin.Context) {
	req := dto.QueryChildrenMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	results, err := c.service.ChildrenMenu(ctx, req.ParentId)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, int64(len(results))).Json()
}
