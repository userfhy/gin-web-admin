package routers

import (
	indexController "gin-web-admin/app/controllers/v1/index"
	"gin-web-admin/app/middleware"
	"text/template"

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

		//SSE HTML
		test.GET("/sse", func(c *gin.Context) {
			t, err := template.ParseFiles("./views/sse/testSSE.html")
			if err != nil {
				panic(err)
			}
			t.Execute(c.Writer, "index")
		})

		//注册SSE路由
		test.GET("/events", indexController.SSEService.Handler())

		// 启动广播
		// go func() {
		// 	ticker := time.NewTicker(5 * time.Second)
		// 	for range ticker.C {
		// 		indexController.SSEService.Broadcast(sse.Message{
		// 			Event: "broadcast",
		// 			Data:  time.Now().Format(time.RFC3339),
		// 		})
		// 	}
		// }()

		test.POST("/send", indexController.SendStream)

		// 获取客户端数量
		test.GET("/count", indexController.SSEClientCount)

	}
}
