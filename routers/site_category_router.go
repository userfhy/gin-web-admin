package routers

import (
	siteCategoryController "gin-web-admin/app/controllers/v1/site_category"

	"github.com/gin-gonic/gin"
)

func InitSiteCategoryRouter(router *gin.RouterGroup, handler *siteCategoryController.Handler) {
	siteCategory := router.Group("/site-category")
	{
		siteCategory.GET("", handler.GetSiteCategoryList)
		siteCategory.GET("/all", handler.GetAllSiteCategories)
		siteCategory.POST("", handler.CreateSiteCategory)
		siteCategory.PUT("/:id", handler.UpdateSiteCategory)
		siteCategory.DELETE("/:id", handler.DeleteSiteCategory)
	}
}
