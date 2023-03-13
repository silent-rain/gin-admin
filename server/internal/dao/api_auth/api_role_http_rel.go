/*角色与Http协议接口关联表*/
package apiauth

import (
	DAO "github.com/silent-rain/gin-admin/internal/dao"
	apiAuthDTO "github.com/silent-rain/gin-admin/internal/dto/api_auth"
	apiAuthModel "github.com/silent-rain/gin-admin/internal/model/api_auth"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"
	"github.com/silent-rain/gin-admin/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ApiRoleHttpRel 角色与Http协议接口关系接口
type ApiRoleHttpRel interface {
	List(req apiAuthDTO.QueryApiRoleHttpRelReq) ([]apiAuthModel.ApiRoleHttpRel, int64, error)
	Update(roleId uint, menuIds []uint) error
}

// 角色与Http协议接口关系
type apiRoleHttpRel struct {
	*DAO.Transaction
	db mysql.DBRepo
}

// NewApiRoleHttpRelDao 创建角色与Http协议接口关系 Dao 对象
func NewApiRoleHttpRelDao() *apiRoleHttpRel {
	return &apiRoleHttpRel{
		Transaction: DAO.NewTransaction(mysql.Instance().GetDbW()),
		db:          mysql.Instance(),
	}
}

// List 角色关联的Http协议接口列表
func (d *apiRoleHttpRel) List(req apiAuthDTO.QueryApiRoleHttpRelReq) ([]apiAuthModel.ApiRoleHttpRel, int64, error) {
	tx := d.db.GetDbR()
	if req.ApiId != 0 {
		tx = tx.Where("api_id = ?", req.ApiId)
	}
	if req.RoleId != 0 {
		tx = tx.Where("role_id = ?", req.RoleId)
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&apiAuthModel.ApiRoleHttpRel{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]apiAuthModel.ApiRoleHttpRel, 0)
	result := tx.Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// Update 更新角色关联的Http协议接口
func (d *apiRoleHttpRel) Update(roleId uint, apiIds []uint) error {
	d.Begin()
	defer func() {
		if err := recover(); err != nil {
			d.Rollback()
			zap.S().Panic("更新角色接口关联关系异常, err: %v", err)
		}
	}()

	// 未传入 apiIds, 不做处理
	if apiIds == nil {
		return nil
	}
	// 获取角色关联的Http协议接口的 ID 列表
	relIds, err := d.getApiIds(roleId)
	if err != nil {
		return err
	}

	// 需要新增的关系列表
	addRels := make([]apiAuthModel.ApiRoleHttpRel, 0)
	for _, id := range apiIds {
		if utils.IndexOfArray(relIds, id) == -1 {
			addRels = append(addRels, apiAuthModel.ApiRoleHttpRel{
				RoleId: roleId,
				ApiId:  id,
			})
		}
	}

	// 需要删除的关系列表
	delRels := make([]uint, 0)
	for _, id := range relIds {
		if utils.IndexOfArray(apiIds, id) == -1 {
			delRels = append(delRels, id)
		}
	}

	if len(addRels) != 0 {
		if result := d.Tx().Debug().Create(&addRels); result.Error != nil {
			return result.Error
		}
	}
	if len(delRels) != 0 {
		if result := d.Tx().Debug().Where("role_id = ? AND api_id in ?", roleId, delRels).
			Delete(&apiAuthModel.ApiRoleHttpRel{}); result.Error != nil {
			return result.Error
		}
	}
	d.Commit()
	return nil
}

// 获取角色关联的Http协议接口的 ID 列表
func (d *apiRoleHttpRel) getApiIds(roleId uint) ([]uint, error) {
	beans := make([]apiAuthModel.ApiRoleHttpRel, 0)
	results := d.Tx().Where("role_id = ?", roleId).Find(&beans)
	if results.Error != nil {
		return nil, results.Error
	}

	apiIds := make([]uint, 0)
	for _, item := range beans {
		apiIds = append(apiIds, item.ApiId)
	}
	return apiIds, nil
}
