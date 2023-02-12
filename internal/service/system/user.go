/*用户管理
 */
package service

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/http"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserService 用户管理
type UserService interface {
	All(ctx *gin.Context)
	List(ctx *gin.Context, req systemDTO.QueryUserReq)
	Add(ctx *gin.Context, req systemDTO.AddUserReq)
	Update(ctx *gin.Context, user systemModel.User, roleIds []uint)
	Delete(ctx *gin.Context, id uint)
	BatchDelete(ctx *gin.Context, ids []uint)
	Status(ctx *gin.Context, id uint, status uint)
	UpdatePassword(ctx *gin.Context, req systemDTO.UpdateUserPasswordReq)
	ResetPassword(ctx *gin.Context, id uint, password string)
	UpdatePhone(ctx *gin.Context, req systemDTO.UpdateUserPhoneReq)
	UpdateEmail(ctx *gin.Context, req systemDTO.UpdateUserEmailReq)
	Info(ctx *gin.Context, userId uint)
}

// 用户管理
type userService struct {
	dao     systemDAO.User
	menuDao systemDAO.Menu
}

// 创建角色对象
func NewUserService() *userService {
	return &userService{
		dao:     systemDAO.NewUserDao(),
		menuDao: systemDAO.NewMenuDao(),
	}
}

// All 获取所有用户列表
func (h *userService) All(ctx *gin.Context) {
	results, total, err := h.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// List 获取用户列表
func (h *userService) List(ctx *gin.Context, req systemDTO.QueryUserReq) {
	results, total, err := h.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(results, total).Json()
}

// Add 添加用户
func (h *userService) Add(ctx *gin.Context, req systemDTO.AddUserReq) {
	// 注册入口检查验证码
	if ctx.Request.URL.Path == "/api/v1/register" {
		if !chechkCaptcha(ctx, req.CaptchaId, req.Captcha) {
			return
		}
	}

	// 判断用户是否存在 邮件/手机号
	if h.chechkPhone(ctx, req.Phone) {
		response.New(ctx).WithCode(statuscode.ExistPhoneError).WithMsg("手机号已存在").Json()
		return
	}
	if h.chechkEmail(ctx, req.Email) {
		response.New(ctx).WithCode(statuscode.ExistEmailError).WithMsg("邮箱已存在").Json()
		return
	}

	// 密码加密
	req.Password = utils.Md5(req.Password)

	// 数据转换
	user := new(systemModel.User)
	if err := http.ApiJsonConvertJson(ctx, req, user); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	user.Status = 1

	// 数据入库
	if err := h.dao.Add(*user, req.RoleIds); err != nil {
		log.New(ctx).WithCode(statuscode.UserRegisterError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.UserRegisterError).Json()
		return
	}
	response.New(ctx).WithMsg("用户注册成功").Json()
}

// 检查手机号是否存在
func (h *userService) chechkPhone(ctx *gin.Context, phone string) bool {
	if phone == "" {
		return false
	}
	if _, ok, err := h.dao.GetUserByPhone(phone); err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return false
	} else if !ok {
		return false
	}
	log.New(ctx).WithCode(statuscode.DBDataExistError).Error("手机号已存在")
	return true
}

// 检查邮箱是否存在
func (h *userService) chechkEmail(ctx *gin.Context, email string) bool {
	if email == "" {
		return false
	}
	if _, ok, err := h.dao.GetUserByEmail(email); err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return false
	} else if !ok {

		return false
	}
	log.New(ctx).WithCode(statuscode.DBDataExistError).Error("邮箱已存在")
	return true
}

// Update 更新用户详情信息
func (h *userService) Update(ctx *gin.Context, user systemModel.User, roleIds []uint) {
	if err := h.dao.Update(user, roleIds); err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateError).Json()
		return
	}
	response.New(ctx).Json()
}

// Delete 删除用户
func (h *userService) Delete(ctx *gin.Context, id uint) {
	row, err := h.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// BatchDelete 批量删除用户
func (h *userService) BatchDelete(ctx *gin.Context, ids []uint) {
	row, err := h.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBBatchDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBBatchDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新用户状态
func (h *userService) Status(ctx *gin.Context, id uint, status uint) {
	row, err := h.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdatePassword 更新密码
func (h *userService) UpdatePassword(ctx *gin.Context, req systemDTO.UpdateUserPasswordReq) {
	// 用户密码验证
	ok, err := h.dao.ExistUserPassword(req.ID, req.OldPassword)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.UserOldPasswordError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.UserOldPasswordError).Json()
		return
	}

	row, err := h.dao.UpdatePassword(req.ID, req.NewPassword)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// ResetPassword 重置密码
func (h *userService) ResetPassword(ctx *gin.Context, id uint, password string) {
	row, err := h.dao.ResetPassword(id, password)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBResetError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBResetError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdatePhone 更新手机号码
func (h *userService) UpdatePhone(ctx *gin.Context, req systemDTO.UpdateUserPhoneReq) {
	// 查看手机号码是否已经被非本人使用
	user, ok, err := h.dao.GetUserByPhone(req.Phone)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	if ok && user.ID == req.ID {
		log.New(ctx).WithCode(statuscode.UserPhoneConsistentError).Error("")
		response.New(ctx).WithCode(statuscode.UserPhoneConsistentError).Json()
		return
	}
	if ok {
		log.New(ctx).WithCode(statuscode.DBDataExistError).Error("手机号码已被使用")
		response.New(ctx).WithCode(statuscode.DBDataExistError).WithMsg("手机号码已被使用").Json()
		return
	}
	row, err := h.dao.UpdatePhone(req.ID, req.Phone)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// UpdateEmail 更新邮箱
func (h *userService) UpdateEmail(ctx *gin.Context, req systemDTO.UpdateUserEmailReq) {
	// 查看邮箱是否已经被非本人使用
	user, ok, err := h.dao.GetUserByEmail(req.Email)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	if ok && user.ID == req.ID {
		log.New(ctx).WithCode(statuscode.UserEmailConsistentError).Error("")
		response.New(ctx).WithCode(statuscode.UserEmailConsistentError).Json()
		return
	}
	if ok {
		log.New(ctx).WithCode(statuscode.DBDataExistError).Error("邮箱已被使用")
		response.New(ctx).WithCode(statuscode.DBDataExistError).WithMsg("邮箱已被使用").Json()
		return
	}
	row, err := h.dao.UpdateEmail(req.ID, req.Email)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Info 获取用户信息
func (h *userService) Info(ctx *gin.Context, userId uint) {
	user, ok, err := h.dao.Info(userId)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
		return
	}
	if !ok {
		log.New(ctx).WithCode(statuscode.DBQueryEmptyError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryEmptyError).Json()
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
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DBQueryError).Json()
	}
	// 菜单路由列表：菜单类型为菜单的数据解析
	menus := h.getMenuList(roleMenus)
	// 菜单列表数据转为树结构
	menus = MenuListToTree(menus, nil)
	// 按钮权限列表：菜单类型为按钮的数据解析
	permissions := h.getPermissionList(roleMenus)

	// 返回值收集
	result := systemDTO.UserInfoRsp{
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
func (h *userService) getRoleMenuList(roles []systemModel.Role) ([]systemModel.Menu, error) {
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
func (h *userService) getMenuList(menus []systemModel.Menu) []systemModel.Menu {
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
func (h *userService) getPermissionList(menus []systemModel.Menu) []systemDTO.ButtonPermission {
	results := make([]systemDTO.ButtonPermission, 0)
	if len(menus) == 0 {
		return results
	}
	for _, item := range menus {
		// 过滤禁用按钮, 过滤菜单路由，过滤空权限
		if item.Status == 1 && item.MenuType == uint(systemModel.MenuTypeByBUtton) &&
			item.Permission != "" {
			results = append(results, systemDTO.ButtonPermission{
				Permission: item.Permission,
				Disabled:   item.Hidden,
			})
		}
	}
	return results
}
