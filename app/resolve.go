package app

import (
	"encoding/json"
	"strconv"
	"time"

	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
	tools "qiscus-omnichannel/tools/parser"
)

func StartResolveListener() {
	log := logger.Logger
	log.Info("🚀 Resolve listener started")

	redisService := service.RedisService(repository.NewRedisRepository())
	agentService := service.AgentService(repository.NewAgentRepository())
	roomService := service.RoomService(repository.NewRoomRepository())

	for {
		payload, err := redisService.Dequeue("assigned")
		if err != nil {
			log.WithError(err).Error("❌ Failed to dequeue assigned message")
			time.Sleep(1 * time.Second)
			continue
		}

		if len(payload) > 0 {
			log.Infof("📥 Dequeued assigned message: %s", string(payload))
			
			assignedMessage, _ := tools.Parser[models.Message](payload)
			log.Infof("👤 Assigned message: %v", assignedMessage)

			_, sdkEmail, sdkToken, err := GetCredentials(redisService)
			if err != nil {
				log.WithError(err).Error("❌ MarkAsResolved failed to get credentials")
				// err := redisService.Backqueue("assigned", payload)
				err := redisService.BackQueueAtomic("assigned", string(payload))
				if err != nil {
					log.WithError(err).Error("❌ Failed to backqueue message due to resolve GetCredentials")
				}
				time.Sleep(3 * time.Second)
				continue
			}

			room, err := roomService.GetRoomById(assignedMessage.RoomId, sdkToken, sdkEmail)
			if err != nil {
				log.WithError(err).Error("❌ MarkAsResolved failed to get room")
				// err := redisService.Backqueue("assigned", payload)
				err := redisService.BackQueueAtomic("assigned", string(payload))
				if err != nil {
					log.WithError(err).Error("❌ Failed to backqueue message due to resolve GetRoomById")
				}
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("✅ LastCommentID: %v", room.Results.Room.LastCommentID)
			lastCommentId := strconv.FormatInt(room.Results.Room.LastCommentID, 10)

			resolved, err := agentService.MarkAsResolved(assignedMessage.RoomId, "Resolved", lastCommentId)
			if err != nil {
				log.WithError(err).Error("❌ Failed to resolved assigned")
				err := redisService.BackQueueAtomic("assigned", string(payload))
				if err != nil {
					log.WithError(err).Error("❌ Failed to mark as resolved message due to unknown error")
				}
				time.Sleep(3 * time.Second)
				continue
			}

			var agents []models.Agents
			cached, err := redisService.GetCache("agents")
			if err == nil && cached != "" {
				err = json.Unmarshal([]byte(cached), &agents)
				if err != nil {
					logger.Logger.Info("⚠️ Failed to unmarshal cached agents, starting fresh")
					agents = []models.Agents{}
				}
			}

			agentID, _ := strconv.Atoi(assignedMessage.AgentID)
			for i, agent := range agents {
				if agent.ID == agentID {
					agents[i].CurrentCustomerCount -= 1
					break
				}
			}

			updatedJSON, _ := json.Marshal(agents)
			redisService.SetCache("agents", string(updatedJSON), 10*time.Minute)
			log.Infof("✅ Resolved: %v", resolved.Data.Service.IsResolved)
		} else {
			log.Debug("⏳ No message in assigned, waiting...")
		}
		
	}
}
