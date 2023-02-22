/*用户管理
 */
package system

import (
	"gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/constant"
	"gin-admin/internal/pkg/core"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/response"
	"gin-admin/internal/pkg/utils"
	systemService "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 用户管理
type userController struct {
	service systemService.UserService
}

// 创建用户对象
func NewUserController() *userController {
	return &userController{
		service: systemService.NewUserService(),
	}
}

// All 获取所有用户列表
func (c *userController) All(ctx *gin.Context) {
	results, total, err := c.service.All(ctx)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// List 获取用户列表
func (c *userController) List(ctx *gin.Context) {
	req := systemDTO.QueryUserReq{}
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

// Add 添加用户
func (c *userController) Add(ctx *gin.Context) {
	req := systemDTO.AddUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	if err := c.service.Add(ctx, req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新用户详情信息
func (c *userController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	// 数据转换
	user := systemModel.User{}
	if err := http.ApiJsonConvertJson(ctx, req, &user); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	roleIds := req.RoleIds

	if err := c.service.Update(ctx, user, roleIds); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除用户
func (c *userController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	if _, err := c.service.Delete(ctx, req.ID); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// BatchDelete 批量删除用户
func (c *userController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	if _, err := c.service.BatchDelete(ctx, req.Ids); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Status 更新用户状态
func (c *userController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	if _, err := c.service.Status(ctx, req.ID, req.Status); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdatePassword 更新密码
func (c *userController) UpdatePassword(ctx *gin.Context) {
	req := systemDTO.UpdateUserPasswordReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	// 密码加密
	req.OldPassword = utils.Md5(req.OldPassword)
	req.NewPassword = utils.Md5(req.NewPassword)

	if _, err := c.service.UpdatePassword(ctx, req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// ResetPassword 重置密码
func (c *userController) ResetPassword(ctx *gin.Context) {
	req := systemDTO.ResetUserPasswordReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	// 默认密码加密
	password := utils.Md5(constant.ServerUserDefaultPwd)

	if _, err := c.service.ResetPassword(ctx, req.ID, password); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdatePhone 更新手机号码
func (c *userController) UpdatePhone(ctx *gin.Context) {
	req := systemDTO.UpdateUserPhoneReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	if _, err := c.service.UpdatePhone(ctx, req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// UpdateEmail 更新邮箱
func (c *userController) UpdateEmail(ctx *gin.Context) {
	req := systemDTO.UpdateUserEmailReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}

	if _, err := c.service.UpdateEmail(ctx, req); err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).Json()
}

// Info 获取用户信息
func (c *userController) Info(ctx *gin.Context) {
	userId := core.GetContext(ctx).UserId

	result, err := c.service.Info(ctx, userId)
	if err != nil {
		response.New(ctx).WithCodeError(err).Json()
		return
	}
	response.New(ctx).WithData(result).Json()
}
