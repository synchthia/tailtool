package misc

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/nebula"
)

func iplookup() *cobra.Command {
	c := &cobra.Command{
		Use:   "iplookup",
		Short: "Check incoming ip via nebula-api",
		Run: func(cmd *cobra.Command, args []string) {
			ip, _ := cmd.Flags().GetString("ip")

			nebula.NewClient()
			defer nebula.Shutdown()

			if res, err := nebula.IPLookup(ip); err != nil {
				logrus.WithError(err).Errorln("Failed check ip")
				os.Exit(1)
			} else {
				pp.Println(res)
			}
		},
	}

	c.Flags().StringP("ip", "i", "", "ip")
	c.MarkFlagRequired("ip")

	return c
}
