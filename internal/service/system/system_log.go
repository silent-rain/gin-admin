/*系统日志
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

// SystemLogService 系统日志接口
type SystemLogService interface {
	List(ctx *gin.Context, req systemDTO.QuerySystemLogReq) ([]systemModel.SystemLog, int64, error)
}

// 系统日志
type systemLogService struct {
	dao systemDAO.SystemLog
}

// NewSystemLogService 创建系统日志对象
func NewSystemLogService() *systemLogService {
	return &systemLogService{
		dao: systemDAO.NewSystemLogDao(),
	}
}

// List 获取系统日志列表
func (s *systemLogService) List(ctx *gin.Context, req systemDTO.QuerySystemLogReq) ([]systemModel.SystemLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.New(errcode.DBQueryError)

	}
	return results, total, nil
}
