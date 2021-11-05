package routers

import (
	sysController "gin-web-admin/app/controllers/v1/sys"
	"gin-web-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysRouter(Router *gin.RouterGroup) {
	sys := Router.Group("/sys").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler())
	{
		sys.GET("/router", sysController.GetRouterList)  // 路由列表
		sys.GET("/menu_list", sysController.GetMenuList) // 菜单列表
	}
}
