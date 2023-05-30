package perms

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the server command
var rootCmd = &cobra.Command{
	Use:   "perms",
	Short: "Manage permss",
}

func New() *cobra.Command {
    rootCmd.AddCommand(add())
    rootCmd.AddCommand(remove())
	return rootCmd
}
