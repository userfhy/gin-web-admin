package routers

import (
	reportController "gin-web-admin/app/controllers/v1/report"

	"github.com/gin-gonic/gin"
)

func InitReportRouter(router *gin.RouterGroup, handler *reportController.Handler) {
	report := router.Group("/report")
	{
		report.POST("", handler.Report)
	}
}
