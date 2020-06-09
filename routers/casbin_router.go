package routers

import (
    casbinController "gin-test/app/controllers/v1/casbin"
    "gin-test/app/middleware"
    "github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
    casbin := Router.Group("/casbin").Use(
        middleware.TranslationHandler(),
        middleware.JWTHandler(),
        middleware.CasbinHandler())
    {
        casbin.GET("", casbinController.GetCasbinList) // 存在规则
        casbin.POST("", casbinController.CreateCasbin)
        casbin.PUT("/:id", casbinController.UpdateCasbin)
        casbin.DELETE("/:id", casbinController.DeleteCasbin)
    }
}