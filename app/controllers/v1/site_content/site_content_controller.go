package siteContentController

import (
	"net/http"
	"strconv"

	siteContentService "gin-web-admin/app/service/v1/site_content"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/query"

	"github.com/gin-gonic/gin"
)

func GetSiteContentList(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c, utils.WithMaxPageSize(100))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	filter := query.NewBuilder()
	if err := filter.FromQuery(c, query.RuleSet{
		"status":     {Field: "status", Op: query.OpEqual, Parser: query.IntEnumParser(0, 1)},
		"categoryId": {Field: "category_id", Op: query.OpEqual, Parser: query.IntParser()},
		"tagId":      {Field: "tag_id", Op: query.OpEqual, Parser: query.IntParser()},
	}); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	var status *int
	if v, ok := filter.Build()["status ="]; ok {
		val := v.(int)
		status = &val
	}
	var categoryID *int
	if v, ok := filter.Build()["category_id ="]; ok {
		val := v.(int)
		categoryID = &val
	}
	var tagID *int
	if v, ok := filter.Build()["tag_id ="]; ok {
		val := v.(int)
		tagID = &val
	}

	result, err := siteContentService.GetSiteContentList(siteContentService.SiteContentQuery{
		Pagination: pg.Clone(),
		Keyword:    c.Query("keyword"),
		Status:     status,
		CategoryID: categoryID,
		TagID:      tagID,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取官网内容列表失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func GetSiteContent(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	data, err := siteContentService.GetSiteContentDetail(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取官网内容详情失败", nil)
		return
	}
	if data == nil {
		appG.Response(http.StatusOK, code.ERROR, "内容不存在", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}

func CreateSiteContent(c *gin.Context) {
	appG := common.Gin{C: c}

	var payload siteContentService.CreateSiteContentStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := siteContentService.CreateSiteContent(payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "创建成功", nil)
}

func UpdateSiteContent(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	var payload siteContentService.UpdateSiteContentStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := siteContentService.UpdateSiteContent(id, payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "更新成功", nil)
}

func DeleteSiteContent(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	if err := siteContentService.DeleteSiteContent(id); err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}

func UpdateSiteContentStatus(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	var payload siteContentService.UpdateSiteContentStatusStruct
	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := siteContentService.UpdateSiteContentStatus(id, payload); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "状态更新成功", nil)
}
