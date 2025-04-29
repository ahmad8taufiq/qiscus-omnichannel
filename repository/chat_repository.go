package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
)

type ChatRepository interface {
	InitiateChat(req *models.InitiateChatRequest) (*models.InitiateChatResponse, error)
}

type chatRepo struct{}

func NewChatRepository() ChatRepository {
	return &chatRepo{}
}

func (r *chatRepo) InitiateChat(reqData *models.InitiateChatRequest) (*models.InitiateChatResponse, error) {
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.AppConfig.QiscusBaseURL+"/api/v2/qiscus/initiate_chat", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response models.InitiateChatResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
