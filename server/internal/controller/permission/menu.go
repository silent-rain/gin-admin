/*菜单*/
package permission

import (
	"github.com/silent-rain/gin-admin/internal/dto"
	permissionDTO "github.com/silent-rain/gin-admin/internal/dto/permission"
	permissionModel "github.com/silent-rain/gin-admin/internal/model/permission"
	"github.com/silent-rain/gin-admin/internal/pkg/core"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	permissionService "github.com/silent-rain/gin-admin/internal/service/permission"

	"github.com/gin-gonic/gin"
)

// 菜单
type menuController struct {
	service permissionService.MenuService
}

// NewMenuController 创建菜单对象
func NewMenuController() *menuController {
	return &menuController{
		service: permissionService.NewMenuService(),
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
	req := permissionDTO.QueryMenuReq{}
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
	req := permissionDTO.AddMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	menu := permissionModel.Menu{}
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

	userId := core.GetContext(ctx).UserId
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
func (c *menuController) setMenuTypeByButtonParams(menu permissionModel.Menu) permissionModel.Menu {
	if menu.MenuType != uint(permissionModel.MenuTypeByButton) {
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
func (c *menuController) setMenuOpenTypeByComponentParams(menu permissionModel.Menu) permissionModel.Menu {
	if !(menu.MenuType == uint(permissionModel.MenuTypeByMenu) &&
		menu.OpenType == uint(permissionModel.MenuOpenTypeByComponent)) {
		return menu
	}
	menu.Permission = ""
	return menu
}

// 设置菜单打开类型为外链接的参数
func (c *menuController) setMenuOpenTypeByOutLinkParams(menu permissionModel.Menu) permissionModel.Menu {
	if !(menu.MenuType == uint(permissionModel.MenuTypeByMenu) &&
		menu.OpenType == uint(permissionModel.MenuOpenTypeByOutLink)) {
		return menu
	}
	menu.Name = ""
	menu.Permission = ""
	menu.Component = ""
	menu.Redirect = ""
	return menu
}

// 设置菜单打开类型为内链接的参数
func (c *menuController) setMenuOpenTypeByInnerLinkParams(menu permissionModel.Menu) permissionModel.Menu {
	if !(menu.MenuType == uint(permissionModel.MenuTypeByMenu) &&
		menu.OpenType == uint(permissionModel.MenuOpenTypeByInnerLink)) {
		return menu
	}
	menu.Permission = ""
	menu.Redirect = ""
	return menu
}

// Update 更新菜单
func (c *menuController) Update(ctx *gin.Context) {
	req := permissionDTO.UpdateMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	menu := permissionModel.Menu{}
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

	userId := core.GetContext(ctx).UserId
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
	req := dto.DeleteReq{}
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
	req := dto.BatchDeleteReq{}
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
	req := dto.UpdateStatusReq{}
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
	req := permissionDTO.QueryChildrenMenuReq{}
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
