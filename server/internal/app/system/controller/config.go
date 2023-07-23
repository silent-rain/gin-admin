// Package controller 应用配置表
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/app/system/service"
	DTO "github.com/silent-rain/gin-admin/internal/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// 配置
type configController struct {
	service *service.ConfigService
}

// NewConfigController 创建配置对象
func NewConfigController() *configController {
	return &configController{
		service: service.NewConfigService(),
	}
}

// AllTree 获取所有配置树
func (c *configController) AllTree(ctx *gin.Context) {
	results, total, err := c.service.AllTree(ctx)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Tree 获取配置树
func (c *configController) Tree(ctx *gin.Context) {
	req := dto.QueryConfigReq{}
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

// List 获取配置列表
func (c *configController) List(ctx *gin.Context) {
	req := dto.QueryConfigReq{}
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

// Info 获取配置信息
func (c *configController) Info(ctx *gin.Context) {
	key, ok := ctx.GetQuery("key")
	if !ok {
		log.New(ctx).WithCode(errcode.ReqParameterParsingError).Errorf("")
		response.New(ctx).
			WithError(errcode.ReqParameterParsingError.WithMsg("key 字段不能为空")).Json()
		return
	}

	result, err := c.service.Info(ctx, key)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}

// ChildrensByKey 通过父 key 获取子配置列表
func (c *configController) ChildrensByKey(ctx *gin.Context) {
	req := dto.QueryConfigReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	results, err := c.service.ChildrensByKey(ctx, req.Key)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, int64(len(results))).Json()
}

// Add 添加配置
func (c *configController) Add(ctx *gin.Context) {
	req := dto.AddConfigReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	config := model.Config{}
	if err := http.ApiJsonConvertJson(ctx, req, &config); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Add(ctx, config)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新配置
func (c *configController) Update(ctx *gin.Context) {
	req := dto.UpdateConfigReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	config := model.Config{}
	if err := http.ApiJsonConvertJson(ctx, req, &config); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Update(ctx, config)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// BatchUpdate 批量更新配置
func (c *configController) BatchUpdate(ctx *gin.Context) {
	req := make([]dto.UpdateConfigReq, 0)
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	var configs []model.Config
	if err := http.ApiJsonConvertJson(ctx, req, &configs); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	err := c.service.BatchUpdate(ctx, configs)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除配置
func (c *configController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除配置, 批量删除，不校验是否存在子配置
func (c *configController) BatchDelete(ctx *gin.Context) {
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

// UpdateStatus 更新配置状态
func (c *configController) UpdateStatus(ctx *gin.Context) {
	req := DTO.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if _, err := c.service.UpdateStatus(ctx, req.ID, req.Status); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// WebSiteConfigList 查询网站配置列表
func (c *configController) WebSiteConfigList(ctx *gin.Context) {
	results, err := c.service.WebSiteConfigList(ctx, constant.WebsiteConfigKey)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, int64(len(results))).Json()
}
