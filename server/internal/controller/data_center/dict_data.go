/*字典数据管理*/
package datacenter

import (
	"github.com/silent-rain/gin-admin/internal/dto"
	dataCenterDTO "github.com/silent-rain/gin-admin/internal/dto/data_center"
	dataCenterModel "github.com/silent-rain/gin-admin/internal/model/data_center"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	dataCenterService "github.com/silent-rain/gin-admin/internal/service/data_center"

	"github.com/gin-gonic/gin"
)

// 字典数据信息
type dictDataCenterController struct {
	service dataCenterService.DictDataService
}

// NewDictDataController 创建字典数据信息控制器对象
func NewDictDataController() *dictDataCenterController {
	return &dictDataCenterController{
		service: dataCenterService.NewDictDataService(),
	}
}

// List 获取所有字典数据信息列表
func (c *dictDataCenterController) List(ctx *gin.Context) {
	req := dataCenterDTO.QueryDictDataReq{}
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
func (c *dictDataCenterController) Add(ctx *gin.Context) {
	req := dataCenterDTO.AddDictDataReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := dataCenterModel.DictData{}
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
func (c *dictDataCenterController) Update(ctx *gin.Context) {
	req := dataCenterDTO.UpdateDictDataReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := dataCenterModel.DictData{}
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
func (c *dictDataCenterController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除字典数据信息
func (c *dictDataCenterController) BatchDelete(ctx *gin.Context) {
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

// Status 更新字典数据信息状态
func (c *dictDataCenterController) Status(ctx *gin.Context) {
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
