/*菜单*/
package system

import (
	systemDao "gin-admin/internal/dao/system"
	"gin-admin/internal/dto"
	systemDto "gin-admin/internal/dto/system"
	systemModel "gin-admin/internal/model/system"
	"gin-admin/internal/pkg/log"
	"gin-admin/internal/pkg/response"
	statuscode "gin-admin/internal/pkg/status_code"
	"gin-admin/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 菜单
type menuHandler struct {
}

// 创建菜单 Handler 对象
func NewMenuHandler() *menuHandler {
	return &menuHandler{}
}

// All 获取所有菜单列表
func (h *menuHandler) All(ctx *gin.Context) {
	menus, total, err := systemDao.NewMenuDao().All()
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(menus, total).Json()
}

// List 获取用菜单列表
func (h *menuHandler) List(ctx *gin.Context) {
	req := new(systemDto.QueryMenuReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}

	menus, total, err := systemDao.NewMenuDao().List(*req)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbQueryError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbQueryError).Json()
		return
	}
	response.New(ctx).WithDataList(menus, total).Json()
}

// Add 添加菜单
func (h *menuHandler) Add(ctx *gin.Context) {
	req := new(systemDto.AddMenuReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	menu := systemModel.Menu{}
	if err := utils.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	if _, err := systemDao.NewMenuDao().Add(menu); err != nil {
		log.New(ctx).WithCode(statuscode.DbAddError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbAddError).Json()
		return
	}
	response.New(ctx).Json()
}

// Update 更新菜单
func (h *menuHandler) Update(ctx *gin.Context) {
	req := new(systemDto.UpdateMenuReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	menu := systemModel.Menu{}
	if err := utils.ApiJsonConvertJson(ctx, req, &menu); err != nil {
		log.New(ctx).WithField("data", req).Errorf("数据转换失败, %v", err)
		return
	}
	row, err := systemDao.NewMenuDao().Update(menu)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Delete 删除菜单
func (h *menuHandler) Delete(ctx *gin.Context) {
	req := new(dto.DeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := systemDao.NewMenuDao().Delete(req.ID)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// BatchDelete 批量删除菜单
func (h *menuHandler) BatchDelete(ctx *gin.Context) {
	req := new(dto.BatchDeleteReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := systemDao.NewMenuDao().BatchDelete(req.Ids)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbBatchDeleteError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbBatchDeleteError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}

// Status 更新菜单状态
func (h *menuHandler) Status(ctx *gin.Context) {
	req := new(dto.UpdateStatusReq)
	if err := utils.ParsingReqParams(ctx, req); err != nil {
		log.New(ctx).WithField("data", req).Errorf("参数解析失败, %v", err)
		return
	}
	row, err := systemDao.NewMenuDao().Status(req.ID, req.Status)
	if err != nil {
		log.New(ctx).WithCode(statuscode.DbUpdateStatusError).Errorf("%v", err)
		response.New(ctx).WithCode(statuscode.DbUpdateStatusError).Json()
		return
	}
	response.New(ctx).WithData(row).Json()
}
