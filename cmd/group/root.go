package group

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the server command
var rootCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage groups",
}

func New() *cobra.Command {
    rootCmd.AddCommand(add())
    rootCmd.AddCommand(list())
	return rootCmd
}
