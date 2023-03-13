/*用户登录表*/
package system

import (
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	systemService "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 用户登录信息
type userLoginController struct {
	service systemService.UserLoginService
}

// NewUserLoginController 创建用户登录信息对象
func NewUserLoginController() *userLoginController {
	return &userLoginController{
		service: systemService.NewUserLoginService(),
	}
}

// List 获取用户登录信息列表
func (c *userLoginController) List(ctx *gin.Context) {
	req := systemDTO.QueryUserLoginReq{}
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

// Status 更新用户登录信息状态
func (c *userLoginController) Status(ctx *gin.Context) {
	req := systemDTO.UpdateUserLoginStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	row, err := c.service.Status(ctx, req.ID, req.UserId, req.Status)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}
