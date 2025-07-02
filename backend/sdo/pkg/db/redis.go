package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/m1sty32/sdo/pkg/config"
	"golang.org/x/net/context"
)

func NewRedis(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("Failed to connect to Redis: %v", err)
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}
	logrus.Info("Successfully connected to Redis")
	return client, nil
}