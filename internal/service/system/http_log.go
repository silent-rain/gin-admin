/*网络请求日志
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

// HttpLogService 网络请求日志接口
type HttpLogService interface {
	List(ctx *gin.Context, req systemDTO.QueryHttpLogReq) *response.ResponseAPI
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
func (s *httpLogService) List(ctx *gin.Context, req systemDTO.QueryHttpLogReq) *response.ResponseAPI {
	roles, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DBQueryError).Errorf("%v", err)
		return response.New().WithCode(statuscode.DBQueryError)

	}
	return response.New().WithDataList(roles, total)
}
