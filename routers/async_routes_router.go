package routers

import (
	sysController "gin-web-admin/app/controllers/v1/sys"

	"github.com/gin-gonic/gin"
)

// GET /v1/api/get-async-routes
// 仅使用 JWT，不使用 Casbin（菜单本身可按角色过滤/或全量返回）
func InitAsyncRoutesRouter(router *gin.RouterGroup, handler *sysController.Handler) {
	router.GET("/get-async-routes", handler.GetAsyncRoutes)
}
