package player

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func cmdInit() *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "Fetch user profile",
		Run: func(cmd *cobra.Command, args []string) {
			systera.NewClient()
			defer systera.Shutdown()

			entry, err := systera.InitPlayerProfile()
			if err != nil {
				logrus.WithError(err).Errorln("Failed init player profile")
				os.Exit(1)
				return
			}

            pp.Println(entry)
		},
	}
	return c
}
