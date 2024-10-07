package player

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the server command
var rootCmd = &cobra.Command{
	Use:   "player",
	Short: "Manage players",
}

func New() *cobra.Command {
    //rootCmd.AddCommand(cmdInit())
    rootCmd.AddCommand(fetch())
    rootCmd.AddCommand(setGroup())
	return rootCmd
}
