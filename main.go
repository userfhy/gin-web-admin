package main

import (
    "context"
    "fmt"
    model "gin-test/app/models"
    "gin-test/common"
    "gin-test/routers"
    "gin-test/utils/gredis"
    "gin-test/utils/setting"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func init() {
    setting.Setup()
    model.Setup()
    common.InitValidate()
    if setting.RedisSetting.Host != "" {
        gredis.Setup()
    }
}

// @title Gin Web Test
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/fanhengyuan/gin-test
// @license.name MIT
// @license.url https://github.com/fanhengyuan/gin-test/blob/master/LICENSE
func main() {
    gin.SetMode(setting.ServerSetting.RunMode)

    routersInit := routers.InitRouter()

    readTimeout := setting.ServerSetting.ReadTimeout
    writeTimeout := setting.ServerSetting.WriteTimeout
    endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
    maxHeaderBytes := 1 << 20

    server := &http.Server{
        Addr:           endPoint,
        Handler:        routersInit,
        ReadTimeout:    readTimeout,
        WriteTimeout:   writeTimeout,
        MaxHeaderBytes: maxHeaderBytes,
    }

    go func() {
        // 服务连接
        err := server.ListenAndServe()
        if err != nil {
            log.Fatalln(err)
        }
    }()

    log.Printf("[info] start http server listening %s", endPoint)
    log.Printf("[info] Actual pid is %d", os.Getpid())

    // 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
    quit := make(chan os.Signal, 1)
    // kill (no param) default send syscanll.SIGTERM
    // kill -2 is syscall.SIGINT
    // kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutdown Server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown: ", err)
    }

    log.Println("Server exiting")
}
