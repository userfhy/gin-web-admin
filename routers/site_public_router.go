package routers

import (
	sitePublicController "gin-web-admin/app/controllers/v1/site_public"

	"github.com/gin-gonic/gin"
)

func InitSitePublicRouter(Router *gin.RouterGroup) {
	site := Router.Group("/site/public")
	{
		site.GET("/categories", sitePublicController.GetPublicCategories)
		site.GET("/tags", sitePublicController.GetPublicTags)
		site.GET("/contents", sitePublicController.GetPublicContentList)
		site.GET("/contents/:id", sitePublicController.GetPublicContentDetail)
		site.GET("/contents/slug/:slug", sitePublicController.GetPublicContentDetailBySlug)
	}
}
