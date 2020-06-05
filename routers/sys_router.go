package routers

import (
    sysController "gin-test/app/controllers/v1/sys"
    "gin-test/app/middleware"
    "github.com/gin-gonic/gin"
)

func InitSysRouter(Router *gin.RouterGroup) {
    sys := Router.Group("/sys").Use(
        middleware.TranslationHandler(),
        middleware.JWTHandler(),
        middleware.CasbinHandler())
    {
        sys.GET("/router", sysController.GetRouterList) // 路由列表
    }
}