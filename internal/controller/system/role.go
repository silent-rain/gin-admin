/*角色
 */
package system

import (
	"gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/http"
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
	c.service.All(ctx).Json(ctx)
}

// List 获取用角色列表
func (c *roleController) List(ctx *gin.Context) {
	req := systemDTO.QueryRoleReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.List(ctx, req).Json(ctx)
}

// Add 添加角色
func (c *roleController) Add(ctx *gin.Context) {
	req := systemDTO.AddRoleReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}
	role := systemModel.Role{}
	if result := http.ApiJsonConvertJson(ctx, req, &role); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Add(ctx, role).Json(ctx)
}

// Update 更新角色
func (c *roleController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateRoleReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}
	role := systemModel.Role{}
	if result := http.ApiJsonConvertJson(ctx, req, &role); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Update(ctx, role).Json(ctx)
}

// Delete 删除角色
func (c *roleController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Delete(ctx, req.ID).Json(ctx)
}

// BatchDelete 批量删除角色
func (c *roleController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.BatchDelete(ctx, req.Ids).Json(ctx)
}

// Status 更新角色状态
func (c *roleController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Status(ctx, req.ID, req.Status).Json(ctx)
}
