// Package initialize 数据库表自动迁移
package initialize

import (
	apiAuthModel "github.com/silent-rain/gin-admin/internal/app/api_auth/model"
	dataCenterModel "github.com/silent-rain/gin-admin/internal/app/data_center/model"
	logModel "github.com/silent-rain/gin-admin/internal/app/log/model"
	permissionModel "github.com/silent-rain/gin-admin/internal/app/permission/model"
	systemModel "github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/global"
)

// 数据库表自动迁移
// 迁移 schema
// 仅让 sqlite3 进行自动迁移
func initTableAutoMigrate() {
	if global.Instance().Config().Sqlite == nil {
		return
	}
	db := global.Instance().Mysql().GetDbW()
	db.AutoMigrate(&apiAuthModel.ApiHttp{})
	db.AutoMigrate(&apiAuthModel.ApiRoleHttpRel{})
	db.AutoMigrate(&dataCenterModel.DictData{})
	db.AutoMigrate(&dataCenterModel.Dict{})
	db.AutoMigrate(&logModel.HttpLog{})
	db.AutoMigrate(&logModel.SystemLog{})
	db.AutoMigrate(&logModel.WebLog{})
	db.AutoMigrate(&permissionModel.Menu{})
	db.AutoMigrate(&permissionModel.RoleMenuRel{})
	db.AutoMigrate(&permissionModel.Role{})
	db.AutoMigrate(&permissionModel.UserApiToken{})
	db.AutoMigrate(&permissionModel.User{})
	db.AutoMigrate(&permissionModel.UserRoleRel{})
	db.AutoMigrate(&systemModel.Config{})
	db.AutoMigrate(&systemModel.UserLogin{})
}
