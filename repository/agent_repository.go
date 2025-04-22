package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
	"strings"
)

type AgentRepository interface {
	AssignAgent(roomID string, agentID int) (*models.AssignAgentResponse, error)
	MarkAsResolved(roomID, notes, lastCommentID string) (*models.MarkAsResolvedResponse, error)
}

type agentRepo struct{}

func NewAgentRepository() AgentRepository {
	return &agentRepo{}
}

func (r *agentRepo) AssignAgent(roomID string, agentID int) (*models.AssignAgentResponse, error) {
	form := url.Values{}
	form.Add("room_id", roomID)
	form.Add("agent_id", fmt.Sprintf("%d", agentID))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/admin/service/assign_agent", config.AppConfig.QiscusBaseURL), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Qiscus-App-Id", config.AppConfig.QiscusAppID)
	req.Header.Add("Qiscus-Secret-Key", config.AppConfig.QiscusSecretKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)
	var response models.AssignAgentResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *agentRepo) MarkAsResolved(roomID, notes, lastCommentID string) (*models.MarkAsResolvedResponse, error) {
	form := url.Values{}
	form.Add("room_id", roomID)
	form.Add("notes", notes)
	form.Add("last_comment_id", lastCommentID)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/admin/service/mark_as_resolved", config.AppConfig.QiscusBaseURL), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Qiscus-App-Id", config.AppConfig.QiscusAppID)
	req.Header.Add("Authorization", config.AppConfig.QiscusSecretKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response models.MarkAsResolvedResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
