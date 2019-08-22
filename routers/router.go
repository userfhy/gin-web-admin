package routers

import (
    indexController "gin-test/app/controllers/index"
    "github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

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
