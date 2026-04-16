package routers

import (
	dictDataController "gin-web-admin/app/controllers/v1/dict_data"

	"github.com/gin-gonic/gin"
)

func InitDictDataRouter(router *gin.RouterGroup, handler *dictDataController.Handler) {
	dictData := router.Group("/dict/data")
	{
		dictData.GET("", handler.GetDictDataList)
		dictData.POST("", handler.CreateDictData)
		dictData.PUT("/:id", handler.UpdateDictData)
		dictData.DELETE("/:id", handler.DeleteDictData)
	}
}
