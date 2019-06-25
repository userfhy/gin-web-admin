package main

import (
    "log"

    "github.com/fanhengyuan/gin-test/app/controllers/index"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong \\(^o^)/~ 测试",
        })
    })

    r.GET("/font", index_controller.Test)

    log.Println("Server startup...")
    r.Run()
}
