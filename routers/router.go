package routers

import (
    indexController "gin-test/app/controllers/index"
    "gin-test/docs"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
    // programatically set swagger info
    docs.SwaggerInfo.Title = "Swagger Example API"
    docs.SwaggerInfo.Description = "This is a sample server Petstore server."
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "127.0.0.1:8080"
    docs.SwaggerInfo.BasePath = "/v1/api/"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    v1 := r.Group("/v1/api")
    {
        test := v1.Group("/test")
        {
            test.GET("/ping", indexController.Ping)
            test.GET("/font", indexController.Test)
            test.GET("/test_users", indexController.GetTestUsers)
        }
    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
