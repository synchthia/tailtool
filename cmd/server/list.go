package server

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/nebula"
)

func list() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "List servers",
		Run: func(cmd *cobra.Command, args []string) {
			nebula.NewClient()
			defer nebula.Shutdown()
			res, err := nebula.ListServers()
			if err != nil {
				logrus.WithError(err).Errorln("Failed lookup servers")
			}

			frmt := "%-10v %-20v %-10v %-10v %-10v %-10v %-5v\n"
			fmt.Printf(frmt, "Name", "DisplayName", "Address", "Port", "Fallback", "Online", "Players")

			for _, v := range res {
				fmt.Printf(frmt, v.Name, v.DisplayName, v.Address, v.Port, v.Fallback, v.Status.Online, fmt.Sprintf("%d/%d",v.Status.Players.Online, v.Status.Players.Max))
			}

		},
	}

	return c
}
