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
	"gin-admin/internal/pkg/log"
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
	c.service.All(ctx)
}

// List 获取用户列表
func (c *userController) List(ctx *gin.Context) {
	req := systemDTO.QueryUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		return
	}

	c.service.List(ctx, req)
}

// Add 添加用户
func (c *userController) Add(ctx *gin.Context) {
	req := systemDTO.AddUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.Add(ctx, req)
}

// Update 更新用户详情信息
func (c *userController) Update(ctx *gin.Context) {
	req := systemDTO.UpdateUserReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 数据转换
	user := systemModel.User{}
	if err := http.ApiJsonConvertJson(ctx, req, &user); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	roleIds := req.RoleIds

	c.service.Update(ctx, user, roleIds)
}

// Delete 删除用户
func (c *userController) Delete(ctx *gin.Context) {
	req := dto.DeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.Delete(ctx, req.ID)
}

// BatchDelete 批量删除用户
func (c *userController) BatchDelete(ctx *gin.Context) {
	req := dto.BatchDeleteReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.BatchDelete(ctx, req.Ids)
}

// Status 更新用户状态
func (c *userController) Status(ctx *gin.Context) {
	req := dto.UpdateStatusReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.Status(ctx, req.ID, req.Status)
}

// UpdatePassword 更新密码
func (c *userController) UpdatePassword(ctx *gin.Context) {
	req := systemDTO.UpdateUserPasswordReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 密码加密
	req.OldPassword = utils.Md5(req.OldPassword)
	req.NewPassword = utils.Md5(req.NewPassword)

	c.service.UpdatePassword(ctx, req)
}

// ResetPassword 重置密码
func (c *userController) ResetPassword(ctx *gin.Context) {
	req := systemDTO.ResetUserPasswordReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 默认密码加密
	password := utils.Md5(conf.ServerUserDefaultPwd)

	c.service.ResetPassword(ctx, req.ID, password)
}

// UpdatePhone 更新手机号码
func (c *userController) UpdatePhone(ctx *gin.Context) {
	req := systemDTO.UpdateUserPhoneReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.UpdatePhone(ctx, req)
}

// UpdateEmail 更新邮箱
func (c *userController) UpdateEmail(ctx *gin.Context) {
	req := systemDTO.UpdateUserEmailReq{}
	if err := http.ParsingReqParams(ctx, &req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	c.service.UpdateEmail(ctx, req)
}

// Info 获取用户信息
func (c *userController) Info(ctx *gin.Context) {
	userId := context.GetUserId(ctx)

	c.service.Info(ctx, userId)
}
