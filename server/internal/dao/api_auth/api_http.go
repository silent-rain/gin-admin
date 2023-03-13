/*Http协议接口管理表*/
package apiauth

import (
	"context"
	"errors"
	"time"

	apiAuthDTO "github.com/silent-rain/gin-admin/internal/dto/api_auth"
	apiAuthModel "github.com/silent-rain/gin-admin/internal/model/api_auth"
	"github.com/silent-rain/gin-admin/internal/pkg/constant"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/redis"

	"gorm.io/gorm"
)

// ApiHttp Http协议接口信息接口
type ApiHttp interface {
	All() ([]apiAuthModel.ApiHttp, int64, error)
	List(req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error)
	InfoByUri(uri string) (apiAuthModel.ApiHttp, bool, error)
	Add(bean apiAuthModel.ApiHttp) (uint, error)
	Update(bean apiAuthModel.ApiHttp) (int64, error)
	Delete(id uint) (int64, error)
	BatchDelete(ids []uint) (int64, error)
	Status(id uint, status uint) (int64, error)
	Children(parentId uint) ([]apiAuthModel.ApiHttp, error)
	GetUriListByToken(token, uri string) (apiAuthModel.ApiHttp, bool, error)
}

// Http协议接口信息
type apiAuth struct {
	db mysql.DBRepo
}

// NewApiHttpDao 创建Http协议接口 Dao 对象
func NewApiHttpDao() *apiAuth {
	return &apiAuth{
		db: mysql.Instance(),
	}
}

// All 获取所有Http协议接口列表
func (d *apiAuth) All() ([]apiAuthModel.ApiHttp, int64, error) {
	tx := d.db.GetDbR().Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&apiAuthModel.ApiHttp{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]apiAuthModel.ApiHttp, 0)
	if result := tx.Order("updated_at ASC").Find(&bean); result.Error != nil {
		return nil, 0, result.Error
	}

	return bean, total, nil
}

// List 查询Http协议接口列表
func (d *apiAuth) List(req apiAuthDTO.QueryApiHttpReq) ([]apiAuthModel.ApiHttp, int64, error) {
	tx := d.db.GetDbR()
	if req.Method != "" {
		tx = tx.Where("method = ?", req.Method)
	}
	if req.Status != nil {
		tx = tx.Where("status = ?", *req.Status)
	}
	if req.Name != "" {
		tx = tx.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Uri != "" {
		tx = tx.Where("uri like ?", req.Uri+"%")
	}
	tx = tx.Session(&gorm.Session{})

	var total int64 = 0
	if result := tx.Model(&apiAuthModel.ApiHttp{}).Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	bean := make([]apiAuthModel.ApiHttp, 0)
	result := tx.Offset(req.Offset()).Limit(req.PageSize).Order("id ASC").
		Find(&bean)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return bean, total, nil
}

// InfoByUri 获取Http协议接口信息
func (d *apiAuth) InfoByUri(uri string) (apiAuthModel.ApiHttp, bool, error) {
	bean := apiAuthModel.ApiHttp{}
	result := d.db.GetDbR().Where("uri = ?", uri).First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// Add 添加Http协议接口
func (d *apiAuth) Add(bean apiAuthModel.ApiHttp) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}

// Update 更新Http协议接口
func (d *apiAuth) Update(bean apiAuthModel.ApiHttp) (int64, error) {
	result := d.db.GetDbW().Select("*").Omit("created_at").Updates(&bean)
	return result.RowsAffected, result.Error
}

// Delete 删除Http协议接口
func (d *apiAuth) Delete(id uint) (int64, error) {
	result := d.db.GetDbW().Delete(&apiAuthModel.ApiHttp{
		ID: id,
	})
	return result.RowsAffected, result.Error
}

// BatchDelete 批量删除Http协议接口
func (d *apiAuth) BatchDelete(ids []uint) (int64, error) {
	beans := make([]apiAuthModel.ApiHttp, len(ids))
	for _, id := range ids {
		beans = append(beans, apiAuthModel.ApiHttp{
			ID: id,
		})
	}
	result := d.db.GetDbW().Delete(&beans)
	return result.RowsAffected, result.Error
}

// Status 更新状态
func (d *apiAuth) Status(id uint, status uint) (int64, error) {
	result := d.db.GetDbW().Select("status").Updates(&apiAuthModel.ApiHttp{
		ID:     id,
		Status: status,
	})
	return result.RowsAffected, result.Error
}

// Children 通过父 ID 获取子配置列表
func (d *apiAuth) Children(parentId uint) ([]apiAuthModel.ApiHttp, error) {
	beans := make([]apiAuthModel.ApiHttp, 0)
	result := d.db.GetDbR().Where("parent_id = ?", parentId).
		Order("sort ASC").Order("id ASC").
		Find(&beans)
	if result.Error != nil {
		return nil, result.Error
	}
	return beans, nil
}

// GetUriListByToken 获取 Token 令牌对应的 URI 资源列表
func (d *apiAuth) GetUriListByToken(token, uri string) (apiAuthModel.ApiHttp, bool, error) {
	bean := apiAuthModel.ApiHttp{}
	result := d.db.GetDbR().Model(&apiAuthModel.ApiHttp{}).
		Joins("left join api_role_http_rel arhr on arhr.api_id = api_http.id").
		Joins("left join perm_user_role_rel purr on purr.role_id = arhr.role_id").
		Joins("left join perm_user_api_token puat on puat.user_id = purr.user_id").
		Where("puat.token = ?", token).
		Where("api_http.uri = ?", uri).
		Group("api_http.id").
		First(&bean)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return bean, false, nil
	}
	if result.Error != nil {
		return bean, false, result.Error
	}
	return bean, true, nil
}

// ApiTokenLoginCache API Token 登录信息缓存接口
type ApiTokenLoginCache interface {
	Set(userId uint, token string) error
	Get(userId uint) (string, error)
}

// API Token 登录信息缓存
type redisApiTokenLogin struct {
	db redis.DBRepo
}

// NewApiTokenLoginCacheDao 创建 API Token 登录信息缓存对象
func NewApiTokenLoginCacheDao() *redisApiTokenLogin {
	return &redisApiTokenLogin{
		db: redis.Instance().DB(redis.ApiTokenLogin),
	}
}

// Set 设置缓存
func (d *redisApiTokenLogin) Set(tokenUri string, userId uint, Nickname string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	user := apiAuthDTO.ApiHttpUserCache{
		UserId:   userId,
		Nickname: Nickname,
	}
	value, err := user.String()
	if err != nil {
		return err
	}
	return d.db.Set(ctx, tokenUri, value, constant.ApiHttpTokenExpire)
}

// Get 获取缓存
func (d *redisApiTokenLogin) Get(tokenUri string) (apiAuthDTO.ApiHttpUserCache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	user := apiAuthDTO.ApiHttpUserCache{}
	value, err := d.db.Get(ctx, tokenUri)
	if err != nil {
		return apiAuthDTO.ApiHttpUserCache{}, err
	}
	if err = user.Unmarshal(value); err != nil {
		return user, err
	}
	return user, nil
}
