package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
)

type AuthRepository interface {
	GetNonce() (*models.NonceResponse, error)
	Authenticate(email, password string) (*models.AuthResponse, error)
}

type authRepository struct {
	apiURL string
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) Authenticate(email, password string) (*models.AuthResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("email", email)
	writer.WriteField("password", password)
	writer.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/auth", config.AppConfig.QiscusBaseURL), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Error: %s", respBody)
	}

	var authResp models.AuthResponse
	if err := json.Unmarshal(respBody, &authResp); err != nil {
		return nil, err
	}

	return &authResp, nil
}

func (r *authRepository) GetNonce() (*models.NonceResponse, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/sdk/auth/nonce", config.AppConfig.QiscusApiURL), bytes.NewBuffer(nil))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Qiscus_sdk_app_id", "rvcbl-fcsngqk40iyo7ks")
	req.Header.Set("Origin", "https://omnichannel.qiscus.com")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get nonce")
	}

	var nonceResp models.NonceResponse
	if err := json.NewDecoder(res.Body).Decode(&nonceResp); err != nil {
		return nil, err
	}

	return &nonceResp, nil
}