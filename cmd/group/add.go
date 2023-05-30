package group

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func add() *cobra.Command {
	c := &cobra.Command{
		Use:   "add",
		Short: "add group",
		Run: func(cmd *cobra.Command, args []string) {
            name, _ := cmd.Flags().GetString("name")
            prefix, _ := cmd.Flags().GetString("prefix")

			systera.NewClient()
			defer systera.Shutdown()

			err := systera.CreateGroup(name, prefix)
			if err != nil {
				logrus.WithError(err).Errorln("Failed create group")
				os.Exit(1)
				return
			}
            fmt.Printf("Perms: %s [%s]", name, prefix)
		},
	}
	c.Flags().StringP("name", "n", "", "name")
	c.MarkFlagRequired("name")

	c.Flags().StringP("prefix", "p", "global", "prefix")
	c.MarkFlagRequired("prefix")

	return c
}
