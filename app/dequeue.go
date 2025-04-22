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
	agentService := service.AgentService(repository.NewAgentRepository())

	for {
		payload, err := redisService.Dequeue("new_session_queue")
		if err != nil {
			log.WithError(err).Error("❌ Failed to dequeue message")
			time.Sleep(1 * time.Second)
			continue
		}

		if len(payload) > 0 {
			log.Infof("📥 Dequeued message: %s", string(payload))
			
			newMessage, _ := tools.Parser[models.Message](payload)

			adminToken, _ := GetAdminToken(redisService)
			availableAgent, err := agentService.GetAvailableAgents(adminToken, newMessage.RoomId)
			if err != nil {
				log.WithError(err).Error("❌ Failed to get available agents")
				continue
			}
			
			log.Infof("👤 Available agents: %v", availableAgent)

			assigned := false
			for _, agent := range availableAgent.Data.Agents {
				if agent.CurrentCustomerCount < 2 {
					log.Infof("🧑 Assigning agent %s (ID: %d)", agent.Name, agent.ID)
			
					assignResp, err := agentService.AssignAgent(newMessage.RoomId, agent.ID)
					if err != nil {
						log.WithError(err).Error("❌ Failed to assign agent")
						continue
					}
			
					assigned = true
					err = redisService.Enqueue("assigned", payload)
					if err != nil {
						log.WithError(err).Error("❌ Failed to assigned queue to Redis")
					} else {
						log.Info("📤 Data assigned to Redis successfully")
					}
			
					log.Infof("✅ Agent assigned successfully: %+v", assignResp.Data)
					break 
				}
			}

			if !assigned {
				log.Warn("⚠️ No available agent (all full), requeueing message...")
				err := redisService.Backqueue("new_session_queue", payload)
				if err != nil {
					log.WithError(err).Error("❌ Failed to requeue message")
				} else {
					time.Sleep(5 * time.Second)
				}
			}
		} else {
			log.Debug("⏳ No message in queue, waiting...")
		}
		
	}
}
