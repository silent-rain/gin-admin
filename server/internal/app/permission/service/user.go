// Package service 用户管理
package service

import (
	"github.com/silent-rain/gin-admin/internal/app/permission/dao"
	"github.com/silent-rain/gin-admin/internal/app/permission/dto"
	"github.com/silent-rain/gin-admin/internal/app/permission/model"
	"github.com/silent-rain/gin-admin/internal/pkg/http"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"
	"github.com/silent-rain/gin-admin/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserService 用户管理
type UserService interface {
	All(ctx *gin.Context) ([]model.User, int64, error)
	List(ctx *gin.Context, req dto.QueryUserReq) ([]model.User, int64, error)
	Add(ctx *gin.Context, req dto.AddUserReq) error
	Update(ctx *gin.Context, user model.User, roleIds []uint) error
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
	UpdatePassword(ctx *gin.Context, req dto.UpdateUserPasswordReq) (int64, error)
	ResetPassword(ctx *gin.Context, id uint, password string) (int64, error)
	UpdatePhone(ctx *gin.Context, req dto.UpdateUserPhoneReq) (int64, error)
	UpdateEmail(ctx *gin.Context, req dto.UpdateUserEmailReq) (int64, error)
	Info(ctx *gin.Context, userId uint) (dto.UserInfoRsp, error)
}

// 用户管理
type userService struct {
	dao     dao.User
	menuDao dao.Menu
}

// 创建角色对象
func NewUserService() *userService {
	return &userService{
		dao:     dao.NewUserDao(),
		menuDao: dao.NewMenuDao(),
	}
}

// All 获取所有用户列表
func (h *userService) All(ctx *gin.Context) ([]model.User, int64, error) {
	results, total, err := h.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	return results, total, nil
}

// List 获取用户列表
func (h *userService) List(ctx *gin.Context, req dto.QueryUserReq) ([]model.User, int64, error) {
	results, total, err := h.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	return results, total, nil
}

// Add 添加用户
func (h *userService) Add(ctx *gin.Context, req dto.AddUserReq) error {
	// 判断用户是否存在 邮件/手机号
	if h.chechkPhone(ctx, req.Phone) {
		return errcode.ExistPhoneError.WithMsg("手机号已存在")
	}
	if h.chechkEmail(ctx, req.Email) {
		return errcode.ExistEmailError.WithMsg("邮箱已存在")
	}

	// 密码加密
	req.Password = utils.EncryptMd5(req.Password)

	// 数据转换
	user := new(model.User)
	if err := http.ApiJsonConvertJson(ctx, req, user); err != nil {
		return err
	}
	user.Status = 1

	// 数据入库
	if err := h.dao.Add(*user, req.RoleIds); err != nil {
		log.New(ctx).WithCode(errcode.UserRegisterError).Errorf("%v", err)
		return errcode.UserRegisterError
	}
	return nil
}

// 检查手机号是否存在
func (h *userService) chechkPhone(ctx *gin.Context, phone string) bool {
	if phone == "" {
		return false
	}
	if _, ok, err := h.dao.GetUserByPhone(phone); err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return false
	} else if !ok {
		return false
	}
	log.New(ctx).WithCode(errcode.DBDataExistError).Error("手机号已存在")
	return true
}

// 检查邮箱是否存在
func (h *userService) chechkEmail(ctx *gin.Context, email string) bool {
	if email == "" {
		return false
	}
	if _, ok, err := h.dao.GetUserByEmail(email); err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return false
	} else if !ok {

		return false
	}
	log.New(ctx).WithCode(errcode.DBDataExistError).Error("邮箱已存在")
	return true
}

// Update 更新用户详情信息
func (h *userService) Update(ctx *gin.Context, user model.User, roleIds []uint) error {
	if err := h.dao.Update(user, roleIds); err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return errcode.DBUpdateError
	}
	return nil
}

// Delete 删除用户
func (h *userService) Delete(ctx *gin.Context, id uint) (int64, error) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}
	return row, nil
}

// BatchDelete 批量删除用户
func (h *userService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}
	return row, nil
}

// Status 更新用户状态
func (h *userService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}
	return row, nil
}

// UpdatePassword 更新密码
func (h *userService) UpdatePassword(ctx *gin.Context, req dto.UpdateUserPasswordReq) (int64, error) {
	// 用户密码验证
	ok, err := h.dao.ExistUserPassword(req.ID, req.OldPassword)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.UserOldPasswordError).Errorf("%v", err)
		return 0, errcode.UserOldPasswordError
	}

	row, err := h.dao.UpdatePassword(req.ID, req.NewPassword)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// ResetPassword 重置密码
func (h *userService) ResetPassword(ctx *gin.Context, id uint, password string) (int64, error) {
	row, err := h.dao.ResetPassword(id, password)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBResetError).Errorf("%v", err)
		return 0, errcode.DBResetError
	}
	return row, nil
}

// UpdatePhone 更新手机号码
func (h *userService) UpdatePhone(ctx *gin.Context, req dto.UpdateUserPhoneReq) (int64, error) {
	// 用户密码验证
	ok, err := h.dao.ExistUserPassword(req.ID, req.Password)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.UserPasswordError).Errorf("%v", err)
		return 0, errcode.UserPasswordError
	}

	// 查看手机号码是否已经被非本人使用
	user, ok, err := h.dao.GetUserByPhone(req.Phone)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if ok && user.ID == req.ID {
		log.New(ctx).WithCode(errcode.UserPhoneConsistentError).Error("")
		return 0, errcode.UserPhoneConsistentError
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Error("手机号码已被使用")
		return 0, errcode.DBDataExistError.WithMsg("手机号码已被使用")
	}
	row, err := h.dao.UpdatePhone(req.ID, req.Phone)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// UpdateEmail 更新邮箱
func (h *userService) UpdateEmail(ctx *gin.Context, req dto.UpdateUserEmailReq) (int64, error) {
	// 用户密码验证
	ok, err := h.dao.ExistUserPassword(req.ID, req.Password)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.UserPasswordError).Errorf("%v", err)
		return 0, errcode.UserPasswordError
	}

	// 查看邮箱是否已经被非本人使用
	user, ok, err := h.dao.GetUserByEmail(req.Email)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if ok && user.ID == req.ID {
		log.New(ctx).WithCode(errcode.UserEmailConsistentError).Error("")
		return 0, errcode.UserEmailConsistentError
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Error("邮箱已被使用")
		return 0, errcode.DBDataExistError.WithMsg("邮箱已被使用")
	}
	row, err := h.dao.UpdateEmail(req.ID, req.Email)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}
	return row, nil
}

// Info 获取用户信息
func (h *userService) Info(ctx *gin.Context, userId uint) (dto.UserInfoRsp, error) {
	user, ok, err := h.dao.Info(userId)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return dto.UserInfoRsp{}, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Errorf("%v", err)
		return dto.UserInfoRsp{}, errcode.DBQueryEmptyError
	}
	// 判断当前用户状态
	if user.Status != 1 {
		log.New(ctx).WithCode(errcode.UserDisableError).Error("")
		return dto.UserInfoRsp{}, errcode.UserDisableError
	}

	// 根据角色获取菜单列表
	roleMenus, err := h.getRoleMenuList(user.Roles)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return dto.UserInfoRsp{}, errcode.DBQueryError
	}
	// 菜单路由列表：菜单类型为菜单的数据解析
	menus := h.getMenuList(roleMenus)
	// 菜单列表数据转为树结构
	menus = menuListToTree(menus, nil)
	// 按钮权限列表：菜单类型为按钮的数据解析
	permissions := h.getPermissionList(roleMenus)

	// 返回值收集
	result := dto.UserInfoRsp{
		User:        user,
		Roles:       user.Roles,
		Menus:       menus,
		Permissions: permissions,
	}
	// 避免角色数据反复嵌套
	result.User.Roles = nil
	return result, nil
}

// 根据角色获取菜单列表
func (h *userService) getRoleMenuList(roles []model.Role) ([]model.Menu, error) {
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
func (h *userService) getMenuList(menus []model.Menu) []model.Menu {
	results := make([]model.Menu, 0)
	if len(menus) == 0 {
		return results
	}
	for _, item := range menus {
		if item.MenuType == uint(model.MenuTypeByMenu) {
			results = append(results, item)
		}
	}
	return results
}

// 按钮权限列表：菜单类型为按钮的数据解析
func (h *userService) getPermissionList(menus []model.Menu) []dto.ButtonPermission {
	results := make([]dto.ButtonPermission, 0)
	if len(menus) == 0 {
		return results
	}
	for _, item := range menus {
		// 过滤禁用按钮, 过滤菜单路由，过滤空权限
		if item.Status == 1 && item.MenuType == uint(model.MenuTypeByButton) &&
			item.Permission != "" {
			results = append(results, dto.ButtonPermission{
				Permission: item.Permission,
				Disabled:   item.Hidden,
			})
		}
	}
	return results
}
