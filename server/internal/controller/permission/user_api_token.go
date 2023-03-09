/*用户API接口Token令牌表*/
package permission

import (
	"gin-admin/internal/dto"
	permissionDTO "gin-admin/internal/dto/permission"
	permissionModel "gin-admin/internal/model/permission"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	permissionService "gin-admin/internal/service/permission"
	"gin-admin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Token 令牌
type userApiTokenController struct {
	service permissionService.UserApiTokenService
}

// NewUserApiTokenController 创建 Token 令牌对象
func NewUserApiTokenController() *userApiTokenController {
	return &userApiTokenController{
		service: permissionService.NewUserApiTokenService(),
	}
}

// List 获取用 Token 令牌列表
func (c *userApiTokenController) List(ctx *gin.Context) {
	req := permissionDTO.QueryUserApiTokenReq{}
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

// Add 添加 Token 令牌
func (c *userApiTokenController) Add(ctx *gin.Context) {
	req := permissionDTO.AddUserApiTokenReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := permissionModel.UserApiToken{}
	if err := http.ApiJsonConvertJson(ctx, req, &bean); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	// 生成用户API接口Token
	bean.Token = utils.GenerateTUserApiToken()
	_, err := c.service.Add(ctx, bean)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新 Token 令牌
func (c *userApiTokenController) Update(ctx *gin.Context) {
	req := permissionDTO.UpdateUserApiTokenReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	bean := permissionModel.UserApiToken{}
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

// Delete 删除 Token 令牌
func (c *userApiTokenController) Delete(ctx *gin.Context) {
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

// BatchDelete 批量删除 Token 令牌
func (c *userApiTokenController) BatchDelete(ctx *gin.Context) {
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

// Status 更新 Token 令牌状态
func (c *userApiTokenController) Status(ctx *gin.Context) {
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
