package main

import (
	"qiscus-omnichannel/cmd"
	"qiscus-omnichannel/tools/logger"
	"qiscus-omnichannel/tools/redis"
)

func main() {
	logger.InitLogger()
	redis.InitRedis()

	if err := cmd.Execute(); err != nil {
		logger.Logger.WithError(err).Fatal("❌ Failed to execute command")
	}
}
