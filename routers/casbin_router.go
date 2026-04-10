package routers

import (
	casbinController "gin-web-admin/app/controllers/v1/casbin"

	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(router *gin.RouterGroup, handler *casbinController.Handler) {
	casbin := router.Group("/casbin")
	{
		casbin.GET("", handler.GetCasbinList) // 存在规则
		casbin.POST("", handler.CreateCasbin)
		casbin.PUT("/:id", handler.UpdateCasbin)
		casbin.DELETE("/:id", handler.DeleteCasbin)
	}
}
