package routers

import (
	siteCategoryController "gin-web-admin/app/controllers/v1/site_category"
	"gin-web-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitSiteCategoryRouter(Router *gin.RouterGroup) {
	siteCategory := Router.Group("/site-category").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler(),
	)
	{
		siteCategory.GET("", siteCategoryController.GetSiteCategoryList)
		siteCategory.GET("/all", siteCategoryController.GetAllSiteCategories)
		siteCategory.POST("", siteCategoryController.CreateSiteCategory)
		siteCategory.PUT("/:id", siteCategoryController.UpdateSiteCategory)
		siteCategory.DELETE("/:id", siteCategoryController.DeleteSiteCategory)
	}
}
