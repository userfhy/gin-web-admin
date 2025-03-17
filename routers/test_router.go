package routers

import (
	indexController "gin-web-admin/app/controllers/v1/index"
	"gin-web-admin/app/middleware"
	"gin-web-admin/common/sse"
	"gin-web-admin/utils/system_monitor"
	"gin-web-admin/views"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func InitTestRouter(Router *gin.RouterGroup) {
	test := Router.Group("/test").Use(
		middleware.TranslationHandler(),
	)
	{
		test.POST("/ping", indexController.Ping)
		test.GET("/ping", indexController.Ping)
		test.GET("/font", indexController.Test)

		// 修改SSE HTML模板加载方式
		test.GET("/sse", func(c *gin.Context) {
			// 使用嵌入的模板文件
			t, err := template.ParseFS(views.SSEStaticFS, "sse/testSSE.html")
			if err != nil {
				panic(err)
			}
			t.Execute(c.Writer, "index")
		})

		//注册SSE路由
		test.GET("/events", indexController.SSEService.Handler())

		// 启动系统监控广播
		go func() {
			ticker := time.NewTicker(5 * time.Second)
			for range ticker.C {
				// 获取系统参数
				stats, err := system_monitor.GetSystemStats()
				if err != nil {
					panic(err)
				}

				indexController.SSEService.Broadcast(sse.Message{
					Event: "system_status",
					Data: gin.H{
						"systemStats": stats.String(),
						"clients":     indexController.SSEService.ClientCount(),
					},
				})
			}
		}()

		test.POST("/send", indexController.SendStream)

		// 获取客户端数量
		test.GET("/count", indexController.SSEClientCount)

	}
}
