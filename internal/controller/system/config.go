/*应用配置表*/
package system

import (
	"gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/http"
	service "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 配置
type configController struct {
	service service.ConfigService
}

// NewConfigController 创建配置对象
func NewConfigController() *configController {
	return &configController{
		service: service.NewConfigService(),
	}
}

// AllTree 获取所有配置树
func (c *configController) AllTree(ctx *gin.Context) {
	c.service.AllTree(ctx).Json(ctx)
}

// Tree 获取配置树
func (c *configController) Tree(ctx *gin.Context) {
	req := systemDTO.QueryConfigReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Tree(ctx, req).Json(ctx)
}

// Add 添加配置
func (c *configController) Add(ctx *gin.Context) {
	req := systemDTO.AddConfigReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}
	config := systemModel.Config{}
	if result := http.ApiJsonConvertJson(ctx, req, &config); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Add(ctx, config).Json(ctx)
}

// Update 更新配置
func (c *configController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateConfigReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}
	config := systemModel.Config{}
	if result := http.ApiJsonConvertJson(ctx, req, &config); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Update(ctx, config).Json(ctx)
}

// Delete 删除配置
func (c *configController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Delete(ctx, req.ID).Json(ctx)
}

// BatchDelete 批量删除配置, 批量删除，不校验是否存在子配置
func (c *configController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.BatchDelete(ctx, req.Ids).Json(ctx)
}

// Status 更新配置状态
func (c *configController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Status(ctx, req.ID, req.Status).Json(ctx)
}
