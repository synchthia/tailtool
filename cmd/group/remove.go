package group

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func remove() *cobra.Command {
	c := &cobra.Command{
		Use:   "remove",
		Short: "remove group",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			systera.NewClient()
			defer systera.Shutdown()

			err := systera.RemoveGroup(name)
			if err != nil {
				logrus.WithError(err).Errorln("Failed remove group")
				os.Exit(1)
				return
			}
		},
	}
	c.Flags().StringP("name", "n", "", "name")
	c.MarkFlagRequired("name")

	return c
}
