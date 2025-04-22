package app

import (
	"time"

	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
	tools "qiscus-omnichannel/tools/parser"
)

func StartDequeueListener() {
	log := logger.Logger
	log.Info("🚀 Dequeue listener started")

	redisService := service.RedisService(repository.NewRedisRepository())

	for {
		payload, err := redisService.Dequeue("new_session_queue")
		if err != nil {
			log.WithError(err).Error("❌ Failed to dequeue message")
			time.Sleep(1 * time.Second)
			continue
		}

		if len(payload) > 0 {
			log.Infof("📥 Dequeued message: %s", string(payload))
			
			data, _ := tools.Parser[models.Message](payload)
			log.Infof("🚀 Processing message: %s", data.RoomId)
		} else {
			log.Debug("⏳ No message in queue, waiting...")
		}
		
	}
}
