// Package dao 角色与Http协议接口关联管理
package dao

import (
	"github.com/silent-rain/gin-admin/internal/app/api_auth/dto"
	"github.com/silent-rain/gin-admin/internal/app/api_auth/model"
	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"
	"github.com/silent-rain/gin-admin/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ApiRoleHttpRel 角色与Http协议接口关系接口
type ApiRoleHttpRel interface {
	List(req dto.QueryApiRoleHttpRelReq) ([]model.ApiRoleHttpRel, int64, error)
	Update(roleId uint, menuIds []uint) error
}

// 角色与Http协议接口关系
type apiRoleHttpRel struct {
	mysql.DBRepo
}

// NewApiRoleHttpRelDao 创建角色与Http协议接口关系 Dao 对象
func NewApiRoleHttpRelDao() *apiRoleHttpRel {
	return &apiRoleHttpRel{
		DBRepo: global.Instance().Mysql(),
	}
}

// List 角色关联的Http协议接口列表
func (d *apiRoleHttpRel) List(req dto.QueryApiRoleHttpRelReq) ([]model.ApiRoleHttpRel, int64, error) {
	tx := d.GetDbR()
	if req.ApiId != 0 {
		tx = tx.Where("api_id = ?", req.ApiId)
	}
	if req.RoleId != 0 {
		tx = tx.Where("role_id = ?", req.RoleId)
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&model.ApiRoleHttpRel{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]model.ApiRoleHttpRel, 0)
	result := tx.Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// Update 更新角色关联的Http协议接口
func (d *apiRoleHttpRel) Update(roleId uint, apiIds []uint) error {
	// 未传入 apiIds, 不做处理
	if apiIds == nil {
		return nil
	}
	// 获取角色关联的Http协议接口的 ID 列表
	relIds, err := d.getApiIds(roleId)
	if err != nil {
		return err
	}

	tx := d.GetDbW().Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			zap.S().Panic("更新角色接口关联关系异常, err: %v", err)
		}
	}()

	// 批量添加关系
	if err := d.addRels(tx, relIds, apiIds, roleId); err != nil {
		tx.Rollback()
		return err
	}
	// 批量删除关系
	if err := d.delRels(tx, relIds, apiIds, roleId); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// 获取角色关联的Http协议接口的 ID 列表
func (d *apiRoleHttpRel) getApiIds(roleId uint) ([]uint, error) {
	beans := make([]model.ApiRoleHttpRel, 0)
	results := d.GetDbR().Where("role_id = ?", roleId).Find(&beans)
	if results.Error != nil {
		return nil, results.Error
	}

	apiIds := make([]uint, 0)
	for _, item := range beans {
		apiIds = append(apiIds, item.ApiId)
	}
	return apiIds, nil
}

// 批量添加关系
func (d *apiRoleHttpRel) addRels(tx *gorm.DB, relIds, apiIds []uint, roleId uint) error {
	// 需要新增的关系列表
	addRels := make([]model.ApiRoleHttpRel, 0)
	for _, id := range apiIds {
		if utils.IndexOfArray(relIds, id) == -1 {
			addRels = append(addRels, model.ApiRoleHttpRel{
				RoleId: roleId,
				ApiId:  id,
			})
		}
	}
	if len(addRels) == 0 {
		return nil
	}
	if result := tx.Create(&addRels); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return nil
}

// 批量删除关系
func (d *apiRoleHttpRel) delRels(tx *gorm.DB, relIds, apiIds []uint, roleId uint) error {
	// 需要删除的关系列表
	delRels := make([]uint, 0)
	for _, id := range relIds {
		if utils.IndexOfArray(apiIds, id) == -1 {
			delRels = append(delRels, id)
		}
	}

	if len(delRels) == 0 {
		return nil
	}
	if result := tx.Where("role_id = ? AND api_id in ?", roleId, delRels).
		Delete(&model.ApiRoleHttpRel{}); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return nil
}
