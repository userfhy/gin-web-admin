package routers

import (
	sysController "gin-web-admin/app/controllers/v1/sys"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/api")
	{
		InitUserRouter(v1) // 用户管理
		InitRoleRouter(v1) // 角色
		InitCasbinRouter(v1)
		InitSysRouter(v1)    // 系统设置
		InitTestRouter(v1)   // 测试路由
		InitReportRouter(v1) // 上报
	}

	InitSwaggerRouter(r) // swagger docs

	// 路由列表
	sysController.Routers = r.Routes()

	return r
}
