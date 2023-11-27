package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Puddlee is a Chat Service for developer.",
}

func Execute() error {
	return rootCmd.Execute()
}
