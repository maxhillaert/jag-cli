package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jag",
	Short: "A brief description of your application",
	Long:  `A longer description of your application`,
}

func Execute() error {
	return rootCmd.Execute()
}
