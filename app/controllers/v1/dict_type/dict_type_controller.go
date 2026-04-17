package dictTypeController

import (
	"net/http"
	"strconv"

	dictTypeService "gin-web-admin/app/service/v1/dict_type"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/query"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *dictTypeService.Service
}

func NewHandler(service *dictTypeService.Service) *Handler {
	if service == nil {
		panic("dict type handler requires non-nil service")
	}
	return &Handler{service: service}
}

func (h *Handler) GetDictTypeList(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c, utils.WithMaxPageSize(100))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	filter := query.NewBuilder()
	if err := filter.FromQuery(c, query.RuleSet{
		"status": query.Rule{Field: "status", Op: query.OpEqual, Parser: query.IntEnumParser(0, 1)},
	}); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	var status *int
	if v, ok := filter.Build()["status ="]; ok {
		val := v.(int)
		status = &val
	}

	result, err := h.service.GetDictTypeList(dictTypeService.DictTypeQuery{
		Pagination: pg.Clone(),
		Name:       c.Query("name"),
		Type:       c.Query("type"),
		Status:     status,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取字典类型列表失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) GetAllDictTypes(c *gin.Context) {
	appG := common.Gin{C: c}

	filter := query.NewBuilder()
	if err := filter.FromQuery(c, query.RuleSet{
		"status": query.Rule{Field: "status", Op: query.OpEqual, Parser: query.IntEnumParser(0, 1)},
	}); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	var status *int
	if v, ok := filter.Build()["status ="]; ok {
		val := v.(int)
		status = &val
	}

	list, err := h.service.GetAllDictTypes(status)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取全部字典类型失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", list)
}

func (h *Handler) CreateDictType(c *gin.Context) {
	appG := common.Gin{C: c}

	var payload dictTypeService.CreateDictTypeStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	if err := h.service.CreateDictType(payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "创建成功", nil)
}

func (h *Handler) RefreshDictCache(c *gin.Context) {
	appG := common.Gin{C: c}

	if err := h.service.RefreshDictCache(); err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "刷新字典缓存失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "刷新字典缓存成功", nil)
}

func (h *Handler) UpdateDictType(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	var payload dictTypeService.UpdateDictTypeStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	if err := h.service.UpdateDictType(id, payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "更新成功", nil)
}

func (h *Handler) DeleteDictType(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	if err := h.service.DeleteDictType(id); err != nil {
		if err.Error() == "dict type in use" {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "该字典类型下存在数据，无法删除", nil)
			return
		}
		if err.Error() == "dict type not found" {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "字典类型不存在", nil)
			return
		}
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}
