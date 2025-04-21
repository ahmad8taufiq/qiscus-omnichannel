package redis

import (
	"context"
	"time"

	"qiscus-omnichannel/tools/logger"

	"github.com/redis/go-redis/v9"

	"qiscus-omnichannel/config"
)

var (
	RedisClient *redis.Client
	RedisCtx    = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisHost,
		Username: config.AppConfig.RedisUser,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	_, err := RedisClient.Ping(RedisCtx).Result()
	if err != nil {
		logger.Logger.WithError(err).Fatal("❌ Failed to connect to Redis")
	} else {
		logger.Logger.Info("🔌 Connected to Redis")
	}
}

func SetCache(key string, value string, ttl time.Duration) error {
	return RedisClient.Set(RedisCtx, key, value, ttl).Err()
}

func GetCache(key string) (string, error) {
	return RedisClient.Get(RedisCtx, key).Result()
}