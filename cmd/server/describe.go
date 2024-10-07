package server

import (
	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/nebula"
)

func describe() *cobra.Command {
	c := &cobra.Command{
		Use:   "describe",
		Short: "Describe servers",
		Run: func(cmd *cobra.Command, args []string) {
			nebula.NewClient()
			defer nebula.Shutdown()
			res, err := nebula.ListServers()
			if err != nil {
				logrus.WithError(err).Errorln("Failed lookup servers")
			}

            pp.Println(res)
		},
	}

	return c
}
