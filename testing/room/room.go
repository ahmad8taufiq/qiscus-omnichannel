package main

import (
	"fmt"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
)

func main() {
	// repo := repository.NewAuthRepository(fmt.Sprintf("%s/api/v1/auth", config.AppConfig.QiscusBaseURL))
	svc := service.NewRoomService(repository.NewRoomRepository())

	authResp, err := svc.GetRoomById("313532428", "Tyd09B0kkpxTQRRxKspq1744880265", "rvcbl-fcsngqk40iyo7ks_admin@qismo.com")
	if err != nil {
		logger.Logger.WithError(err).Error("Login failed")
		return
	}

	fmt.Println("✅ Auth Token:", authResp.Results.Room.LastCommentID)
}