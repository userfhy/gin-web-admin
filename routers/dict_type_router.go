package routers

import (
	dictTypeController "gin-web-admin/app/controllers/v1/dict_type"

	"github.com/gin-gonic/gin"
)

func InitDictTypeRouter(router *gin.RouterGroup, handler *dictTypeController.Handler) {
	dictType := router.Group("/dict/type")
	{
		dictType.GET("", handler.GetDictTypeList)
		dictType.GET("/all", handler.GetAllDictTypes)
		dictType.POST("", handler.CreateDictType)
		dictType.PUT("/:id", handler.UpdateDictType)
		dictType.DELETE("/:id", handler.DeleteDictType)
	}
}
