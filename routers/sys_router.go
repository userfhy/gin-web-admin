package routers

import (
	sysController "gin-web-admin/app/controllers/v1/sys"

	"github.com/gin-gonic/gin"
)

func InitSysRouter(router *gin.RouterGroup, handler *sysController.Handler) {
	sys := router.Group("/sys")
	{
		sys.GET("/router", handler.GetRouterList) // 路由列表
		sys.GET("/permission-visualization", handler.GetPermissionVisualization)
		sys.GET("/online-users", handler.GetOnlineUsers)
		sys.DELETE("/online-users/:userId", handler.ForceOffline)
		sys.GET("/server-monitor", handler.GetServerMonitor)
		sys.GET("/server-monitor/ws", handler.StreamServerMonitor)
		sys.GET("/login-logs", handler.GetLoginLogs)
		sys.DELETE("/login-logs", handler.DeleteLoginLogs)
		sys.GET("/operation-logs", handler.GetOperationLogs)
		sys.DELETE("/operation-logs", handler.DeleteOperationLogs)
		sys.GET("/system-logs", handler.GetSystemLogs)
		sys.GET("/system-logs/:id", handler.GetSystemLogDetail)
		sys.DELETE("/system-logs", handler.DeleteSystemLogs)
		// sys.GET("/menu_list", handler.GetMenuList) // 菜单列表
	}
}
