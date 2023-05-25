/*应用配置表*/
package service

import (
	"errors"

	"github.com/silent-rain/gin-admin/internal/app/system/dao"
	"github.com/silent-rain/gin-admin/internal/app/system/dto"
	"github.com/silent-rain/gin-admin/internal/app/system/model"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ConfigService 配置接口
type ConfigService interface {
	AllTree(ctx *gin.Context) ([]model.Config, int64, error)
	Tree(ctx *gin.Context, req dto.QueryConfigReq) ([]model.Config, int64, error)
	List(ctx *gin.Context, req dto.QueryConfigReq) ([]model.Config, int64, error)
	Info(ctx *gin.Context, key string) (model.Config, error)
	Add(ctx *gin.Context, config model.Config) (uint, error)
	Update(ctx *gin.Context, config model.Config) (int64, error)
	BatchUpdate(ctx *gin.Context, configs []model.Config) error
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
	ChildrenByKey(ctx *gin.Context, key string) ([]model.Config, error)
	WebSiteConfigList(ctx *gin.Context, key string) ([]model.Config, error)
}

// 配置
type configService struct {
	dao        dao.Config
	innerCache dao.WebSiteConfigCache
}

// NewConfigService 创建配置对象
func NewConfigService() *configService {
	return &configService{
		dao:        dao.NewConfigDao(),
		innerCache: dao.NewWebSiteConfigCache(),
	}
}

// AllTree 获取所有配置树
func (s *configService) AllTree(ctx *gin.Context) ([]model.Config, int64, error) {
	results, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	// 配置列表数据转为树结构
	tree := configListToTree(results, nil)
	return tree, int64(len(tree)), nil
}

// Tree 获取配置树
func (s *configService) Tree(ctx *gin.Context, req dto.QueryConfigReq) ([]model.Config, int64, error) {
	configList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	configAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}

	// 配置列表数据转为树结构
	tree := configListToTree(configAll, nil)

	// 过滤
	treeFilter := make([]model.Config, 0)
	for _, item := range configList {
		for _, itemA := range tree {
			if item.Key == itemA.Key {
				treeFilter = append(treeFilter, itemA)
			}
		}
	}
	return treeFilter, int64(len(tree)), nil
}

// List 获取配置列表
func (s *configService) List(ctx *gin.Context, req dto.QueryConfigReq) ([]model.Config, int64, error) {
	configList, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError
	}
	return configList, total, nil
}

// Info 获取配置信息
func (s *configService) Info(ctx *gin.Context, key string) (model.Config, error) {
	result, ok, err := s.dao.Info(key)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return result, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Errorf("%v", err)
		return result, errcode.DBQueryEmptyError
	}
	return result, nil
}

// ChildrenByKey 通过父 key 获取子配置列表
func (s *configService) ChildrenByKey(ctx *gin.Context, key string) ([]model.Config, error) {
	configList, err := s.dao.ChildrenByKey(key)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, errcode.DBQueryError
	}
	return configList, nil
}

// Add 添加配置
func (s *configService) Add(ctx *gin.Context, config model.Config) (uint, error) {
	_, ok, err := s.dao.Info(config.Key)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("配置项已存在")
		return 0, errcode.DBDataExistError.WithMsg("配置项已存在")
	}

	id, err := s.dao.Add(config)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}

	// 设置站点配置缓存
	s.innerCache.Set()
	return id, nil
}

// Update 更新配置
func (s *configService) Update(ctx *gin.Context, config model.Config) (int64, error) {
	row, err := s.dao.Update(config)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.DBUpdateError
	}

	// 设置站点配置缓存
	s.innerCache.Set()
	return row, nil
}

// BatchUpdate 批量更新配置
func (s *configService) BatchUpdate(ctx *gin.Context, configs []model.Config) error {
	err := s.dao.BatchUpdate(configs)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return errcode.DBUpdateError
	}

	// 设置站点配置缓存
	s.innerCache.Set()
	return nil
}

// Delete 删除配置
func (s *configService) Delete(ctx *gin.Context, id uint) (int64, error) {
	childrenConfig, err := s.dao.Children(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.DBQueryError
	}
	if len(childrenConfig) > 0 {
		log.New(ctx).WithCode(errcode.DBDataExistChildrenError).Errorf("删除失败, 存在子配置, %v", err)
		return 0, errcode.DBDataExistChildrenError.WithMsg("删除失败, 存在子配置")
	}

	row, err := s.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.DBDeleteError
	}

	// 设置站点配置缓存
	s.innerCache.Set()
	return row, nil
}

// BatchDelete 批量删除配置, 批量删除，不校验是否存在子配置
func (s *configService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := s.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.DBBatchDeleteError
	}

	// 设置站点配置缓存
	s.innerCache.Set()
	return row, nil
}

// Status 更新配置状态
func (s *configService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.DBUpdateStatusError
	}

	// 设置站点配置缓存
	s.innerCache.Set()
	return row, nil
}

// 配置列表数据转为树结构
func configListToTree(src []model.Config, parentId *uint) []model.Config {
	tree := make([]model.Config, 0)
	for _, item := range src {
		if (item.ParentId == nil && parentId == nil) ||
			(item.ParentId != nil && parentId != nil && *item.ParentId == *parentId) {
			tree = append(tree, item)
		}
	}

	for i := range tree {
		children := configListToTree(src, &tree[i].ID)
		if tree[i].Children == nil {
			tree[i].Children = children
		} else {
			tree[i].Children = append(tree[i].Children, children...)
		}
	}
	return tree
}

// WebSiteConfigList 查询网站配置列表
func (s *configService) WebSiteConfigList(ctx *gin.Context, key string) ([]model.Config, error) {
	results, err := s.innerCache.Get()
	if err != nil {
		return nil, err
	}
	return results, nil
}
