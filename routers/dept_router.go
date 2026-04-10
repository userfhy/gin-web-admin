package routers

import (
	deptController "gin-web-admin/app/controllers/v1/dept"

	"github.com/gin-gonic/gin"
)

func InitDeptRouter(router *gin.RouterGroup, handler *deptController.Handler) {
	dept := router.Group("/dept")
	{
		dept.GET("", handler.GetDeptList)
		dept.POST("", handler.CreateDept)
		dept.PUT("/:id", handler.UpdateDept)
		dept.DELETE("/:id", handler.DeleteDept)
	}
}
