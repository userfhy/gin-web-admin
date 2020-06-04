package routers

import (
    authController "gin-test/app/controllers/v1/auth"
    indexController "gin-test/app/controllers/v1/index"
    reportController "gin-test/app/controllers/v1/report"
    sysController "gin-test/app/controllers/v1/sys"
    userController "gin-test/app/controllers/v1/user"
    roleController "gin-test/app/controllers/v1/role"
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

func InitRouter() *gin.Engine {
    // programmatically set swagger info
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
        v1.POST("/login", authController.UserLogin) // 登录

        // 用户管理
        user := v1.Group("/user").Use(
            middleware.TranslationHandler(),
            middleware.JWTHandler(),
            middleware.CasbinHandler())
        {
            user.POST("", userController.CreateUser) // 创建用户
            user.GET("", userController.GetUsers) // 用户列表
            user.PUT("/logout", authController.UserLogout) // 登出
            user.PUT("/change_password", authController.ChangePassword) // 修改密码
            user.GET("/logged_in", authController.GetLoggedInUser) // 当前登录用户信息
        }

        // 角色
        role := v1.Group("/role").Use(
            middleware.TranslationHandler(),
            middleware.JWTHandler(),
            middleware.CasbinHandler())
        {
            role.GET("", roleController.GetRoles) // 获取角色列表
            role.PUT("/:role_id", roleController.UpdateRole) // 修改角色
        }

        report := v1.Group("/report").Use(
            middleware.TranslationHandler())
        {
            report.POST("", reportController.Report)
        }

        // 系统设置
        sys := v1.Group("/sys").Use(
            middleware.TranslationHandler(),
            middleware.JWTHandler(),
            middleware.CasbinHandler())
        {
            sys.GET("/router", sysController.GetRouterList)
        }

        // 测试
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

            test.POST("/ping", indexController.Ping)
            test.GET("/ping", indexController.Ping)
            test.GET("/font", indexController.Test)
        }

    }

    // swagger docs
    r.GET("/swagger", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // 路由列表
    sysController.Routers = r.Routes()

    return r
}