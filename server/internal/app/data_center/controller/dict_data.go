// Package controller 字典数据管理
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/data_center/dto"
	"github.com/silent-rain/gin-admin/internal/app/data_center/model"
	"github.com/silent-rain/gin-admin/internal/app/data_center/service"
	DTO "github.com/silent-rain/gin-admin/internal/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 字典数据信息
type dictDataController struct {
	service *service.DictDataService
}

// NewDictDataController 创建字典数据信息控制器对象
func NewDictDataController() *dictDataController {
	return &dictDataController{
		service: service.NewDictDataService(),
	}
}

// List 获取所有字典数据信息列表
func (c *dictDataController) List(ctx *gin.Context) {
	req := dto.QueryDictDataReq{}
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

// Add 添加字典数据信息
func (c *dictDataController) Add(ctx *gin.Context) {
	req := dto.AddDictDataReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := model.DictData{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Add(ctx, bean)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新字典数据信息
func (c *dictDataController) Update(ctx *gin.Context) {
	req := dto.UpdateDictDataReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := model.DictData{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.Update(ctx, bean)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除字典数据信息
func (c *dictDataController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除字典数据信息
func (c *dictDataController) BatchDelete(ctx *gin.Context) {
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

// UpdateStatus 更新字典数据信息状态
func (c *dictDataController) UpdateStatus(ctx *gin.Context) {
	req := DTO.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	_, err := c.service.UpdateStatus(ctx, req.ID, req.Status)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
