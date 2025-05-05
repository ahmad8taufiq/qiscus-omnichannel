package app

import (
	"encoding/json"
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

	agentRepo := repository.NewAgentRepository()
	redisRepo := repository.NewRedisRepository()

	for {
		payload, err := redisService.Dequeue("new_session_queue")
		if err != nil {
			log.Debug("⏳ No message in queue, waiting...")
			time.Sleep(1 * time.Second)
			continue
		}

		log.Infof("📥 Dequeued message: %s", string(payload))

		newMessage, err := tools.Parser[models.Message](payload)
		if err != nil {
			log.WithError(err).Error("❌ Failed to parse message payload")
			continue
		}

		adminToken, _, sdkToken, err := GetCredentials(redisService)
		if err != nil {
			log.WithError(err).Error("❌ Dequeue failed to get credentials")
			err := redisService.BackQueueAtomic("new_session_queue", string(payload))
			if err != nil {
				log.WithError(err).Error("❌ Failed to backqueue message due to dequeue error")
			}
			time.Sleep(3 * time.Second)
			continue
		}

		log.Infof("👤 SDK Token: %s", sdkToken)

		var agents []models.Agents
		cached, err := redisService.GetCache("agents")
		if err == nil && cached != "" {
			err = json.Unmarshal([]byte(cached), &agents)
			if err != nil {
				logger.Logger.Info("⚠️ Failed to unmarshal cached agents, starting fresh")
				agents = []models.Agents{}
			}
		} else {
			agents = []models.Agents{}
		}

		availableAgent, _ := agentRepo.GetAvailableAgents(adminToken, newMessage.RoomId)

		for _, agent := range availableAgent.Data.Agents {
			exists := false
			for _, a := range agents {
				if a.ID == agent.ID {
					exists = true
					break
				}
			}
			if !exists {
				agents = append(agents, models.Agents{
					ID:                    agent.ID,
					CurrentCustomerCount:  agent.CurrentCustomerCount,
				})
			}
		}

		jsonBytes, _ := json.Marshal(agents)
		jsonString := string(jsonBytes)

		redisRepo.SetCache("agents", jsonString, 10*time.Minute)

		cached, err = redisService.GetCache("agents")
		if err == nil && cached != "" {
			err = json.Unmarshal([]byte(cached), &agents)
			if err != nil {
				logger.Logger.Info("⚠️ Failed to unmarshal cached agents, starting fresh")
				agents = []models.Agents{}
			}
		} else {
			agents = []models.Agents{}
		}

		if len(agents) == 0 {
			log.Warn("⚠️ No available agent, requeueing message...")
			err := redisService.BackQueueAtomic("new_session_queue", string(payload))
			if err != nil {
				log.WithError(err).Error("❌ Failed to requeue message")
			} else {
				time.Sleep(5 * time.Second)
			}
		} else {
			minAgent := agents[0]
			for _, agent := range agents[1:] {
				if agent.CurrentCustomerCount < minAgent.CurrentCustomerCount {
					minAgent = agent
				}
			}
	
			if minAgent.CurrentCustomerCount < 2 {
				assignResp, err := agentService.AssignAgent(newMessage.RoomId, minAgent.ID)
				if err != nil {
					log.WithError(err).Error("❌ Failed to assign agent")
					continue
				}

				for i, agent := range agents {
					if agent.ID == minAgent.ID {
						agents[i].CurrentCustomerCount += 1
						break
					}
				}
				updatedJSON, err := json.Marshal(agents)
				if err != nil {
					log.Fatalf("Failed to marshal updated agents: %v", err)
				}

				err = redisService.SetCache("agents", string(updatedJSON), 10*time.Minute)
				if err != nil {
					log.Fatalf("SetCache failed: %v", err)
				}

				err = redisService.Enqueue("assigned", newMessage)
				if err != nil {
					log.WithError(err).Error("❌ Failed to enqueue to assigned queue")
				} else {
					log.Info("📤 Data assigned to Redis successfully")
				}
				
				log.Infof("✅ Agent assigned successfully: %+v", assignResp.Data)
			}
		}
	}
}
