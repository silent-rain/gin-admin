/*网络请求日志 DAO
 */
package systemDAO

import (
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/repository/mysql"

	"gorm.io/gorm"
)

// HttpLog 网络请求日志接口
type HttpLog interface {
	List(req systemDTO.QueryHttpLogReq) ([]systemModel.HttpLog, int64, error)
	Add(bean systemModel.HttpLog) (uint, error)
}

// 网络请求日志结构
type httpLog struct {
	db mysql.DBRepo
}

// 创建网络请求日志 Dao 对象
func NewHttpLogDao() *httpLog {
	return &httpLog{
		db: mysql.Instance(),
	}
}

// List 查询网络请求日志列表
func (d *httpLog) List(req systemDTO.QueryHttpLogReq) ([]systemModel.HttpLog, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db.GetDbR()
		if req.UserId != 0 {
			stats = stats.Where("user_id = ?", req.UserId)
		}
		if req.TraceId != "" {
			stats = stats.Where("trace_id = ? OR parent_trace_id = ?", req.TraceId, req.TraceId)
		}
		if req.StatusCode != 0 {
			stats = stats.Where("status_code = ?", req.StatusCode)
		}
		if req.Method != "" {
			stats = stats.Where("method = ?", req.Method)
		}
		if req.RemoteAddr != "" {
			stats = stats.Where("remote_addr = ?", req.RemoteAddr)
		}
		if req.HttpType != "" {
			stats = stats.Where("htpp_type = ?", req.HttpType)
		}
		if req.Path != "" {
			stats = stats.Where("path LIKE ?", "%"+req.Path+"%")
		}
		return stats
	}

	beans := make([]systemModel.HttpLog, 0)
	result := stats().Offset(req.Offset()).Limit(req.PageSize).
		Order("created_at DESC").
		Find(&beans)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.HttpLog{}).Count(&total)
	return beans, total, nil
}

// Add 添加网络请求日志
func (d *httpLog) Add(bean systemModel.HttpLog) (uint, error) {
	result := d.db.GetDbW().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}