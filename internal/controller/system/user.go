/*用户管理
 */
package system

import (
	"gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/context"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/utils"
	service "gin-admin/internal/service/system"

	"github.com/gin-gonic/gin"
)

// 用户管理
type userController struct {
	service service.UserService
}

// 创建用户对象
func NewUserController() *userController {
	return &userController{
		service: service.NewUserService(),
	}
}

// All 获取所有用户列表
func (c *userController) All(ctx *gin.Context) {
	c.service.All(ctx).Json(ctx)
}

// List 获取用户列表
func (c *userController) List(ctx *gin.Context) {
	req := systemDTO.QueryUserReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.List(ctx, req).Json(ctx)
}

// Add 添加用户
func (c *userController) Add(ctx *gin.Context) {
	req := systemDTO.AddUserReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Add(ctx, req).Json(ctx)
}

// Update 更新用户详情信息
func (c *userController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateUserReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	// 数据转换
	user := systemModel.User{}
	if result := http.ApiJsonConvertJson(ctx, req, &user); result.Error() != nil {
		result.Json(ctx)
		return
	}
	roleIds := req.RoleIds

	c.service.Update(ctx, user, roleIds).Json(ctx)
}

// Delete 删除用户
func (c *userController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Delete(ctx, req.ID).Json(ctx)
}

// BatchDelete 批量删除用户
func (c *userController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.BatchDelete(ctx, req.Ids).Json(ctx)
}

// Status 更新用户状态
func (c *userController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.Status(ctx, req.ID, req.Status).Json(ctx)
}

// UpdatePassword 更新密码
func (c *userController) UpdatePassword(ctx *gin.Context) {
	req := systemDTO.UpdateUserPasswordReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	// 密码加密
	req.OldPassword = utils.Md5(req.OldPassword)
	req.NewPassword = utils.Md5(req.NewPassword)

	c.service.UpdatePassword(ctx, req).Json(ctx)
}

// ResetPassword 重置密码
func (c *userController) ResetPassword(ctx *gin.Context) {
	req := systemDTO.ResetUserPasswordReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	// 默认密码加密
	password := utils.Md5(conf.ServerUserDefaultPwd)

	c.service.ResetPassword(ctx, req.ID, password).Json(ctx)
}

// UpdatePhone 更新手机号码
func (c *userController) UpdatePhone(ctx *gin.Context) {
	req := systemDTO.UpdateUserPhoneReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.UpdatePhone(ctx, req).Json(ctx)
}

// UpdateEmail 更新邮箱
func (c *userController) UpdateEmail(ctx *gin.Context) {
	req := systemDTO.UpdateUserEmailReq{}
	if result := http.ParsingReqParams(ctx, &req); result.Error() != nil {
		result.Json(ctx)
		return
	}

	c.service.UpdateEmail(ctx, req).Json(ctx)
}

// Info 获取用户信息
func (c *userController) Info(ctx *gin.Context) {
	userId := context.GetUserId(ctx)

	c.service.Info(ctx, userId).Json(ctx)
}
