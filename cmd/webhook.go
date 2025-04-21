package cmd

import (
	"fmt"
	"net/http"
	"qiscus-omnichannel/app"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"

	"github.com/spf13/cobra"
)

var port int

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Start webhook listener",
	Run:   runWebhookServer,
}

func init() {
	webhookCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the webhook server on")
	rootCmd.AddCommand(webhookCmd)
}

func runWebhookServer(cmd *cobra.Command, args []string) {
	log := logger.Logger
	svc := service.NewWebhookService()

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.WebhookHandler(log, svc))

	addr := fmt.Sprintf(":%d", port)
	log.Infof("🚀 Webhook is running on port %d", port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
