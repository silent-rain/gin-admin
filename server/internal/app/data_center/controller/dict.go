// Package controller 字典维度管理
package controller

import (
	"github.com/silent-rain/gin-admin/internal/app/data_center/dto"
	"github.com/silent-rain/gin-admin/internal/app/data_center/model"
	"github.com/silent-rain/gin-admin/internal/app/data_center/service"
	DTO "github.com/silent-rain/gin-admin/internal/dto"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// 字典维度信息
type dictController struct {
	service *service.DictService
}

// NewDictController 创建字典维度信息控制器对象
func NewDictController() *dictController {
	return &dictController{
		service: service.NewDictService(),
	}
}

// List 获取所有字典维度信息列表
func (c *dictController) List(ctx *gin.Context) {
	req := dto.QueryDictReq{}
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

// Add 添加字典维度信息
func (c *dictController) Add(ctx *gin.Context) {
	req := dto.AddDictReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := model.Dict{}
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

// Update 更新字典维度信息
func (c *dictController) Update(ctx *gin.Context) {
	req := dto.UpdateDictReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := model.Dict{}
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

// Delete 删除字典维度信息
func (c *dictController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除字典维度信息
func (c *dictController) BatchDelete(ctx *gin.Context) {
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

// UpdateStatus 更新字典维度信息状态
func (c *dictController) UpdateStatus(ctx *gin.Context) {
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
