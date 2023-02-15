/*菜单*/
package system

import (
	"gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/context"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 菜单
type menuController struct {
	service systemService.MenuService
}

// NewMenuController 创建菜单对象
func NewMenuController() *menuController {
	return &menuController{
		service: systemService.NewMenuService(),
	}
}

// AllTree 获取所有菜单树
func (c *menuController) AllTree(ctx *gin.Context) {
	results, total, err := c.service.AllTree(ctx)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Tree 获取菜单树
func (c *menuController) Tree(ctx *gin.Context) {
	req := systemDTO.QueryMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	results, total, err := c.service.Tree(ctx, req)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Add 添加菜单
func (c *menuController) Add(ctx *gin.Context) {
	req := systemDTO.AddMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	menu := systemModel.Menu{}
	if err := http.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	userId := context.GetUserId(ctx)
	menu.CreateUserId = userId
	menu.UpdateUserId = userId

	_, err := c.service.Add(ctx, menu)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新菜单
func (c *menuController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	menu := systemModel.Menu{}
	if err := http.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	userId := context.GetUserId(ctx)
	menu.UpdateUserId = userId

	_, err := c.service.Update(ctx, menu)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除菜单
func (c *menuController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	_, err := c.service.Delete(ctx, req.ID)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (c *menuController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	_, err := c.service.BatchDelete(ctx, req.Ids)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Status 更新菜单状态
func (c *menuController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	_, err := c.service.Status(ctx, req.ID, req.Status)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}
