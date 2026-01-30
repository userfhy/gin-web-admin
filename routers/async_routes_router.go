package routers

import (
	sysController "gin-web-admin/app/controllers/v1/sys"
	"gin-web-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// GET /v1/api/get-async-routes
// 仅使用 JWT，不使用 Casbin（菜单本身可按角色过滤/或全量返回）
func InitAsyncRoutesRouter(Router *gin.RouterGroup) {
	Router.GET("/get-async-routes", middleware.JWTHandler(), sysController.GetAsyncRoutes)
}
