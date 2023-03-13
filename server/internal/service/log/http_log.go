/*网络请求日志
 */
package log

import (
	logDAO "gin-admin/internal/dao/log"
	logDTO "gin-admin/internal/dto/log"
	logModel "gin-admin/internal/model/log"
	"gin-admin/internal/pkg/log"
	systemVO "gin-admin/internal/vo/system"
	"gin-admin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// HttpLogService 网络请求日志接口
type HttpLogService interface {
	List(ctx *gin.Context, req logDTO.QueryHttpLogReq) ([]logModel.HttpLog, int64, error)
	GetBody(ctx *gin.Context, id uint) (systemVO.QueryHttpLogBody, error)
}

// 网络请求日志
type httpLogService struct {
	dao logDAO.HttpLog
}

// NewHttpLogService 创建网络请求日志对象
func NewHttpLogService() *httpLogService {
	return &httpLogService{
		dao: logDAO.NewHttpLogDao(),
	}
}

// List 获取网络请求日志列表
func (s *httpLogService) List(ctx *gin.Context, req logDTO.QueryHttpLogReq) ([]logModel.HttpLog, int64, error) {
	results, total, err := s.dao.List(req)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return nil, 0, errcode.DBQueryError

	}
	return results, total, nil
}

// GetBody 获取 body 信息
func (h *httpLogService) GetBody(ctx *gin.Context, id uint) (systemVO.QueryHttpLogBody, error) {
	result := systemVO.QueryHttpLogBody{
		Body: "",
	}
	resp, ok, err := h.dao.GetBody(id)
	if err != nil {
		log.New(ctx).WithCode(errcode.DBQueryError).Errorf("%v", err)
		return result, errcode.DBQueryError
	}
	if !ok {
		log.New(ctx).WithCode(errcode.DBQueryEmptyError).Errorf("%v", err)
		return result, errcode.DBQueryEmptyError
	}

	result.Body = resp.Body
	return result, nil
}
