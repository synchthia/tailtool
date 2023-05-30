package proxy

import (
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/cmd/proxy/motd"
)

var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Manage proxy",
}

func New() *cobra.Command {
	rootCmd.AddCommand(motd.New())
	//rootCmd.AddCommand(favicon())

	return rootCmd
}
