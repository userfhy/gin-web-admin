package dictDataController

import (
	"net/http"
	"strconv"

	dictDataService "gin-web-admin/app/service/v1/dict_data"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/query"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *dictDataService.Service
}

func NewHandler(service *dictDataService.Service) *Handler {
	if service == nil {
		panic("dict data handler requires non-nil service")
	}
	return &Handler{service: service}
}

func (h *Handler) GetDictDataList(c *gin.Context) {
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

	result, err := h.service.GetDictDataList(dictDataService.DictDataQuery{
		Pagination: pg.Clone(),
		DictType:   c.Query("dictType"),
		Label:      c.Query("label"),
		Status:     status,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取字典数据列表失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) CreateDictData(c *gin.Context) {
	appG := common.Gin{C: c}

	var payload dictDataService.CreateDictDataStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	if err := h.service.CreateDictData(payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "创建成功", nil)
}

func (h *Handler) UpdateDictData(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	var payload dictDataService.UpdateDictDataStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	if err := h.service.UpdateDictData(id, payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "更新成功", nil)
}

func (h *Handler) DeleteDictData(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}
	if err := h.service.DeleteDictData(id); err != nil {
		if err.Error() == "dict data not found" {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "字典数据不存在", nil)
			return
		}
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}
