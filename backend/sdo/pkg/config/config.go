package config

import (
	"os"
	"strconv"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Redis struct {
		Host string
		Port int
		Addr string
	}
	Server struct {
		Port int
	}
	SecretKey string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	
	cfg.Database.Host = os.Getenv("DB_HOST")
	if cfg.Database.Host == "" {
		cfg.Database.Host = "postgres"
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 5432
	}
	cfg.Database.Port = port
	cfg.Database.User = os.Getenv("DB_USER")
	if cfg.Database.User == "" {
		cfg.Database.User = "kktd_user"
	}
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	if cfg.Database.Password == "" {
		cfg.Database.Password = "111"
	}
	cfg.Database.Name = os.Getenv("DB_NAME")
	if cfg.Database.Name == "" {
		cfg.Database.Name = "sdo_db"
	}
	
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	if cfg.Redis.Host == "" {
		cfg.Redis.Host = "redis"
	}
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		redisPort = 6379
	}
	cfg.Redis.Port = redisPort
	cfg.Redis.Addr = cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port)
	
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		appPort = 8080
	}
	cfg.Server.Port = appPort
	
	cfg.SecretKey = os.Getenv("SECRET_KEY")
	if cfg.SecretKey == "" {
		cfg.SecretKey = "EuuwIes2/2fDPnsgCuCWsOgX/lOLOhnKl29+HBgwiJ8="
	}
	
	return cfg, nil
}