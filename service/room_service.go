package service

import (
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
)

type RoomService interface {
	GetRoomById(roomID, sdkToken, sdkUserID string) (*models.SdkRoomResponse, error)
}

type roomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{repo: repo}
}

func (s *roomService) GetRoomById(roomID, sdkToken, sdkUserID string) (*models.SdkRoomResponse, error) {
	return s.repo.GetRoomById(roomID, sdkToken, sdkUserID)
}
