/*系统日志 DAO
 */
package systemDAO

import (
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"

	"gorm.io/gorm"
)

// HttpLog 系统日志接口
type SystemLog interface {
	List(req systemDTO.QuerySystemLogReq) ([]systemModel.SystemLog, int64, error)
	Add(bean systemModel.SystemLog) (uint, error)
}

// 系统日志
type systemLog struct {
	db *gorm.DB
}

// 创建系统日志对象
func NewSystemLogDao() *systemLog {
	return &systemLog{
		db: database.Instance(),
	}
}

// List 查询系统日志列表
func (d *systemLog) List(req systemDTO.QuerySystemLogReq) ([]systemModel.SystemLog, int64, error) {
	var stats = func() *gorm.DB {
		stats := d.db
		if req.UserId != 0 {
			stats = stats.Where("user_id = ?", req.UserId)
		}
		if req.TraceId != "" {
			stats = stats.Where("trace_id = ?", req.TraceId)
		}
		if req.Level != "" {
			stats = stats.Where("level = ?", req.Level)
		}
		if req.ErrorCode != 0 {
			stats = stats.Where("error_code = ?", req.ErrorCode)
		}
		if req.ErrorMsg != "" {
			stats = stats.Where("error_msg LIKE ?", req.ErrorMsg+"%")
		}
		if req.Msg != "" {
			stats = stats.Where("msg LIKE ?", req.Msg+"%")
		}
		return stats
	}

	beans := make([]systemModel.SystemLog, 0)
	result := stats().Offset(req.Offset()).Limit(req.PageSize).
		Order("updated_at DESC").
		Find(&beans)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64 = 0
	stats().Model(&systemModel.SystemLog{}).Count(&total)
	return beans, total, nil
}

// Add 添加系统日志
func (d *systemLog) Add(bean systemModel.SystemLog) (uint, error) {
	result := d.db.Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}
