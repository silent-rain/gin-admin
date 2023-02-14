/*应用配置表*/
package service

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// ConfigService 配置接口
type ConfigService interface {
	AllTree(ctx *gin.Context) *response.ResponseAPI
	Tree(ctx *gin.Context, req systemDTO.QueryConfigReq) *response.ResponseAPI
	Add(ctx *gin.Context, menu systemModel.Config) *response.ResponseAPI
	Update(ctx *gin.Context, menu systemModel.Config) *response.ResponseAPI
	Delete(ctx *gin.Context, id uint) *response.ResponseAPI
	BatchDelete(ctx *gin.Context, ids []uint) *response.ResponseAPI
	Status(ctx *gin.Context, id uint, status uint) *response.ResponseAPI
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
func (s *configService) AllTree(ctx *gin.Context) *response.ResponseAPI {
	results, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	// 配置列表数据转为树结构
	tree := configListToTree(results, nil)
	return response.New().WithDataList(tree, int64(len(tree)))
}

// Tree 获取配置树
func (s *configService) Tree(ctx *gin.Context, req systemDTO.QueryConfigReq) *response.ResponseAPI {
	configList, _, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	configAll, _, err := s.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}

	// 配置列表数据转为树结构
	tree := configListToTree(configAll, nil)

	// 过滤
	treeFilter := make([]systemModel.Config, 0)
	for _, itemA := range tree {
		for _, item := range configList {
			if itemA.Key == item.Key {
				treeFilter = append(treeFilter, itemA)
			}
		}
	}
	return response.New().WithDataList(treeFilter, int64(len(tree)))
}

// Add 添加配置
func (s *configService) Add(ctx *gin.Context, config systemModel.Config) *response.ResponseAPI {
	if _, err := s.dao.Add(config); err != nil {
		log.New(ctx).WithCode(statuscode.DBAddError).Errorf("%v", err)

		return response.New().WithCode(statuscode.DBAddError)
	}
	return response.New()
}

// Update 更新配置
func (s *configService) Update(ctx *gin.Context, config systemModel.Config) *response.ResponseAPI {
	row, err := s.dao.Update(config)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBUpdateError)
	}
	return response.New().WithData(row)
}

// Delete 删除配置
func (s *configService) Delete(ctx *gin.Context, id uint) *response.ResponseAPI {
	childrenConfig, err := s.dao.Children(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)
	}
	if len(childrenConfig) > 0 {
		log.New(ctx).WithCode(statuscode.DBDataExistChildrenError).Errorf("删除失败, 存在子配置, %v", err)
		return response.New().WithCode(statuscode.DBDataExistChildrenError).WithMsg("删除失败, 存在子配置")
	}

	row, err := s.dao.Delete(id)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBDeleteError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBDeleteError)
	}
	return response.New().WithData(row)
}

// BatchDelete 批量删除配置, 批量删除，不校验是否存在子配置
func (s *configService) BatchDelete(ctx *gin.Context, ids []uint) *response.ResponseAPI {
	row, err := s.dao.BatchDelete(ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBBatchDeleteError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBBatchDeleteError)
	}
	return response.New().WithData(row)
}

// Status 更新配置状态
func (s *configService) Status(ctx *gin.Context, id uint, status uint) *response.ResponseAPI {
	row, err := s.dao.Status(id, status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBUpdateStatusError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBUpdateStatusError)
	}
	return response.New().WithData(row)
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
