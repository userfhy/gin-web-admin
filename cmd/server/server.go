package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	authService "gin-web-admin/app/service/v1/auth"
	casbinService "gin-web-admin/app/service/v1/casbin"
	deptService "gin-web-admin/app/service/v1/dept"
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
	"gin-web-admin/internal/bootstrap"
	"gin-web-admin/routers"
	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/setting"
)

// Options 控制服务器启动行为
type Options struct {
	ConfigPath      string
	ShutdownTimeout time.Duration
}

func Run(opts Options) error {
	container, err := bootstrap.Initialize(bootstrap.Options{
		ConfigPath: opts.ConfigPath,
	})
	if err != nil {
		return err
	}

	userSvc := userService.NewService(container.Store)

	authSvc := authService.NewService(container.Store, userSvc)

	roleSvc := roleService.NewService(container.Store)

	menuSvc := menuService.NewService(container.Store)

	casbinSvc := casbinService.NewService(container.Store)

	reportSvc := reportService.NewService(container.Store)

	deptSvc := deptService.NewService(container.Store)

	siteCategorySvc := siteCategoryService.NewService(container.Store)

	siteContentSvc := siteContentService.NewService(container.Store)

	siteTagSvc := siteTagService.NewService(container.Store)

	sitePublicSvc := sitePublicService.NewService(container.Store)

	sysSvc := sysService.NewService(container.Store)

	testSvc := testService.NewService()

	engine := container.NewEngine()
	routersInit := routers.InitRouter(engine, routers.Dependencies{
		CasbinEnforcer:      container.CasbinEnforcer,
		AuthService:         authSvc,
		CasbinService:       casbinSvc,
		UserService:         userSvc,
		RoleService:         roleSvc,
		MenuService:         menuSvc,
		ReportService:       reportSvc,
		DeptService:         deptSvc,
		SiteCategoryService: siteCategorySvc,
		SiteContentService:  siteContentSvc,
		SiteTagService:      siteTagSvc,
		SitePublicService:   sitePublicSvc,
		SysService:          sysSvc,
		TestService:         testSvc,
	})

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	errCh := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			errCh <- err
		} else {
			errCh <- nil
		}
	}()

	logging.Infof("start http server listening %s", endPoint)
	logging.Infof("Actual pid is %d", os.Getpid())

	timeout := opts.ShutdownTimeout
	if timeout <= 0 {
		timeout = 5 * time.Second
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		logging.Infof("Shutdown signal received: %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			return err
		}
		<-errCh
		logging.Info("Server exiting")
		return nil
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		logging.Info("Server stopped")
		return nil
	}
}
