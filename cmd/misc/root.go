package misc

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "misc",
	Short: "Misc commands",
}

func New() *cobra.Command {
	rootCmd.AddCommand(announce())
	rootCmd.AddCommand(dispatch())
    rootCmd.AddCommand(iplookup())
	return rootCmd
}
