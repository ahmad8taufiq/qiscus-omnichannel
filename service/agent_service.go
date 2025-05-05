package service

import (
	"encoding/json"
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/tools/logger"
	"time"
)

type AgentService interface {
	AssignAgent(roomID string, agentID int) (*models.AssignAgentResponse, error)
	MarkAsResolved(roomID, notes, lastCommentID string) (*models.MarkAsResolvedResponse, error)
	GetAvailableAgents(adminToken, roomID string) (*models.AvailableAgentsResponse, error)
}

type agentService struct {
	repo repository.AgentRepository
}

func NewAssignService(repo repository.AgentRepository) AgentService {
	return &agentService{repo: repo}
}

func (s *agentService) AssignAgent(roomID string, agentID int) (*models.AssignAgentResponse, error) {
	return s.repo.AssignAgent(roomID, agentID)
}

func (s *agentService) MarkAsResolved(roomID, notes, lastCommentID string) (*models.MarkAsResolvedResponse, error) {
	return s.repo.MarkAsResolved(roomID, notes, lastCommentID)
}

func (s *agentService) GetAvailableAgents(adminToken, roomID string) (*models.AvailableAgentsResponse, error) {
	redisRepo := repository.NewRedisRepository()
	redisService := RedisService(redisRepo)

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

	availableAgent, _ := s.repo.GetAvailableAgents(adminToken, roomID)
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

	return availableAgent, nil
}