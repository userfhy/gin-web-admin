package routers

import (
	siteTagController "gin-web-admin/app/controllers/v1/site_tag"
	"gin-web-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitSiteTagRouter(Router *gin.RouterGroup) {
	siteTag := Router.Group("/site-tag").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler(),
	)
	{
		siteTag.GET("", siteTagController.GetSiteTagList)
		siteTag.GET("/all", siteTagController.GetAllSiteTags)
		siteTag.POST("", siteTagController.CreateSiteTag)
		siteTag.PUT("/:id", siteTagController.UpdateSiteTag)
		siteTag.DELETE("/:id", siteTagController.DeleteSiteTag)
	}
}
