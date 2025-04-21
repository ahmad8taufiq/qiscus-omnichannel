package main

import (
	"qiscus-omnichannel/cmd"
	"qiscus-omnichannel/tools/logger"
)

func main() {
	logger.InitLogger()

	if err := cmd.Execute(); err != nil {
		logger.Logger.WithError(err).Fatal("❌ Failed to execute command")
	}
}
