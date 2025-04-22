package service

import (
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
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
	return s.repo.GetAvailableAgents(adminToken, roomID)
}