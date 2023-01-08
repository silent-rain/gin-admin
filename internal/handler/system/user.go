/*
 * @Author: silent-rain
 * @Date: 2023-01-08 21:24:21
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:35:48
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/user.go
 * @Descripttion: 用户管理
 */
package system

import (
	"gin-admin/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// UserManageImpl 用户登录注册对象
var UserManageImpl = new(userManageHandler)

// 用户管理
type userManageHandler struct {
}

// UserInfo 获取用户信息
func (h *userManageHandler) UserInfo(ctx *gin.Context) {
	response.New(ctx).Json()
}

/*
   userInfo: {
     headImgUrl: "logo_W2xEX_2x.jpg",
     phone: "13302254696",
     roleId: "[13]",
     name: "panda",
     id: 5,
   },
*/
