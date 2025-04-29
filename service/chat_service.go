package service

import (
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
)

type ChatService interface {
	InitiateChat(req *models.InitiateChatRequest) (*models.InitiateChatResponse, error)
}

type chatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) ChatService {
	return &chatService{repo: repo}
}

func (s *chatService) InitiateChat(req *models.InitiateChatRequest) (*models.InitiateChatResponse, error) {
	return s.repo.InitiateChat(req)
}
