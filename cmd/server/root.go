package server

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the server command
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage servers",
}

func New() *cobra.Command {
	rootCmd.AddCommand(add())
	rootCmd.AddCommand(remove())
	rootCmd.AddCommand(list())
	rootCmd.AddCommand(describe())
	return rootCmd
}
