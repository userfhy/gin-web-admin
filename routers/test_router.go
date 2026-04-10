package routers

import (
	indexController "gin-web-admin/app/controllers/v1/index"
	"gin-web-admin/views"
	"text/template"

	"github.com/gin-gonic/gin"
)

func InitTestRouter(router *gin.RouterGroup, handler *indexController.Handler) {
	test := router.Group("/test")
	{
		test.POST("/ping", handler.Ping)
		test.GET("/ping", handler.Ping)
		test.GET("/font", handler.Test)

		// 修改SSE HTML模板加载方式
		test.GET("/sse", func(c *gin.Context) {
			// 使用嵌入的模板文件
			t, err := template.ParseFS(views.SSEStaticFS, "sse/testSSE.html")
			if err != nil {
				c.String(500, "template parse error")
				return
			}
			if err := t.Execute(c.Writer, "index"); err != nil {
				c.String(500, "template execute error")
			}
		})

		//注册SSE路由
		test.GET("/events", handler.Stream())

		handler.StartSystemMonitorBroadcast()

		test.POST("/send", handler.SendStream)

		// 获取客户端数量
		test.GET("/count", handler.SSEClientCount)

	}
}
