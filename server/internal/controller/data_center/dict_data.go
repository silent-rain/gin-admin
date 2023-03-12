/*字典数据管理*/
package datacenter

import (
	"gin-admin/internal/dto"
	dictCenterDTO "gin-admin/internal/dto/data_center"
	dictCenterModel "gin-admin/internal/model/data_center"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	dictCenterService "gin-admin/internal/service/data_center"

	"github.com/gin-gonic/gin"
)

// 字典数据信息
type dictDataCenterController struct {
	service dictCenterService.DictDataService
}

// NewDictDataController 创建字典数据信息控制器对象
func NewDictDataController() *dictDataCenterController {
	return &dictDataCenterController{
		service: dictCenterService.NewDictDataService(),
	}
}

// List 获取所有字典数据信息列表
func (c *dictDataCenterController) List(ctx *gin.Context) {
	req := dictCenterDTO.QueryDictDataReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Add 添加字典数据信息
func (c *dictDataCenterController) Add(ctx *gin.Context) {
	req := dictCenterDTO.AddDictDataReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := dictCenterModel.DictData{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	_, err := c.service.Add(ctx, bean)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新字典数据信息
func (c *dictDataCenterController) Update(ctx *gin.Context) {
	req := dictCenterDTO.UpdateDictDataReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := dictCenterModel.DictData{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	_, err := c.service.Update(ctx, bean)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除字典数据信息
func (c *dictDataCenterController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除字典数据信息
func (c *dictDataCenterController) BatchDelete(ctx *gin.Context) {
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

// Status 更新字典数据信息状态
func (c *dictDataCenterController) Status(ctx *gin.Context) {
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