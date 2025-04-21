package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"qiscus-omnichannel/models"
)

type AuthRepository interface {
	Authenticate(email, password string) (*models.AuthResponse, error)
}

type authRepository struct {
	apiURL string
}

func NewAuthRepository(apiURL string) AuthRepository {
	return &authRepository{apiURL: apiURL}
}

func (r *authRepository) Authenticate(email, password string) (*models.AuthResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("email", email)
	writer.WriteField("password", password)
	writer.Close()

	req, err := http.NewRequest("POST", r.apiURL, body)
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