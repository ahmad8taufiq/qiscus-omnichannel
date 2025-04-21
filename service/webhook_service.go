package service

import (
	"encoding/json"
	"fmt"

	"qiscus-omnichannel/models"
)

type WebhookService interface {
	ProcessWebhook(body []byte) (*models.Message, error)
}

type webhookService struct{}

func NewWebhookService() WebhookService {
	return &webhookService{}
}

func (s *webhookService) ProcessWebhook(body []byte) (*models.Message, error) {
	var data models.Message
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return &data, nil
}
