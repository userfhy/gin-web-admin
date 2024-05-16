package routers

import (
	"gin-web-admin/docs"
	"gin-web-admin/utils/setting"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRouter(Router *gin.Engine) {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Gin Web Admin."
	docs.SwaggerInfo.Description = "Build web admin with gin-gonic."
	docs.SwaggerInfo.Version = "1.0"

	configHost := ""
	prefixUrl := setting.AppSetting.PrefixUrl
	if strings.HasPrefix(prefixUrl, "http://") {
		configHost = strings.Replace(prefixUrl, "http://", "", -1)
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else if strings.HasPrefix(prefixUrl, "https://") {
		configHost = strings.Replace(prefixUrl, "https://", "", -1)
		docs.SwaggerInfo.Schemes = []string{"https"}
	}
	docs.SwaggerInfo.Host = configHost
	docs.SwaggerInfo.BasePath = "/v1/api/"

	Router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
