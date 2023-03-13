/*WEB 日志
 */
package log

import (
	logDAO "gin-admin/internal/dao/log"
	logDTO "gin-admin/internal/dto/log"
	logModel "gin-admin/internal/model/log"
	"gin-admin/internal/pkg/log"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// WebLogService WEB 日志接口
type WebLogService interface {
	List(ctx *gin.Context, req logDTO.QueryWebLogReq) ([]logModel.WebLog, int64, error)
	Add(ctx *gin.Context, bean logModel.WebLog) (uint, error)
}

// WEB 日志
type webLogService struct {
	dao logDAO.WebLog
}

// NewWebLogService 创建 WEB 日志对象
func NewWebLogService() *webLogService {
	return &webLogService{
		dao: logDAO.NewWebLogDao(),
	}
}

// List 获取 WEB 日志列表
func (s *webLogService) List(ctx *gin.Context, req logDTO.QueryWebLogReq) ([]logModel.WebLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}

// Add 添加 WEB 日志
func (h *webLogService) Add(ctx *gin.Context, bean logModel.WebLog) (uint, error) {
	id, err := h.dao.Add(bean)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBAddError).Errorf("%v", err)
		return 0, errcode.DBAddError
	}
	return id, nil
}
