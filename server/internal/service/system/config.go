/*应用配置表*/
package system

import (
	"errors"

	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ConfigService 配置接口
type ConfigService interface {
	AllTree(ctx *gin.Context) ([]systemModel.Config, int64, error)
	Tree(ctx *gin.Context, req systemDTO.QueryConfigReq) ([]systemModel.Config, int64, error)
	Add(ctx *gin.Context, menu systemModel.Config) (uint, error)
	Update(ctx *gin.Context, menu systemModel.Config) (int64, error)
	Delete(ctx *gin.Context, id uint) (int64, error)
	BatchDelete(ctx *gin.Context, ids []uint) (int64, error)
	Status(ctx *gin.Context, id uint, status uint) (int64, error)
}

// 配置
type configService struct {
	dao systemDAO.Config
}

// NewConfigService 创建配置对象
func NewConfigService() *configService {
	return &configService{
		dao: systemDAO.NewConfigDao(),
	}
}

// AllTree 获取所有配置树
func (s *configService) AllTree(ctx *gin.Context) ([]systemModel.Config, int64, error) {
	results, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}
	// 配置列表数据转为树结构
	tree := configListToTree(results, nil)
	return tree, int64(len(tree)), nil
}

// Tree 获取配置树
func (s *configService) Tree(ctx *gin.Context, req systemDTO.QueryConfigReq) ([]systemModel.Config, int64, error) {
	configList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}
	configAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)
	}

	// 配置列表数据转为树结构
	tree := configListToTree(configAll, nil)

	// 过滤
	treeFilter := make([]systemModel.Config, 0)
	for _, item := range configList {
		for _, itemA := range tree {
			if item.Key == itemA.Key {
				treeFilter = append(treeFilter, itemA)
			}
		}
	}
	return treeFilter, int64(len(tree)), nil
}

// Add 添加配置
func (s *configService) Add(ctx *gin.Context, config systemModel.Config) (uint, error) {
	_, ok, err := s.dao.InfoByKey(config.Key)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBQueryError)
	}
	if ok {
		log.New(ctx).WithCode(errcode.DBDataExistError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBDataExistError).WithMsg("配置项已存在")
	}

	id, err := s.dao.Add(config)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}

// Update 更新配置
func (s *configService) Update(ctx *gin.Context, config systemModel.Config) (int64, error) {
	row, err := s.dao.Update(config)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateError)
	}
	return row, nil
}

// Delete 删除配置
func (s *configService) Delete(ctx *gin.Context, id uint) (int64, error) {
	childrenConfig, err := s.dao.Children(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBQueryError)
	}
	if len(childrenConfig) > 0 {
		log.New(ctx).WithCode(errcode.DBDataExistChildrenError).Errorf("删除失败, 存在子配置, %v", err)
		return 0, errcode.New(errcode.DBDataExistChildrenError).WithMsg("删除失败, 存在子配置")
	}

	row, err := s.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBDeleteError)
	}
	return row, nil
}

// BatchDelete 批量删除配置, 批量删除，不校验是否存在子配置
func (s *configService) BatchDelete(ctx *gin.Context, ids []uint) (int64, error) {
	row, err := s.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBBatchDeleteError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBBatchDeleteError)
	}
	return row, nil
}

// Status 更新配置状态
func (s *configService) Status(ctx *gin.Context, id uint, status uint) (int64, error) {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBUpdateStatusError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBUpdateStatusError)
	}
	return row, nil
}

// 配置列表数据转为树结构
func configListToTree(src []systemModel.Config, parentId *uint) []systemModel.Config {
	tree := make([]systemModel.Config, 0)
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
