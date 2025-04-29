package main

import (
	"fmt"
	"qiscus-omnichannel/app"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
)

func main() {
	redisRepo := repository.NewRedisRepository()
	redisSvc := service.RedisService(redisRepo)

	nonceResp, err := app.GetNonce(redisSvc)
	if err != nil {
		logger.Logger.WithError(err).Error("Login failed")
		return
	}

	fmt.Println("✅ Nonce:", nonceResp)
}