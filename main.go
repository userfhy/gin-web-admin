package main

import (
    "fmt"
    "gin-test/routers"
    "gin-test/utils/setting"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func init() {
    setting.Setup()
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

    log.Printf("[info] start http server listening %s", endPoint)

    err := server.ListenAndServe()
    if err != nil {
        log.Println(err)
        return
    }
}
