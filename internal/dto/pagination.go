/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:42:08
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 22:21:26
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/pagination.go
 * @Descripttion: 分页, 标页码
 */
package dto

// Pagination 分页查询条件
type Pagination struct {
	Page     int `json:"page" form:"page"`           // 页码，从1开始
	PageSize int `json:"page_size" form:"page_size"` // 每页数, 默认100
}

// 获取页数
func (p *Pagination) getPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page - 1
}

// Offset 偏移行数
func (p *Pagination) Offset() int {
	if p.PageSize == 0 {
		p.PageSize = 100
	}
	return p.getPage() * p.PageSize
}
