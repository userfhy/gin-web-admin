package routers

import (
	authController "gin-web-admin/app/controllers/v1/auth"
	casbinController "gin-web-admin/app/controllers/v1/casbin"
	deptController "gin-web-admin/app/controllers/v1/dept"
	dictDataController "gin-web-admin/app/controllers/v1/dict_data"
	dictTypeController "gin-web-admin/app/controllers/v1/dict_type"
	indexController "gin-web-admin/app/controllers/v1/index"
	menuController "gin-web-admin/app/controllers/v1/menu"
	reportController "gin-web-admin/app/controllers/v1/report"
	roleController "gin-web-admin/app/controllers/v1/role"
	siteCategoryController "gin-web-admin/app/controllers/v1/site_category"
	siteContentController "gin-web-admin/app/controllers/v1/site_content"
	sitePublicController "gin-web-admin/app/controllers/v1/site_public"
	siteTagController "gin-web-admin/app/controllers/v1/site_tag"
	sysController "gin-web-admin/app/controllers/v1/sys"
	userController "gin-web-admin/app/controllers/v1/user"
	"gin-web-admin/app/middleware"
	authService "gin-web-admin/app/service/v1/auth"
	casbinService "gin-web-admin/app/service/v1/casbin"
	deptService "gin-web-admin/app/service/v1/dept"
	dictDataService "gin-web-admin/app/service/v1/dict_data"
	dictTypeService "gin-web-admin/app/service/v1/dict_type"
	menuService "gin-web-admin/app/service/v1/menu"
	reportService "gin-web-admin/app/service/v1/report"
	roleService "gin-web-admin/app/service/v1/role"
	siteCategoryService "gin-web-admin/app/service/v1/site_category"
	siteContentService "gin-web-admin/app/service/v1/site_content"
	sitePublicService "gin-web-admin/app/service/v1/site_public"
	siteTagService "gin-web-admin/app/service/v1/site_tag"
	sysService "gin-web-admin/app/service/v1/sys"
	testService "gin-web-admin/app/service/v1/test"
	userService "gin-web-admin/app/service/v1/user"

	"github.com/casbin/casbin/v3"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	CasbinEnforcer      *casbin.SyncedEnforcer
	AuthService         *authService.Service
	CasbinService       *casbinService.Service
	UserService         *userService.Service
	RoleService         *roleService.Service
	MenuService         *menuService.Service
	ReportService       *reportService.Service
	DeptService         *deptService.Service
	DictTypeService     *dictTypeService.Service
	DictDataService     *dictDataService.Service
	SiteCategoryService *siteCategoryService.Service
	SiteContentService  *siteContentService.Service
	SiteTagService      *siteTagService.Service
	SitePublicService   *sitePublicService.Service
	SysService          *sysService.Service
	TestService         *testService.Service
}

func InitRouter(r *gin.Engine, deps Dependencies) *gin.Engine {
	v1 := r.Group("/v1/api")
	public := v1

	casbinMiddleware := middleware.CasbinHandler(deps.CasbinEnforcer)
	casbinHandler := casbinController.NewHandler(deps.CasbinService)
	userHandler := userController.NewHandler(deps.UserService)
	authHandler := authController.NewHandler(deps.AuthService)
	roleHandler := roleController.NewHandler(deps.RoleService)
	menuHandler := menuController.NewHandler(deps.MenuService)
	reportHandler := reportController.NewHandler(deps.ReportService)
	deptHandler := deptController.NewHandler(deps.DeptService)
	dictTypeHandler := dictTypeController.NewHandler(deps.DictTypeService)
	dictDataHandler := dictDataController.NewHandler(deps.DictDataService)
	siteCategoryHandler := siteCategoryController.NewHandler(deps.SiteCategoryService)
	siteContentHandler := siteContentController.NewHandler(deps.SiteContentService)
	siteTagHandler := siteTagController.NewHandler(deps.SiteTagService)
	sitePublicHandler := sitePublicController.NewHandler(deps.SitePublicService)
	sysHandler := sysController.NewHandler(deps.SysService)
	testHandler := indexController.NewHandler(deps.TestService)

	translationOnly := v1.Group("")
	translationOnly.Use(middleware.TranslationHandler())

	authzGroup := v1.Group("")
	authzGroup.Use(
		middleware.TranslationHandler(),
		middleware.JWTHandler(deps.UserService),
		casbinMiddleware,
	)
	authzGroup.Use(middleware.OperationLogger())

	jwtOnly := v1.Group("")
	jwtOnly.Use(
		middleware.JWTHandler(deps.UserService),
		middleware.OperationLogger(),
	)

	InitUserRouter(public, authzGroup, userHandler, authHandler) // 用户管理
	InitRoleRouter(authzGroup, roleHandler)                      // 角色
	InitCasbinRouter(authzGroup, casbinHandler)                  // Casbin
	InitSysRouter(authzGroup, sysHandler)                        // 系统设置
	InitTestRouter(translationOnly, testHandler)                 // 测试路由
	InitReportRouter(translationOnly, reportHandler)             // 上报
	InitMenuRouter(authzGroup, menuHandler)                      // 菜单管理
	InitDeptRouter(authzGroup, deptHandler)                      // 部门管理
	InitDictTypeRouter(authzGroup, dictTypeHandler)              // 字典类型管理
	InitDictDataRouter(authzGroup, dictDataHandler)              // 字典数据管理
	InitSiteContentRouter(authzGroup, siteContentHandler)        // 企业官网内容管理
	InitSiteCategoryRouter(authzGroup, siteCategoryHandler)      // 企业官网分类管理
	InitSiteTagRouter(authzGroup, siteTagHandler)                // 企业官网标签管理
	InitSitePublicRouter(public, sitePublicHandler)              // 企业官网公开接口
	InitAsyncRoutesRouter(jwtOnly, sysHandler)                   // 动态路由

	if gin.Mode() == gin.DebugMode {
		InitSwaggerRouter(r) // swagger docs
	}

	// 路由列表
	sysController.Routers = r.Routes()

	return r
}
