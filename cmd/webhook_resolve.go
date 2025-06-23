package cmd

import (
	"fmt"
	"net/http"
	"qiscus-omnichannel/app"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/console"
	"qiscus-omnichannel/tools/logger"

	"github.com/spf13/cobra"
)

var webhookResolvePort int

var webhookResolveCmd = &cobra.Command{
	Use:   "webhook-resolve",
	Short: "Start webhook resolve listener",
	Run:   runWebhookResolveServer,
}

func init() {
	webhookResolveCmd.Flags().IntVarP(&webhookResolvePort, "port", "p", 8082, "Port to run the webhook resolve server on")
	rootCmd.AddCommand(webhookResolveCmd)
}

func runWebhookResolveServer(cmd *cobra.Command, args []string) {
	log := logger.Logger
	redisService := service.RedisService(repository.NewRedisRepository())

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.WebhookResolveHandler(redisService))

	addr := fmt.Sprintf(":%d", webhookResolvePort)
	console.ConsoleGreet("Webhook Resolve", "1.0.0", "", webhookResolvePort)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Failed to start webhook resolve server: %v", err)
	}
}
