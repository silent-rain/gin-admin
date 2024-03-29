// Package dto 网络请求日志
package dto

import (
	DTO "github.com/silent-rain/gin-admin/internal/dto"
)

// QueryHttpLogReq 查询条件
type QueryHttpLogReq struct {
	DTO.Pagination
	UserId     uint   `json:"user_id" form:"user_id"`
	TraceId    string `json:"trace_id" form:"trace_id"`
	StatusCode int    `json:"status_code" form:"status_code"`
	Method     string `json:"method" form:"method"`
	Path       string `json:"path" form:"path"`
	RemoteAddr string `json:"remote_addr" form:"remote_addr"`
	HttpType   string `json:"htpp_type" form:"htpp_type"`
}

// QueryHttpLogBodyReq 查询条件 网络请求日志的 body 信息
type QueryHttpLogBodyReq struct {
	Id uint `json:"user_id" form:"id" binding:"required"`
}

// QueryHttpLogBody 网络请求日志的 body 信息
type QueryHttpLogBody struct {
	Body string `json:"body"`
}
