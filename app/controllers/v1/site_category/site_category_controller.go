package siteCategoryController

import (
	"net/http"
	"strconv"

	siteCategoryService "gin-web-admin/app/service/v1/site_category"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/query"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *siteCategoryService.Service
}

func NewHandler(service *siteCategoryService.Service) *Handler {
	if service == nil {
		panic("site category handler requires non-nil service")
	}
	return &Handler{service: service}
}

func (h *Handler) GetSiteCategoryList(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c, utils.WithMaxPageSize(100))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	filter := query.NewBuilder()
	if err := filter.FromQuery(c, query.RuleSet{
		"status": {Field: "status", Op: query.OpEqual, Parser: query.IntEnumParser(0, 1)},
	}); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	var status *int
	if v, ok := filter.Build()["status ="]; ok {
		val := v.(int)
		status = &val
	}

	result, err := h.service.GetSiteCategoryList(siteCategoryService.SiteCategoryQuery{
		Pagination: pg.Clone(),
		Keyword:    c.Query("keyword"),
		Status:     status,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取分类列表失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) GetAllSiteCategories(c *gin.Context) {
	appG := common.Gin{C: c}

	var status *int
	if statusStr := c.Query("status"); statusStr != "" {
		statusInt, err := strconv.Atoi(statusStr)
		if err != nil || (statusInt != 0 && statusInt != 1) {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "status 参数无效，只能为 0 或 1", nil)
			return
		}
		status = &statusInt
	}

	list, err := h.service.GetAllSiteCategories(status)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取全部分类失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", list)
}

func (h *Handler) CreateSiteCategory(c *gin.Context) {
	appG := common.Gin{C: c}

	var payload siteCategoryService.CreateSiteCategoryStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := h.service.CreateSiteCategory(payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "创建成功", nil)
}

func (h *Handler) UpdateSiteCategory(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	var payload siteCategoryService.UpdateSiteCategoryStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := h.service.UpdateSiteCategory(id, payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "更新成功", nil)
}

func (h *Handler) DeleteSiteCategory(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	if err := h.service.DeleteSiteCategory(id); err != nil {
		if err.Error() == "category in use" {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "该分类已关联内容，无法删除", nil)
			return
		}
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}
