package bootstrap

import (
	"fmt"
	"os"

	casbinv3 "github.com/casbin/casbin/v3"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-web-admin/app/middleware"
	model "gin-web-admin/app/models"
	"gin-web-admin/common"
	"gin-web-admin/internal/data"
	casbinutil "gin-web-admin/utils/casbin"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/setting"
)

// Options 配置初始化所需的参数
type Options struct {
	ConfigPath string
}

// Container 保存初始化后的依赖，便于在应用中传递
type Container struct {
	RunMode        string
	DB             *gorm.DB
	RedisPool      *redis.Pool
	CasbinEnforcer *casbinv3.SyncedEnforcer
	Store          *data.Store
}

func Initialize(opts Options) (*Container, error) {
	if opts.ConfigPath != "" {
		setting.SetConfigPath(opts.ConfigPath)
	}

	setting.Setup()

	runMode := setting.ServerSetting.RunMode
	if envRunMode := os.Getenv("RUN_MODE"); envRunMode != "" {
		runMode = envRunMode
	}

	gin.SetMode(runMode)

	logLevel := gin.DebugMode
	if runMode != gin.DebugMode {
		logLevel = "info"
	}

	logging.Setup("main-logger", &logging.Option{
		LogLevel:   logLevel,
		Formatter:  "text",
		OutputPath: "",
	})

	logging.Info("runMode:", gin.Mode())

	model.Setup()
	common.InitValidate()

	if setting.RedisSetting.Host != "" {
		gredis.Setup()
	}

	enforcer := casbinutil.SetupCasbin()
	if enforcer == nil {
		return nil, fmt.Errorf("failed to initialize casbin")
	}

	var redisPool *redis.Pool
	if gredis.RedisConn != nil {
		redisPool = gredis.RedisConn
	}

	db := model.DB()

	return &Container{
		RunMode:        runMode,
		DB:             db,
		RedisPool:      redisPool,
		CasbinEnforcer: enforcer,
		Store:          data.NewStore(db, redisPool),
	}, nil
}

func (c *Container) NewEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	if setting.AppSetting.EnabledCORS {
		engine.Use(middleware.CORS())
	}

	return engine
}
