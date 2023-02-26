/*WEB 日志
 */
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// WebLogService WEB 日志接口
type WebLogService interface {
	List(ctx *gin.Context, req systemDTO.QueryWebLogReq) ([]systemModel.WebLog, int64, error)
	Add(ctx *gin.Context, bean systemModel.WebLog) (uint, error)
}

// WEB 日志
type webLogService struct {
	dao systemDAO.WebLog
}

// NewWebLogService 创建 WEB 日志对象
func NewWebLogService() *webLogService {
	return &webLogService{
		dao: systemDAO.NewWebLogDao(),
	}
}

// List 获取 WEB 日志列表
func (s *webLogService) List(ctx *gin.Context, req systemDTO.QueryWebLogReq) ([]systemModel.WebLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return results, total, nil
}

// Add 添加 WEB 日志
func (h *webLogService) Add(ctx *gin.Context, bean systemModel.WebLog) (uint, error) {
	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.New(errcode.DBAddError)
	}
	return id, nil
}
