package main

import (
	"context"
	"fmt"
	"gin-web-admin/app/middleware"
	model "gin-web-admin/app/models"
	"gin-web-admin/common"
	"gin-web-admin/utils/casbin"
	"gin-web-admin/utils/logging"

	//"gin-web-admin/common/validator"
	"gin-web-admin/routers"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/setting"

	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin/binding"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger = logging.Setup("main-logger", nil)

func init() {
	setting.Setup()
	model.Setup()
	common.InitValidate()

	// 初始化路由权限 在这初始化的目的： 避免每次访问路由查询数据库
	// 如果更改路由权限 需要重新调用一下这个方法
	casbin.SetupCasbin()

	if setting.RedisSetting.Host != "" {
		gredis.Setup()
	}
}

// @termsOfService https://github.com/userfhy/gin-web-admin
// @license.name MIT
// @license.url https://github.com/userfhy/gin-web-admin/blob/master/LICENSE

// @securityDefinitions.apikey ApiKeyAuth
// @in header like: Bearer xxxx
// @name Authorization
func main() {
	//binding.Validator = new(validator.DefaultValidator)
	gin.SetMode(setting.ServerSetting.RunMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if setting.AppSetting.EnabledCORS {
		r.Use(middleware.CORS())
	}

	routersInit := routers.InitRouter(r)

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
			logger.Fatalln(err)
		}
	}()

	logger.Printf("[info] start http server listening %s", endPoint)
	logger.Printf("[info] Actual pid is %d", os.Getpid())

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}

	logger.Println("Server exiting")
}
