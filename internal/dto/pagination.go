/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:42:08
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 00:45:20
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/pagination.go
 * @Descripttion: 分页, 标页码
 */
package dto

// Pagination 分页查询条件
type Pagination struct {
	Page  int `json:"page" form:"page"`   // 页码
	Limit int `json:"limit" form:"limit"` // 每页数
}
