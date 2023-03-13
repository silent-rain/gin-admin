/*用户登录/登出/注册
 */
package system

import (
	permissionDTO "github.com/silent-rain/gin-admin/internal/dto/permission"
	systemDTO "github.com/silent-rain/gin-admin/internal/dto/system"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/response"
	systemService "github.com/silent-rain/gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 用户登录/登出/注册
type userLoginRegisterController struct {
	service systemService.UserLoginRegisterService
}

// NewUserLoginRegisterController 创建用户登录/登出/注册 对象
func NewUserLoginRegisterController() *userLoginRegisterController {
	return &userLoginRegisterController{
		service: systemService.NewUserLoginRegisterService(),
	}
}

// Login 登录
func (c *userLoginRegisterController) Login(ctx *gin.Context) {
	req := systemDTO.UserLoginReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	result, err := c.service.Login(ctx, req)
	if err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}

// Logout 注销系统
func (c *userLoginRegisterController) Logout(ctx *gin.Context) {
	if _, err := c.service.Logout(ctx); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Register 注册用户
func (c *userLoginRegisterController) Register(ctx *gin.Context) {
	req := permissionDTO.AddUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}

	if err := c.service.Register(ctx, req); err != nil {
		response.New(ctx).WithError(err).Json()
		return
	}
	response.New(ctx).Json()
}
