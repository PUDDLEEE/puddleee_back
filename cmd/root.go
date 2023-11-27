package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "puddlee",
	Short: "Puddlee is a Chat Service for developer.",
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
func Execute() error {
	return rootCmd.Execute()
}
