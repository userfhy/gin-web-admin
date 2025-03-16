package setting

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

// 统一配置结构体
type Config struct {
	App      App
	Server   Server
	Database Database
	Redis    Redis
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

var cfg *Config

// Setup 初始化配置
func Setup() {
	cfg = &Config{}
	if _, err := toml.DecodeFile("conf/app.toml", cfg); err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.toml': %v", err)
	}

	// 映射配置到全局变量
	AppSetting = &cfg.App
	ServerSetting = &cfg.Server
	DatabaseSetting = &cfg.Database
	RedisSetting = &cfg.Redis

	// 转换时间单位（TOML解析可以直接处理time.Duration类型，但需要确保配置文件中的数值单位）
	// 如果配置文件中是秒数，需要手动转换
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}
