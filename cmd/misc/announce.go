package misc

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
	"github.com/synchthia/tailtool/utils"
)

func announce() *cobra.Command {
	c := &cobra.Command{
		Use:   "announce",
		Short: "Announce message in all or specified servers.",
		Run: func(cmd *cobra.Command, args []string) {
			server, _ := cmd.Flags().GetString("server")
			message, _ := cmd.Flags().GetString("message")

			systera.NewClient()
			defer systera.Shutdown()

			if err := systera.Announce(server, message); err != nil {
				logrus.WithError(err).Errorln("Failed dispatch announce")
				os.Exit(1)
			} else {
				logrus.WithFields(logrus.Fields{
					"server":  server,
					"message": utils.ColorText(message),
				}).Infof("Message Announced")
			}
		},
	}

	c.Flags().StringP("server", "s", "global", "server")

	c.Flags().StringP("message", "m", "", "message")
	c.MarkFlagRequired("message")

	return c
}
