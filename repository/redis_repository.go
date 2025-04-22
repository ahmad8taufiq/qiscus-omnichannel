package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"qiscus-omnichannel/config"
	"qiscus-omnichannel/tools/logger"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	SetCache(key string, value string, ttl time.Duration) error
	GetCache(key string) (string, error)
	SetJSON(key string, value interface{}, ttl time.Duration) error
	GetJSON(key string, target interface{}) error
	UpdateJSONField(key string, jsonPath string, value interface{}) error
	DeleteJSON(key string) error
	Enqueue(key string, value interface{}) error
	Dequeue(key string) ([]byte, error)
	Backqueue(key string, value interface{}) error
	BackQueueAtomic(key string, value string) error
}

type redisRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepository() RedisRepository {
	addr := fmt.Sprintf("%s:%s", config.AppConfig.RedisHost, config.AppConfig.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: config.AppConfig.RedisUser,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		logger.Logger.WithError(err).Fatal("❌ Failed to connect to Redis")
	} else {
		logger.Logger.Info("🔌 Connected to Redis")
	}

	return &redisRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *redisRepository) SetCache(key string, value string, ttl time.Duration) error {
	return r.client.Set(r.ctx, key, value, ttl).Err()
}

func (r *redisRepository) GetCache(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *redisRepository) SetJSON(key string, value interface{}, ttl time.Duration) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := r.client.Do(r.ctx, "JSON.SET", key, "$", jsonBytes).Err(); err != nil {
		return err
	}

	if ttl > 0 {
		return r.client.Expire(r.ctx, key, ttl).Err()
	}

	return nil
}

func (r *redisRepository) GetJSON(key string, target interface{}) error {
	res, err := r.client.Do(r.ctx, "JSON.GET", key).Text()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(res), target)
}

func (r *redisRepository) UpdateJSONField(key string, jsonPath string, value interface{}) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Do(r.ctx, "JSON.SET", key, jsonPath, valueBytes).Err()
}

func (r *redisRepository) DeleteJSON(key string) error {
	return r.client.Do(r.ctx, "JSON.DEL", key).Err()
}

func (r *redisRepository) Enqueue(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.RPush(r.ctx, key, data).Err()
}

func (r *redisRepository) Dequeue(key string) ([]byte, error) {
	return r.client.LPop(r.ctx, key).Bytes()
}

func (r *redisRepository) Backqueue(key string, value interface{}) error {
	return r.client.LPush(r.ctx, key, value).Err()
}

func (r *redisRepository) BackQueueAtomic(key string, value string) error {
    script := `return redis.call('LPUSH', KEYS[1], ARGV[1])`
    _, err := r.client.Eval(context.Background(), script, []string{key}, value).Result()
    return err
}
