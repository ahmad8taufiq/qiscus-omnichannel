package cmd

import (
	"qiscus-omnichannel/app"

	"github.com/spf13/cobra"
)

var dequeueCmd = &cobra.Command{
	Use:   "dequeue",
	Short: "Start Redis dequeue listener",
	Run: func(cmd *cobra.Command, args []string) {
		app.StartDequeueListener()
	},
}

func init() {
	rootCmd.AddCommand(dequeueCmd)
}
