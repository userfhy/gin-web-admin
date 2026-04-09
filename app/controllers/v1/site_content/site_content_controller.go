package siteContentController

import (
	"net/http"
	"strconv"

	siteContentService "gin-web-admin/app/service/v1/site_content"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
)

func GetSiteContentList(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c, utils.WithMaxPageSize(100))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	var status *int
	if statusStr := c.Query("status"); statusStr != "" {
		statusInt, err := strconv.Atoi(statusStr)
		if err != nil || (statusInt != 0 && statusInt != 1) {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "status 参数无效，只能为 0 或 1", nil)
			return
		}
		status = &statusInt
	}
	var categoryID *int
	if categoryStr := c.Query("categoryId"); categoryStr != "" {
		categoryInt, err := strconv.Atoi(categoryStr)
		if err != nil || categoryInt <= 0 {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "categoryId 参数无效", nil)
			return
		}
		categoryID = &categoryInt
	}
	var tagID *int
	if tagStr := c.Query("tagId"); tagStr != "" {
		tagInt, err := strconv.Atoi(tagStr)
		if err != nil || tagInt <= 0 {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "tagId 参数无效", nil)
			return
		}
		tagID = &tagInt
	}

	data, err := siteContentService.GetSiteContentList(siteContentService.SiteContentQuery{
		PageNum:    pg.Page,
		PageSize:   pg.PageSize,
		Keyword:    c.Query("keyword"),
		Status:     status,
		CategoryID: categoryID,
		TagID:      tagID,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取官网内容列表失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
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
