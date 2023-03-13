/*系统日志
 */
package log

import (
	logDAO "github.com/silent-rain/gin-admin/internal/dao/log"
	logDTO "github.com/silent-rain/gin-admin/internal/dto/log"
	logModel "github.com/silent-rain/gin-admin/internal/model/log"
	"github.com/silent-rain/gin-admin/internal/pkg/log"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// SystemLogService 系统日志接口
type SystemLogService interface {
	List(ctx *gin.Context, req logDTO.QuerySystemLogReq) ([]logModel.SystemLog, int64, error)
}

// 系统日志
type systemLogService struct {
	dao logDAO.SystemLog
}

// NewSystemLogService 创建系统日志对象
func NewSystemLogService() *systemLogService {
	return &systemLogService{
		dao: logDAO.NewSystemLogDao(),
	}
}

// List 获取系统日志列表
func (s *systemLogService) List(ctx *gin.Context, req logDTO.QuerySystemLogReq) ([]logModel.SystemLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}
