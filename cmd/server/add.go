package server

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/nebula"
)

func add() *cobra.Command {
	c := &cobra.Command{
		Use:   "add",
		Short: "Add server",
		Run: func(cmd *cobra.Command, args []string) {
			// Add
			name, _ := cmd.Flags().GetString("name")
			displayName, _ := cmd.Flags().GetString("display-name")
			address, _ := cmd.Flags().GetString("address")
			port, _ := cmd.Flags().GetInt("port")
			fallback, _ := cmd.Flags().GetBool("fallback")

			if len(displayName) == 0 {
				displayName = name
			}

			nebula.NewClient()
			defer nebula.Shutdown()

			if err := nebula.AddServerEntry(name, displayName, address, int32(port), fallback); err != nil {
				logrus.WithError(err).Fatal("Failed add server")
			}
		},
	}

	c.Flags().StringP("name", "n", "", "name")
	c.MarkFlagRequired("name")

	c.Flags().String("display-name", "", "display-name")

	c.Flags().StringP("address", "a", "", "address")
	c.MarkFlagRequired("address")

	c.Flags().IntP("port", "p", 33400, "port")
	c.Flags().BoolP("fallback", "f", false, "fallback")

	return c
}
