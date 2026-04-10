package siteTagController

import (
	"net/http"
	"strconv"

	siteTagService "gin-web-admin/app/service/v1/site_tag"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/query"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *siteTagService.Service
}

func NewHandler(service *siteTagService.Service) *Handler {
	if service == nil {
		panic("site tag handler requires non-nil service")
	}
	return &Handler{service: service}
}

func (h *Handler) GetSiteTagList(c *gin.Context) {
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

	result, err := h.service.GetSiteTagList(siteTagService.SiteTagQuery{
		Pagination: pg.Clone(),
		Keyword:    c.Query("keyword"),
		Status:     status,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取标签列表失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) GetAllSiteTags(c *gin.Context) {
	appG := common.Gin{C: c}

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

	list, err := h.service.GetAllSiteTags(status)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取全部标签失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", list)
}

func (h *Handler) CreateSiteTag(c *gin.Context) {
	appG := common.Gin{C: c}

	var payload siteTagService.CreateSiteTagStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := h.service.CreateSiteTag(payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "创建成功", nil)
}

func (h *Handler) UpdateSiteTag(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	var payload siteTagService.UpdateSiteTagStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := h.service.UpdateSiteTag(id, payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "更新成功", nil)
}

func (h *Handler) DeleteSiteTag(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	if err := h.service.DeleteSiteTag(id); err != nil {
		if err.Error() == "tag in use" {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "该标签已关联内容，无法删除", nil)
			return
		}
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}
