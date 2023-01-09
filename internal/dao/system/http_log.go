/*
 * @Author: silent-rain
 * @Date: 2023-01-09 23:08:33
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-09 23:49:12
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/http_log.go
 * @Descripttion: 网络请求日志
 */
package systemDao

import (
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/database"
)

// HttpLogImpl 网络请求日志对象
var HttpLogImpl = new(httpLog)

// HttpLog 网络请求日志接口
type HttpLog interface {
	Add(bean *systemModel.HttpLog) (uint, error)
}

// 网络请求日志结构
type httpLog struct{}

// Add 添加网络请求日志
func (d *httpLog) Add(bean systemModel.HttpLog) (uint, error) {
	result := database.Instance().Create(&bean)
	if result.Error != nil {
		return 0, result.Error
	}
	return bean.ID, nil
}
