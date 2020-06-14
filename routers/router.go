package routers

import (
	sysController "gin-test/app/controllers/v1/sys"
	"gin-test/app/middleware"
	"gin-test/utils/casbin"
	"gin-test/utils/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
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
