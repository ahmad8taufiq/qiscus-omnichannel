package cmd

import (
	"qiscus-omnichannel/app"

	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Start Redis resolve assigned listener",
	Run: func(cmd *cobra.Command, args []string) {
		app.StartResolveListener()
	},
}

func init() {
	rootCmd.AddCommand(resolveCmd)
}
