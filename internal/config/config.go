package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env           string `yaml:"env" env-default:"local"`
	StorageConfig `yaml:"storage_config" env-required:"true"`
	HTTPServer    `yaml:"http_server"`
}

type NewConfig struct {
	Env         string
	Addr        string
	DbConn      string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type StorageConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Username string `yaml:"username" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	DBName   string `yaml:"db_name" env-default:"postgres"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./config/local.yaml", "path to config file")
}

func MustLoad() *Config {
	flag.Parse()

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

func New() *NewConfig {
	envMode := os.Getenv("ENV_MODE")
	if envMode == "" {
		log.Fatalf("enviroment variable ENV_MODE is not set")
	}

	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		log.Fatalf("enviroment variable SERVER_ADDR is not set")
	}

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatalf("enviroment variable DB_CONNECTION_STRING is not set")
	}

	serverTimeout := os.Getenv("SERVER_TIMEOUT")
	if serverTimeout == "" {
		log.Fatalf("enviroment variable SERVER_TIMEOUT is not set")
	}
	timeout, err := time.ParseDuration(serverTimeout)
	if err != nil {
		log.Fatalf("no valid data formet in SERVER_TIMEOUT variable")
	}

	idleTimeout := os.Getenv("IDLE_TIMEOUT")
	if idleTimeout == "" {
		log.Fatalf("enviroment variable IDLE_TIMEOUT is not set")
	}
	idle, err := time.ParseDuration(idleTimeout)
	if err != nil {
		log.Fatalf("no valid data formet in IDLE_TIMEOUT variable")
	}

	return &NewConfig{
		Env:         envMode,
		Addr:        addr,
		DbConn:      connectionString,
		Timeout:     timeout,
		IdleTimeout: idle,
	}
}
