package service

import (
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
)

type AgentService interface {
	Assign(roomID string, agentID int) (*models.AssignAgentResponse, error)
}

type agentService struct {
	repo repository.AgentRepository
}

func NewAssignService(repo repository.AgentRepository) AgentService {
	return &agentService{repo: repo}
}

func (s *agentService) Assign(roomID string, agentID int) (*models.AssignAgentResponse, error) {
	return s.repo.AssignAgent(roomID, agentID)
}
