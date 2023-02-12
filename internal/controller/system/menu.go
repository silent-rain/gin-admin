/*菜单*/
package system

import (
	"gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/context"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/log"
	service "gin-admin/internal/service/system"

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
	c.service.AllTree(ctx)
}

// Tree 获取菜单树
func (c *menuController) Tree(ctx *gin.Context) {
	req := systemDTO.QueryMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		return
	}

	c.service.Tree(ctx, req)
}

// Add 添加菜单
func (c *menuController) Add(ctx *gin.Context) {
	req := systemDTO.AddMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		return
	}
	menu := systemModel.Menu{}
	if err := http.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		return
	}
	userId := context.GetUserId(ctx)
	menu.CreateUserId = userId
	menu.UpdateUserId = userId

	c.service.Add(ctx, menu)
}

// Update 更新菜单
func (c *menuController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateMenuReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	menu := systemModel.Menu{}
	if err := http.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	userId := context.GetUserId(ctx)
	menu.UpdateUserId = userId

	c.service.Update(ctx, menu)
}

// Delete 删除菜单
func (c *menuController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.Delete(ctx, req.ID)
}

// BatchDelete 批量删除菜单, 批量删除，不校验是否存在子菜单
func (c *menuController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.BatchDelete(ctx, req.Ids)
}

// Status 更新菜单状态
func (c *menuController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.Status(ctx, req.ID, req.Status)
}
