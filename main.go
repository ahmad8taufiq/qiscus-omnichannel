package main

import (
	"qiscus-omnichannel/cmd"
	"qiscus-omnichannel/tools/logger"
)

func main() {
	logger.InitLogger()

	cmd.Execute()
}
