package routers

import (
	menuController "gin-web-admin/app/controllers/v1/menu"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(router *gin.RouterGroup, handler *menuController.Handler) {
	menu := router.Group("/menu")
	{

		menu.GET("", handler.GetMenuList)       // 获取菜单列表（分页）
		menu.GET("/all", handler.GetAllMenus)   // 获取全部菜单（无分页）
		menu.POST("", handler.CreateMenu)       // 新增菜单
		menu.PUT("/:id", handler.UpdateMenu)    // 更新菜单
		menu.DELETE("/:id", handler.DeleteMenu) // 删除菜单
	}
}
