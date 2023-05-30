package server

import (
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/nebula"
)

func remove() *cobra.Command {
	c := &cobra.Command{
		Use:   "remove",
		Short: "Remove server",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			nebula.NewClient()
			defer nebula.Shutdown()
			if err := nebula.RemoveServerEntry(name); err != nil {
				panic(err)
			}
		},
	}

	c.Flags().StringP("name", "n", "", "name")
	c.MarkFlagRequired("name")

	return c
}
