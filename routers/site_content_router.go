package routers

import (
	siteContentController "gin-web-admin/app/controllers/v1/site_content"

	"github.com/gin-gonic/gin"
)

func InitSiteContentRouter(router *gin.RouterGroup, handler *siteContentController.Handler) {
	siteContent := router.Group("/site-content")
	{
		siteContent.GET("", handler.GetSiteContentList)
		siteContent.GET("/:id", handler.GetSiteContent)
		siteContent.POST("", handler.CreateSiteContent)
		siteContent.PUT("/:id", handler.UpdateSiteContent)
		siteContent.PATCH("/:id/status", handler.UpdateSiteContentStatus)
		siteContent.DELETE("/:id", handler.DeleteSiteContent)
	}
}
