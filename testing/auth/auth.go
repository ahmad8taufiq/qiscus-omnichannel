package main

import (
	"fmt"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
)

func main() {
	repo := repository.NewAuthRepository(config.AppConfig.QiscusAuthURL)
	svc := service.NewAuthService(repo)

	authResp, err := svc.Login(config.AppConfig.QiscusEmail, config.AppConfig.QiscusPassword)
	if err != nil {
		logger.Logger.WithError(err).Error("Login failed")
		return
	}

	fmt.Println("✅ Auth Token:", authResp.Data.User.AuthenticationToken)
}