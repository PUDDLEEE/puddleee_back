package cmd

import (
	"github.com/PUDDLEEE/puddleee_back/internal/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Puddlee Backend Server",
	Long: `Run Puddlee Backend Server on local.
	configuration file lies on the config directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.InitApp()
	},
}
