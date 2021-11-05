package routers

import (
	sysController "gin-web-admin/app/controllers/v1/sys"
	"gin-web-admin/utils/casbin"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	// 初始化路由权限 在这初始化的目的： 避免每次访问路由查询数据库
	// 如果更改路由权限 需要重新调用一下这个方法
	casbin.SetupCasbin()

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
