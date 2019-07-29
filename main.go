package main

import (
    indexController "gin-test/app/controllers/index"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong \\(^o^)/~ 测试",
        })
    })

    r.GET("/font", indexController.Test)

    log.Println("Server startup...")
    r.Run()
}
