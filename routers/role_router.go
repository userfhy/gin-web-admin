package routers

import (
	roleController "gin-web-admin/app/controllers/v1/role"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(router *gin.RouterGroup, handler *roleController.Handler) {
	role := router.Group("/role")
	{
		role.GET("", handler.GetRoles)               // 获取角色列表
		role.POST("", handler.CreateRole)            // 创建角色
		role.PUT("/:role_id", handler.UpdateRole)    // 修改角色
		role.DELETE("/:role_id", handler.DeleteRole) // 删除

		role.GET("/:role_id/menu_ids", handler.GetRoleMenuIds)
		role.PUT("/:role_id/menu_ids", handler.SaveRoleMenuIds)
	}
}
