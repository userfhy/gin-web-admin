package routers

import (
    "gin-test/docs"
    "gin-test/utils/setting"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "net/http"
    "strings"
)

func InitSwaggerRouter(Router *gin.Engine) {
    // programmatically set swagger info
    docs.SwaggerInfo.Title = "Gin Web Test."
    docs.SwaggerInfo.Description = "An example of gin."
    docs.SwaggerInfo.Version = "1.0"

    configHost := ""
    prefixUrl := setting.AppSetting.PrefixUrl
    if strings.HasPrefix(prefixUrl, "http://") {
        configHost = strings.Replace(prefixUrl, "http://", "", -1)
        docs.SwaggerInfo.Schemes = []string{"http"}
    }else if strings.HasPrefix(prefixUrl, "https://") {
        configHost = strings.Replace(prefixUrl, "https://", "", -1)
        docs.SwaggerInfo.Schemes = []string{"https"}
    }
    docs.SwaggerInfo.Host = configHost
    docs.SwaggerInfo.BasePath = "/v1/api/"

    Router.GET("/swagger", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })
    Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
