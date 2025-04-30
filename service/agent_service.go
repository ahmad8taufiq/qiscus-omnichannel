package service

import (
	"fmt"
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
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

// func (s *agentService) GetAvailableAgents(adminToken, roomID string) (*models.AvailableAgentsResponse, error) {
// 	return s.repo.GetAvailableAgents(adminToken, roomID)
// }

func (s *agentService) GetAvailableAgents(adminToken, roomID string) (*models.AvailableAgentsResponse, error) {
	redisRepo := repository.NewRedisRepository()
	redisService := RedisService(redisRepo)

	availableAgent, _ := s.repo.GetAvailableAgents(adminToken, roomID)

	for _, agent := range availableAgent.Data.Agents {
		_, err := redisService.GetCache(fmt.Sprintf("%d", agent.ID))
		if err != nil {
			redisService.SetCache(fmt.Sprintf("%d", agent.ID), fmt.Sprintf("%d", agent.CurrentCustomerCount), time.Minute * 10)
		}
	}

	return availableAgent, nil
}