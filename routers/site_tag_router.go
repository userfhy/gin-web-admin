package routers

import (
	siteTagController "gin-web-admin/app/controllers/v1/site_tag"

	"github.com/gin-gonic/gin"
)

func InitSiteTagRouter(router *gin.RouterGroup, handler *siteTagController.Handler) {
	siteTag := router.Group("/site-tag")
	{
		siteTag.GET("", handler.GetSiteTagList)
		siteTag.GET("/all", handler.GetAllSiteTags)
		siteTag.POST("", handler.CreateSiteTag)
		siteTag.PUT("/:id", handler.UpdateSiteTag)
		siteTag.DELETE("/:id", handler.DeleteSiteTag)
	}
}
