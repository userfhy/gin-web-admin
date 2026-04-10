package routers

import (
	sitePublicController "gin-web-admin/app/controllers/v1/site_public"

	"github.com/gin-gonic/gin"
)

func InitSitePublicRouter(Router *gin.RouterGroup, handler *sitePublicController.Handler) {
	site := Router.Group("/site/public")
	{
		site.GET("/categories", handler.GetPublicCategories)
		site.GET("/tags", handler.GetPublicTags)
		site.GET("/contents", handler.GetPublicContentList)
		site.GET("/contents/:id", handler.GetPublicContentDetail)
		site.GET("/contents/slug/:slug", handler.GetPublicContentDetailBySlug)
	}
}
