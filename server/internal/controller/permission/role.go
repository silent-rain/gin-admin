/*角色
 */
package permission

import (
	"github.com/silent-rain/gin-admin/internal/dto"
	permissionDTO "github.com/silent-rain/gin-admin/internal/dto/permission"
	permissionModel "github.com/silent-rain/gin-admin/internal/model/permission"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	permissionService "github.com/silent-rain/gin-admin/internal/service/permission"
	"github.com/silent-rain/gin-admin/pkg/tracer"

	"github.com/gin-gonic/gin"
)

// 角色
type roleController struct {
	service permissionService.RoleService
}

// NewRoleController 创建角色对象
func NewRoleController() *roleController {
	return &roleController{
		service: permissionService.NewRoleService(),
	}
}

// All 获取所有角色列表
func (c *roleController) All(ctx *gin.Context) {
	results, total, err := c.service.All(ctx)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// List 获取用角色列表
func (c *roleController) List(ctx *gin.Context) {
	span := tracer.SpanStart(ctx)
	defer span.Finish()

	req := permissionDTO.QueryRoleReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Add 添加角色
func (c *roleController) Add(ctx *gin.Context) {
	req := permissionDTO.AddRoleReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	role := permissionModel.Role{}
	if err := http.ApiJsonConvertJson(ctx, req, &role); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Add(ctx, role)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新角色
func (c *roleController) Update(ctx *gin.Context) {
	req := permissionDTO.UpdateRoleReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	role := permissionModel.Role{}
	if err := http.ApiJsonConvertJson(ctx, req, &role); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Update(ctx, role)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除角色
func (c *roleController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除角色
func (c *roleController) BatchDelete(ctx *gin.Context) {
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

// Status 更新角色状态
func (c *roleController) Status(ctx *gin.Context) {
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
