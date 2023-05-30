package group

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func list() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "list group",
		Run: func(cmd *cobra.Command, args []string) {
            server, _ := cmd.Flags().GetString("server")

			systera.NewClient()
			defer systera.Shutdown()

			res, err := systera.GetGroups(server)
			if err != nil {
				logrus.WithError(err).Errorln("Failed fetch group profile")
				os.Exit(1)
				return
			}
			pp.Println(res.Groups)
		},
	}

	c.Flags().StringP("server", "s", "global", "server")

	return c
}
