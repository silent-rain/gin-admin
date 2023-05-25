// Package dto 分页, 标页码
package dto

// Pagination 分页查询条件
type Pagination struct {
	Page     int `json:"page" form:"page"`           // 页码，从1开始
	PageSize int `json:"page_size" form:"page_size"` // 每页数, 默认100
}

// GetPage 获取页数
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page - 1
}

// GetPageSize 获取页面数据大小
func (p *Pagination) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 100
	}
	return p.PageSize
}

// Offset 偏移行数
func (p *Pagination) Offset() int {
	return p.GetPage() * p.GetPageSize()
}
