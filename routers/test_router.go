package routers

import (
	indexController "gin-web-admin/app/controllers/v1/index"
	"gin-web-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestRouter(Router *gin.RouterGroup) {
	test := Router.Group("/test").Use(
		middleware.TranslationHandler(),
	)
	{
		test.POST("/ping", indexController.Ping)
		test.GET("/ping", indexController.Ping)
		test.GET("/font", indexController.Test)
	}
}
