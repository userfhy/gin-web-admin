package routers

import (
	authController "gin-web-admin/app/controllers/v1/auth"
	userController "gin-web-admin/app/controllers/v1/user"
	"gin-web-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	Router.POST("/login", authController.UserLogin) // 登录

	user := Router.Group("/user").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler())
	{
		// user.POST("", userController.CreateUser)                    // 创建用户
		user.GET("", userController.GetUsers)                       // 用户列表
		user.PUT("/logout", authController.UserLogout)              // 登出
		user.PUT("/change_password", authController.ChangePassword) // 修改密码
		user.GET("/logged_in", authController.GetLoggedInUser)      // 当前登录用户信息
	}
}
