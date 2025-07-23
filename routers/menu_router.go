package routers

import (
	menuController "gin-web-admin/app/controllers/v1/menu"
	"gin-web-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	menu := Router.Group("/menu").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler())
	{

		menu.GET("", menuController.GetMenuList)       // 获取菜单列表（分页）
		menu.GET("/all", menuController.GetAllMenus)   // 获取全部菜单（无分页）
		menu.POST("", menuController.CreateMenu)       // 新增菜单
		menu.PUT("/:id", menuController.UpdateMenu)    // 更新菜单
		menu.DELETE("/:id", menuController.DeleteMenu) // 删除菜单
	}
}
