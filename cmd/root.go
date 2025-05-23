package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qiscus-omnichannel",
	Short: "CLI for Qiscus Omnichannel Integration",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(webhookCmd)
	rootCmd.AddCommand(serverCmd)
}