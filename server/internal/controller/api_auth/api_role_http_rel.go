/*角色与Http协议接口关联表*/
package apiauth

import (
	apiAuthDTO "github.com/silent-rain/gin-admin/internal/dto/api_auth"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	apiAuthService "github.com/silent-rain/gin-admin/internal/service/api_auth"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// 角色与Http协议接口关系
type apiRoleHttpRelController struct {
	service apiAuthService.ApiRoleHttpRelService
}

// NewApiRoleHttpRelController 创建角色与Http协议接口关系对象
func NewApiRoleHttpRelController() *apiRoleHttpRelController {
	return &apiRoleHttpRelController{
		service: apiAuthService.NewApiRoleHttpRelService(),
	}
}

// List 获取角色与Http协议接口关系列表
func (c *apiRoleHttpRelController) List(ctx *gin.Context) {
	req := apiAuthDTO.QueryApiRoleHttpRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	if req.RoleId == 0 && req.ApiId == 0 {
		log.New(ctx).WithField("data", req).Errorf("role_id/api_id 不能同时为空")
		response.New(ctx).WithCode(errcode.ReqParameterParsingError).WithMsg("role_id/api_id 不能同时为空")
		return
	}

	results, total, err := c.service.List(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Update 更新角色与Http协议接口关系
func (c *apiRoleHttpRelController) Update(ctx *gin.Context) {
	req := apiAuthDTO.UpdateApiRoleHttpRelReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if err := c.service.Update(ctx, req.RoleId, req.ApiIds); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
