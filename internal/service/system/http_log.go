/*网络请求日志
 */
package service

import (
	systemDAO "gin-admin/internal/dao/system"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/code_errors"
	"gin-admin/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// HttpLogService 网络请求日志接口
type HttpLogService interface {
	List(ctx *gin.Context, req systemDTO.QueryHttpLogReq) ([]systemModel.HttpLog, int64, error)
}

// 网络请求日志
type httpLogService struct {
	dao systemDAO.HttpLog
}

// NewHttpLogService 创建网络请求日志对象
func NewHttpLogService() *httpLogService {
	return &httpLogService{
		dao: systemDAO.NewHttpLogDao(),
	}
}

// List 获取网络请求日志列表
func (s *httpLogService) List(ctx *gin.Context, req systemDTO.QueryHttpLogReq) ([]systemModel.HttpLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(code_errors.DBQueryError).Errorf("%v", err)
		return nil, 0, code_errors.New(code_errors.DBQueryError)

	}
	return results, total, nil
}
