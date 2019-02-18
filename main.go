package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/fanhengyuan/gin-test/app/controllers/index"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong \\(^o^)/~ 测试",
        })
    })

    r.GET("/", index_controller.Test)

    log.Println("Server startup...")
    r.Run()
}
