package cmd

import (
	"fmt"
	"net/http"
	handler "qiscus-omnichannel/services/webhook"
	logger "qiscus-omnichannel/tools/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var port int
var log = logrus.New()

func init() {
	log.SetFormatter(&logger.ColorFormatter{})
	log.SetLevel(logrus.DebugLevel)
}

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Start webhook listener",
	Run: func(cmd *cobra.Command, args []string) {
		log := logrus.New()
		log.SetFormatter(&logger.ColorFormatter{})

		http.HandleFunc("/", handler.WebhookHandler(log))

		addr := fmt.Sprintf(":%d", port)
		log.Infof("🚀 Listening on port %d", port)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	webhookCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the webhook server on")
	rootCmd.AddCommand(webhookCmd)
}
