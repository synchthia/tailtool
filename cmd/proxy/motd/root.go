package motd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "Manage motd",
}

func New() *cobra.Command {
	rootCmd.AddCommand(set())
	return rootCmd
}
