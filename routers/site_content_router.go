package routers

import (
	siteContentController "gin-web-admin/app/controllers/v1/site_content"
	"gin-web-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitSiteContentRouter(Router *gin.RouterGroup) {
	siteContent := Router.Group("/site-content").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler(),
	)
	{
		siteContent.GET("", siteContentController.GetSiteContentList)
		siteContent.GET("/:id", siteContentController.GetSiteContent)
		siteContent.POST("", siteContentController.CreateSiteContent)
		siteContent.PUT("/:id", siteContentController.UpdateSiteContent)
		siteContent.PATCH("/:id/status", siteContentController.UpdateSiteContentStatus)
		siteContent.DELETE("/:id", siteContentController.DeleteSiteContent)
	}
}
