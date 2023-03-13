/*登录*/
package system

import (
	"errors"

	permissionModel "github.com/silent-rain/gin-admin/internal/model/permission"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// Login 登录接口
type Login interface {
	Login(username, password string) (permissionModel.User, bool, error)
}

// 登录
type login struct {
	db mysql.DBRepo
}

// NewLoginDao 创建登录对象
func NewLoginDao() *login {
	return &login{
		db: mysql.Instance(),
	}
}

// Login 查询登录用户信息 邮件/手机号
func (d *login) Login(username, password string) (permissionModel.User, bool, error) {
	bean := permissionModel.User{}
	result := d.db.GetDbR().
		Where("(phone = ? OR email = ?) AND password = ?", username, username, password).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return permissionModel.User{}, false, nil
	}
	if result.Error != nil {
		return permissionModel.User{}, false, result.Error
	}
	return bean, true, nil
}
