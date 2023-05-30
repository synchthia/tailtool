package player

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func fetch() *cobra.Command {
	c := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch user profile",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			systera.NewClient()
			defer systera.Shutdown()

			res, err := systera.FetchPlayerProfileByName(name)
			if err != nil {
				logrus.WithError(err).Errorln("Failed fetch player profile")
				os.Exit(1)
				return
			}
			pp.Println(res)
		},
	}
	c.Flags().StringP("name", "n", "", "name")
	c.MarkFlagRequired("name")

	return c
}
