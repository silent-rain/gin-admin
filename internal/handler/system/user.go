/*
 * @Author: silent-rain
 * @Date: 2023-01-08 21:24:21
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 23:03:08
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/user.go
 * @Descripttion: 用户管理
 */
package system

import (
	systemDao "gin-admin/internal/dao/system"
	"gin-admin/internal/dto"
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/conf"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 用户管理
type userHandler struct {
	dao     systemDao.User
	menuDao systemDao.Menu
}

// 创建角色 Handler 对象
func NewUserHandler() *userHandler {
	return &userHandler{
		dao:     systemDao.NewUserDao(),
		menuDao: systemDao.NewMenuDao(),
	}
}

// All 获取所有用户列表
func (h *userHandler) All(ctx *gin.Context) {
	results, total, err := h.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// List 获取用户列表
func (h *userHandler) List(ctx *gin.Context) {
	req := new(systemDto.QueryUserReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	results, total, err := h.dao.List(*req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Update 更新用户详情信息
func (h *userHandler) Update(ctx *gin.Context) {
	req := new(systemDto.UpdateUserReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 数据转换
	user := new(systemModel.User)
	if err := utils.ApiJsonConvertJson(ctx, req, user); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	roleIds := req.RoleIds
	if err := h.dao.Update(*user, roleIds); err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除用户
func (h *userHandler) Delete(ctx *gin.Context) {
	req := new(dto.DeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.Delete(req.ID)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// BatchDelete 批量删除用户
func (h *userHandler) BatchDelete(ctx *gin.Context) {
	req := new(dto.BatchDeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.BatchDelete(req.Ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbBatchDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbBatchDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新用户状态
func (h *userHandler) Status(ctx *gin.Context) {
	req := new(dto.UpdateStatusReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.Status(req.ID, req.Status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdatePassword 更新密码
func (h *userHandler) UpdatePassword(ctx *gin.Context) {
	req := new(systemDto.UpdateUserPasswordReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 密码加密
	req.OldPassword = utils.Md5(req.OldPassword)
	req.NewPassword = utils.Md5(req.NewPassword)

	// 用户密码验证
	ok, err := h.dao.ExistUserPassword(req.ID, req.OldPassword)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.UserOldPasswordError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.UserOldPasswordError).Json()
		return
	}

	row, err := h.dao.UpdatePassword(req.ID, req.NewPassword)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// ResetPassword 重置密码
func (h *userHandler) ResetPassword(ctx *gin.Context) {
	req := new(systemDto.ResetUserPasswordReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	// 默认密码加密
	password := utils.Md5(conf.ServerUserDefaultPwd)
	row, err := h.dao.ResetPassword(req.ID, password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbResetError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbResetError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdatePhone 更新手机号码
func (h *userHandler) UpdatePhone(ctx *gin.Context) {
	req := new(systemDto.UpdateUserPhoneReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.UpdatePhone(req.ID, req.Phone)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdateEmail 更新邮箱
func (h *userHandler) UpdateEmail(ctx *gin.Context) {
	req := new(systemDto.UpdateUserEmailReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.UpdateEmail(req.ID, req.Email)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Info 获取用户信息
func (h *userHandler) Info(ctx *gin.Context) {
	userId := utils.GetUserId(ctx)
	user, ok, err := h.dao.Info(userId)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.DbQueryEmptyError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryEmptyError).Json()
		return
	}
	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(statuscode.UserDisableError).Error("")
		response.New(ctx).WithCode(statuscode.UserDisableError).Json()
		return
	}

	// 根据角色获取菜单列表
	roleMenus, err := h.getRoleMenuList(user.Roles)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
	}
	// 菜单路由列表：菜单类型为菜单的数据解析
	menus := h.getMenuList(roleMenus)
	// 按钮权限列表：菜单类型为按钮的数据解析
	permissions := h.getPermissionList(roleMenus)

	// 返回值收集
	result := systemDto.UserInfoRsp{
		User:        user,
		Roles:       user.Roles,
		Menus:       menus,
		Permissions: permissions,
	}
	// 避免角色数据反复嵌套
	result.User.Roles = nil
	response.New(ctx).WithData(result).Json()
}

// 根据角色获取菜单列表
func (h *userHandler) getRoleMenuList(roles []systemModel.Role) ([]systemModel.Menu, error) {
	// 获取角色ID
	roleIds := make([]uint, 0)
	for _, item := range roles {
		roleIds = append(roleIds, item.ID)
	}
	// 获取菜单列表
	menus, err := h.menuDao.ListByRoleIds(roleIds)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

// 菜单路由列表：菜单类型为菜单的数据解析
func (h *userHandler) getMenuList(menus []systemModel.Menu) []systemModel.Menu {
	results := make([]systemModel.Menu, 0)
	if len(menus) == 0 {
		return results
	}
	for _, item := range menus {
		if item.MenuType == uint(systemModel.MenuTypeByMenu) {
			results = append(results, item)
		}
	}
	return results
}

// 按钮权限列表：菜单类型为按钮的数据解析
func (h *userHandler) getPermissionList(menus []systemModel.Menu) []string {
	results := make([]string, 0)
	if len(menus) == 0 {
		return results
	}
	for _, item := range menus {
		// 过滤菜单路由，过滤空权限，过滤隐藏按钮
		if item.MenuType == uint(systemModel.MenuTypeByBUtton) &&
			item.Permission != "" &&
			item.Hide == uint(systemModel.MenuHideTypeByShow) {
			results = append(results, item.Permission)
		}
	}
	return results
}
