package routers

import (
    indexController "gin-test/app/controllers/index"
    "gin-test/common"
    _ "gin-test/docs"
    "gin-test/utils/code"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "net/http"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    testApi := r.Group("/test")
    {
        testApi.GET("/ping", func(c *gin.Context) {
            appG := common.Gin{C: c}
            appG.Response(http.StatusOK, code.SUCCESS, "pong", nil)
        })
    }

    r.GET("/font", indexController.Test)

    return r
}
