package routers

import (
	deptController "gin-web-admin/app/controllers/v1/dept"
	"gin-web-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitDeptRouter(Router *gin.RouterGroup) {
	dept := Router.Group("/dept").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler(),
	)
	{
		dept.GET("", deptController.GetDeptList)
		dept.POST("", deptController.CreateDept)
		dept.PUT("/:id", deptController.UpdateDept)
		dept.DELETE("/:id", deptController.DeleteDept)
	}
}
