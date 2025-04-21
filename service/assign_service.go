package service

import (
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
)

type AssignService interface {
	Assign(roomID string, agentID int) (*models.AssignAgentResponse, error)
}

type assignService struct {
	repo repository.AssignRepository
}

func NewAssignService(repo repository.AssignRepository) AssignService {
	return &assignService{repo: repo}
}

func (s *assignService) Assign(roomID string, agentID int) (*models.AssignAgentResponse, error) {
	return s.repo.AssignAgent(roomID, agentID)
}
