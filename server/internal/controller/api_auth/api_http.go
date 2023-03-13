/*Http协议接口管理表*/
package apiauth

import (
	"github.com/silent-rain/gin-admin/internal/dto"
	apiAuthDTO "github.com/silent-rain/gin-admin/internal/dto/api_auth"
	apiAuthModel "github.com/silent-rain/gin-admin/internal/model/api_auth"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	apiAuthService "github.com/silent-rain/gin-admin/internal/service/api_auth"

	"github.com/gin-gonic/gin"
)

// Http协议接口信息
type apiAuthController struct {
	service apiAuthService.ApiHttpService
}

// NewApiHttpController 创建Http协议接口信息对象
func NewApiHttpController() *apiAuthController {
	return &apiAuthController{
		service: apiAuthService.NewApiHttpService(),
	}
}

// AllTree 获取所有Http协议接口信息树
func (c *apiAuthController) AllTree(ctx *gin.Context) {
	results, total, err := c.service.AllTree(ctx)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Tree 获取Http协议接口信息树
func (c *apiAuthController) Tree(ctx *gin.Context) {
	req := apiAuthDTO.QueryApiHttpReq{}
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

// Add 添加Http协议接口信息
func (c *apiAuthController) Add(ctx *gin.Context) {
	req := apiAuthDTO.AddApiHttpReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := apiAuthModel.ApiHttp{}
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

// Update 更新Http协议接口信息
func (c *apiAuthController) Update(ctx *gin.Context) {
	req := apiAuthDTO.UpdateApiHttpReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	bean := apiAuthModel.ApiHttp{}
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

// Delete 删除Http协议接口信息
func (c *apiAuthController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除Http协议接口信息
func (c *apiAuthController) BatchDelete(ctx *gin.Context) {
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

// Status 更新Http协议接口信息状态
func (c *apiAuthController) Status(ctx *gin.Context) {
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
