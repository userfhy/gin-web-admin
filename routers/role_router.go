package routers

import (
	roleController "gin-web-admin/app/controllers/v1/role"
	"gin-web-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	role := Router.Group("/role").Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(),
		middleware.CasbinHandler())
	{
		role.GET("", roleController.GetRoles)               // 获取角色列表
		role.POST("", roleController.CreateRole)            // 创建角色
		role.PUT("/:role_id", roleController.UpdateRole)    // 修改角色
		role.DELETE("/:role_id", roleController.DeleteRole) // 删除
	}
}
