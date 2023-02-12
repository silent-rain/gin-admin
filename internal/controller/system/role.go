/*角色
 */
package system

import (
	"gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/log"
	service "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 角色
type roleController struct {
	service service.RoleService
}

// NewRoleController 创建角色对象
func NewRoleController() *roleController {
	return &roleController{
		service: service.NewRoleService(),
	}
}

// All 获取所有角色列表
func (c *roleController) All(ctx *gin.Context) {
	c.service.All(ctx)
}

// List 获取用角色列表
func (c *roleController) List(ctx *gin.Context) {
	req := systemDTO.QueryRoleReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		return
	}

	c.service.List(ctx, req)
}

// Add 添加角色
func (c *roleController) Add(ctx *gin.Context) {
	req := systemDTO.AddRoleReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	role := systemModel.Role{}
	if err := http.ApiJsonConvertJson(ctx, req, &role); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}

	c.service.Add(ctx, role)
}

// Update 更新角色
func (c *roleController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateRoleReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	role := systemModel.Role{}
	if err := http.ApiJsonConvertJson(ctx, req, &role); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}

	c.service.Update(ctx, role)
}

// Delete 删除角色
func (c *roleController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.Delete(ctx, req.ID)
}

// BatchDelete 批量删除角色
func (c *roleController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.BatchDelete(ctx, req.Ids)
}

// Status 更新角色状态
func (c *roleController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.Status(ctx, req.ID, req.Status)
}
