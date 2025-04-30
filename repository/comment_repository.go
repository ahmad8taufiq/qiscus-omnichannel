package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
)

type CommentRepository interface {
	PostComment(token, userId string, req *models.PostCommentRequest) (*models.PostCommentResponse, error)
}

type commentRepo struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepo{}
}

func (r *commentRepo) PostComment(token, userId string, reqData *models.PostCommentRequest) (*models.PostCommentResponse, error) {
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.AppConfig.QiscusApi21URL+"/api/v2/sdk/post_comment", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Qiscus-Sdk-App-Id", config.AppConfig.QiscusAppID)
	req.Header.Set("Qiscus-Sdk-Token", token)
	req.Header.Set("Qiscus-Sdk-User-Id", userId)
	req.Header.Set("Qiscus-Sdk-Version", "WEB_2.10.9")
	req.Header.Set("Origin", config.AppConfig.QiscusBaseURL)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response models.PostCommentResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
