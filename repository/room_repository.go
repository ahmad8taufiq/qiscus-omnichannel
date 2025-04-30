package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
)

type RoomRepository interface {
	GetRoomById(roomID, sdkToken, sdkUserID string) (*models.SdkRoomResponse, error)
}

type roomRepo struct{}

func NewRoomRepository() RoomRepository {
	return &roomRepo{}
}

func (r *roomRepo) GetRoomById(roomID, sdkToken, sdkUserID string) (*models.SdkRoomResponse, error) {
	url := fmt.Sprintf("%s/api/v2/sdk/get_room_by_id?id=%s", config.AppConfig.QiscusApiURL, roomID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Qiscus-Sdk-App-Id", config.AppConfig.QiscusAppID)
	req.Header.Set("Qiscus-Sdk-Token", sdkToken)

	if sdkUserID != "" {
		req.Header.Set("Qiscus-Sdk-User-Id", sdkUserID)
	}

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

	var response models.SdkRoomResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}