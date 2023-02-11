/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:55:36
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 17:14:09
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/system/role.go
 * @Descripttion: 角色
 */
package system

import (
	systemDAO "gin-admin/internal/dao/system"
	DTO "gin-admin/internal/dto"
	systemDTO "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 角色
type roleHandler struct {
	dao systemDAO.Role
}

// 创建角色 Handler 对象
func NewRoleHandler() *roleHandler {
	return &roleHandler{
		dao: systemDAO.NewRoleDao(),
	}
}

// All 获取所有角色列表
func (h *roleHandler) All(ctx *gin.Context) {
	roles, total, err := h.dao.All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(roles, total).Json()
}

// List 获取用角色列表
func (h *roleHandler) List(ctx *gin.Context) {
	req := new(systemDTO.QueryRoleReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	roles, total, err := h.dao.List(*req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(roles, total).Json()
}

// Add 添加角色
func (h *roleHandler) Add(ctx *gin.Context) {
	req := new(systemDTO.AddRoleReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	role := systemModel.Role{}
	if err := utils.ApiJsonConvertJson(ctx, req, &role); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}

	_, ok, err := h.dao.InfoByName(role.Name)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	if ok {
		log.New(ctx).WithCode(statuscode.DbDataExistError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbDataExistError).WithMsg("角色已存在").Json()
		return
	}

	if _, err := h.dao.Add(role); err != nil {
		log.New(ctx).WithCode(statuscode.DbAddError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbAddError).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新角色
func (h *roleHandler) Update(ctx *gin.Context) {
	req := new(systemDTO.UpdateRoleReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	role := systemModel.Role{}
	if err := utils.ApiJsonConvertJson(ctx, req, &role); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	row, err := h.dao.Update(role)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Delete 删除角色
func (h *roleHandler) Delete(ctx *gin.Context) {
	req := new(DTO.DeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.Delete(req.ID)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// BatchDelete 批量删除角色
func (h *roleHandler) BatchDelete(ctx *gin.Context) {
	req := new(DTO.BatchDeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.BatchDelete(req.Ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbBatchDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbBatchDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新角色状态
func (h *roleHandler) Status(ctx *gin.Context) {
	req := new(DTO.UpdateStatusReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := h.dao.Status(req.ID, req.Status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}
