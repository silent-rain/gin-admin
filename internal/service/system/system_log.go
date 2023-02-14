/*系统日志
 */
package service

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
)

// SystemLogService 系统日志接口
type SystemLogService interface {
	List(ctx *gin.Context, req systemDTO.QuerySystemLogReq) *response.ResponseAPI
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
func (s *systemLogService) List(ctx *gin.Context, req systemDTO.QuerySystemLogReq) *response.ResponseAPI {
	roles, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)

	}
	return response.New().WithDataList(roles, total)
}
