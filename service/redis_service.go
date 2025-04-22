package service

import (
	"time"

	"qiscus-omnichannel/repository"
)

type RedisService interface {
	SetCache(key string, value string, ttl time.Duration) error
	GetCache(key string) (string, error)
	SetJSON(key string, value interface{}, ttl time.Duration) error
	GetJSON(key string, target interface{}) error
	UpdateJSONField(key string, jsonPath string, value interface{}) error
	DeleteJSON(key string) error
	Enqueue(key string, value interface{}) error
	Dequeue(key string) ([]byte, error)
	Backqueue(key string, value interface{}) error
}

type redisService struct {
	repo repository.RedisRepository
}

func NewRedisService(repo repository.RedisRepository) RedisService {
	return &redisService{repo: repo}
}

func (s *redisService) SetCache(key string, value string, ttl time.Duration) error {
	return s.repo.SetCache(key, value, ttl)
}

func (s *redisService) GetCache(key string) (string, error) {
	return s.repo.GetCache(key)
}

func (s *redisService) SetJSON(key string, value interface{}, ttl time.Duration) error {
	return s.repo.SetJSON(key, value, ttl)
}

func (s *redisService) GetJSON(key string, target interface{}) error {
	return s.repo.GetJSON(key, target)
}

func (s *redisService) UpdateJSONField(key string, jsonPath string, value interface{}) error {
	return s.repo.UpdateJSONField(key, jsonPath, value)
}

func (s *redisService) DeleteJSON(key string) error {
	return s.repo.DeleteJSON(key)
}

func (s *redisService) Enqueue(key string, value interface{}) error {
	return s.repo.Enqueue(key, value)
}

func (s *redisService) Dequeue(key string) ([]byte, error) {
	return s.repo.Dequeue(key)
}

func (s *redisService) Backqueue(key string, value interface{}) error {
	return s.repo.Backqueue(key, value)
}
