package routers

import (
	authController "gin-web-admin/app/controllers/v1/auth"
	userController "gin-web-admin/app/controllers/v1/user"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(public, protected *gin.RouterGroup, userHandler *userController.Handler, authHandler *authController.Handler) {
	public.POST("/login", authHandler.UserLogin)                  // 登录
	public.POST("/refresh_token", authHandler.RefreshAccessToken) // 刷新access_token

	user := protected.Group("/user")
	{
		user.GET("", userHandler.GetUsers)                       // 用户列表
		user.PUT("/logout", authHandler.UserLogout)              // 登出
		user.PUT("/change_password", authHandler.ChangePassword) // 修改密码
		user.GET("/logged_in", authHandler.GetLoggedInUser)      // 当前登录用户信息
	}
}
