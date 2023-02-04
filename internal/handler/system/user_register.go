/*
 * @Author: silent-rain
 * @Date: 2023-01-08 14:12:59
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 23:38:16
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/user_register.go
 * @Descripttion: 用户注册服务
 */
package system

import (
	systemDao "gin-admin/internal/dao/system"
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserRegisterHandlerImpl 用户注册对象
var UserRegisterHandlerImpl = new(userRegisterHandler)

// 用户注册结构
type userRegisterHandler struct{}

// Add 添加用户
func (h *userRegisterHandler) Add(ctx *gin.Context) {
	// 解析参数
	req := new(systemDto.UserRegisterReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
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
	if err := utils.ApiJsonConvertJson(ctx, req, user); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	user.Status = 1
	roleIds := req.RoleIds

	// 数据入库
	if err := systemDao.NewDaoUserRegister().Add(*user, roleIds); err != nil {
		log.New(ctx).WithCode(statuscode.UserRegisterError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.UserRegisterError).Json()
		return
	}
	response.New(ctx).WithMsg("用户注册成功").Json()
}

// 检查手机号是否存在
func (h *userRegisterHandler) chechkPhone(ctx *gin.Context, phone string) bool {
	if phone == "" {
		return false
	}
	if _, ok, err := systemDao.NewDaoUser().GetUserByPhone(phone); err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		return false
	} else if !ok {
		return false
	}
	log.New(ctx).WithCode(statuscode.DbDataExistError).Error("手机号已存在")
	return true
}

// 检查邮箱是否存在
func (h *userRegisterHandler) chechkEmail(ctx *gin.Context, email string) bool {
	if email == "" {
		return false
	}
	if _, ok, err := systemDao.NewDaoUser().GetUserByEmail(email); err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		return false
	} else if !ok {

		return false
	}
	log.New(ctx).WithCode(statuscode.DbDataExistError).Error("邮箱已存在")
	return true
}
