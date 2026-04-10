package routers

import (
	sysController "gin-web-admin/app/controllers/v1/sys"

	"github.com/gin-gonic/gin"
)

func InitSysRouter(router *gin.RouterGroup, handler *sysController.Handler) {
	sys := router.Group("/sys")
	{
		sys.GET("/router", handler.GetRouterList) // 路由列表
		// sys.GET("/menu_list", handler.GetMenuList) // 菜单列表
	}
}
