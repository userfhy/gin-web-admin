package routers

import (
    authController "gin-test/app/controllers/v1/auth"
    indexController "gin-test/app/controllers/v1/index"
    reportController "gin-test/app/controllers/v1/report"
    "gin-test/app/middleware"
    "gin-test/docs"
    "gin-test/utils/casbin"
    "gin-test/utils/setting"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "net/http"
    "strings"
)

var Routers gin.RoutesInfo

func InitRouter() *gin.Engine {
    // programatically set swagger info
    docs.SwaggerInfo.Title = "Swagger Example API"
    docs.SwaggerInfo.Description = "This is a sample server Petstore server."
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

    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    if setting.AppSetting.EnabledCORS {
        r.Use(middleware.CORS())
    }

    // 初始化路由权限 在这初始化的目的： 避免每次访问路由查询数据库
    // 如果更改路由权限 需要重新调用一下这个方法
    casbin.SetupCasbin()

    v1 := r.Group("/v1/api")
    {
        v1.POST("/login", authController.GetAuth) // 登录

        report := v1.Group("/report").Use(
            middleware.TranslationHandler())
        {
            report.POST("", reportController.Report)
        }

        test := v1.Group("/test").Use(
            middleware.TranslationHandler(),
            middleware.JWTHandler(),
            middleware.CasbinHandler())
        {
            // 参数测试
            test.GET("/article/:id", func(c *gin.Context) {
              id := c.Param("id")
              c.JSON(200, gin.H{"id": id})
            })

            test.GET("/article/:id/user/:user", func(c *gin.Context) {
              id := c.Param("id")
              user := c.Param("user")
              c.JSON(200, gin.H{"id": id, "user": user})
            })

            test.GET("/routers", GetRouterList)
            test.POST("/ping", indexController.Ping)
            test.GET("/ping", indexController.Ping)
            test.GET("/font", indexController.Test)
            test.GET("/test_users", indexController.GetTestUsers)
            test.POST("/login", indexController.UserLogin)
        }

    }

    // swagger docs
    r.GET("/swagger", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    Routers = r.Routes()
    //for index := range Routers{
    //    log.Println(routes[index].Path, routes[index].Method, routes[index].HandlerFunc)
    //}

    return r
}

// @Summary Routers
// @Description get router list
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Tags Test
// @Success 200 {object} common.Response
// @Router /test/routers [get]
func GetRouterList(c *gin.Context) {
    type Router struct {
        Path   string `json:"path"`
        Method string `json:"method"`
    }

    data := make([]Router, 0)

    Routers := Routers
    for index := range Routers{
        var router Router
        router.Method = Routers[index].Method
        router.Path = Routers[index].Path
        data = append(data, router)
    }

    c.JSON(200, gin.H{
        "code": 200,
        "msg": "ok",
        "data": data,
    })
}