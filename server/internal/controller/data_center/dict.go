/*字典维度管理*/
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

// 字典维度信息
type dictCenterController struct {
	service dictCenterService.DictService
}

// NewDictController 创建字典维度信息控制器对象
func NewDictController() *dictCenterController {
	return &dictCenterController{
		service: dictCenterService.NewDictService(),
	}
}

// List 获取所有字典维度信息列表
func (c *dictCenterController) List(ctx *gin.Context) {
	req := dictCenterDTO.QueryDictReq{}
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

// Add 添加字典维度信息
func (c *dictCenterController) Add(ctx *gin.Context) {
	req := dictCenterDTO.AddDictReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := dictCenterModel.Dict{}
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

// Update 更新字典维度信息
func (c *dictCenterController) Update(ctx *gin.Context) {
	req := dictCenterDTO.UpdateDictReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := dictCenterModel.Dict{}
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

// Delete 删除字典维度信息
func (c *dictCenterController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除字典维度信息
func (c *dictCenterController) BatchDelete(ctx *gin.Context) {
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

// Status 更新字典维度信息状态
func (c *dictCenterController) Status(ctx *gin.Context) {
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
