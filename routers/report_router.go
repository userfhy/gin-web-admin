package routers

import (
	reportController "gin-web-admin/app/controllers/v1/report"
	"gin-web-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitReportRouter(Router *gin.RouterGroup) {
	report := Router.Group("/report").Use(
		middleware.TranslationHandler())
	{
		report.POST("", reportController.Report)
	}
}
