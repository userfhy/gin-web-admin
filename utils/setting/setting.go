package setting

import (
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

// 统一配置结构体
type Config struct {
	App      App
	Server   Server
	Database Database
	Redis    Redis
	Security Security
}

type App struct {
	JwtSecret        string
	PasswordSalt     string
	PrefixUrl        string
	TimeFormat       string
	EnabledCORS      bool
	ExpireTimeFormat string `toml:"ExpireTimeFormat"` // 解决字段名不一致问题
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	EchoSql     bool
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	DB          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Security struct {
	PasswordMinLength   int
	RequireUppercase    bool
	RequireLowercase    bool
	RequireNumber       bool
	RequireSpecial      bool
	LoginMaxAttempts    int
	LoginLockoutMinutes int
	IPWhitelist         []string
}

var SecuritySetting = &Security{}

var (
	cfg        *Config
	configPath = "conf/app.toml"
)

// SetConfigPath 允许在运行时覆盖配置文件路径，便于多环境部署或测试。
func SetConfigPath(path string) {
	if path != "" {
		configPath = path
	}
}

// Setup 初始化配置
func Setup() {
	path := configPath
	if envPath := os.Getenv("APP_CONFIG_PATH"); envPath != "" {
		path = envPath
	}

	cfg = &Config{}
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.toml': %v", err)
	}

	// 映射配置到全局变量
	AppSetting = &cfg.App
	ServerSetting = &cfg.Server
	DatabaseSetting = &cfg.Database
	RedisSetting = &cfg.Redis
	SecuritySetting = &cfg.Security

	// 从环境变量覆盖敏感配置
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		AppSetting.JwtSecret = jwtSecret
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		DatabaseSetting.User = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		DatabaseSetting.Password = dbPassword
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		DatabaseSetting.Host = dbHost
	}
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		RedisSetting.Host = redisHost
	}
	if redisPassword := os.Getenv("REDIS_PASSWORD"); redisPassword != "" {
		RedisSetting.Password = redisPassword
	}

	// 转换时间单位（TOML解析可以直接处理time.Duration类型，但需要确保配置文件中的数值单位）
	// 如果配置文件中是秒数，需要手动转换
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	if SecuritySetting.PasswordMinLength <= 0 {
		SecuritySetting.PasswordMinLength = 8
	}
	if SecuritySetting.LoginMaxAttempts <= 0 {
		SecuritySetting.LoginMaxAttempts = 5
	}
	if SecuritySetting.LoginLockoutMinutes <= 0 {
		SecuritySetting.LoginLockoutMinutes = 15
	}
}
