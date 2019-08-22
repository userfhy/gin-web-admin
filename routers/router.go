package routers

import (
    indexController "gin-test/app/controllers/index"
    _ "gin-test/docs"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    testApi := r.Group("/test")
    {
        testApi.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong \\(^o^)/~ 测试",
            })
        })
    }

    r.GET("/font", indexController.Test)

    return r
}
