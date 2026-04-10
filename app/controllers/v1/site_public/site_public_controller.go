package sitePublicController

import (
	"net/http"
	"strconv"

	sitePublicService "gin-web-admin/app/service/v1/site_public"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
)

func GetPublicCategories(c *gin.Context) {
	appG := common.Gin{C: c}
	list, err := sitePublicService.GetPublicCategories()
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取分类失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", list)
}

func GetPublicTags(c *gin.Context) {
	appG := common.Gin{C: c}
	list, err := sitePublicService.GetPublicTags()
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取标签失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", list)
}

func GetPublicContentList(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c, utils.WithMaxPageSize(100))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	data, err := sitePublicService.GetPublicContentList(sitePublicService.PublicContentQuery{
		Pagination:   pg.Clone(),
		Keyword:      c.Query("keyword"),
		CategorySlug: c.Query("categorySlug"),
		TagSlug:      c.Query("tagSlug"),
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取内容列表失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}

func GetPublicContentDetail(c *gin.Context) {
	appG := common.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "id 参数无效", nil)
		return
	}
	data, err := sitePublicService.GetPublicContentDetailByID(id)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	if data == nil {
		appG.Response(http.StatusOK, code.ERROR, "内容不存在", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}

func GetPublicContentDetailBySlug(c *gin.Context) {
	appG := common.Gin{C: c}
	data, err := sitePublicService.GetPublicContentDetailBySlug(c.Param("slug"))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	if data == nil {
		appG.Response(http.StatusOK, code.ERROR, "内容不存在", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}
