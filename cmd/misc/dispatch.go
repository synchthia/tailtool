package misc

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func dispatch() *cobra.Command {
	c := &cobra.Command{
		Use:   "dispatch",
		Short: "Send command in all or specified servers.",
		Run: func(cmd *cobra.Command, args []string) {
			server, _ := cmd.Flags().GetString("server")
			command, _ := cmd.Flags().GetString("command")

			systera.NewClient()
			defer systera.Shutdown()

			if err := systera.Dispatch(server, command); err != nil {
				logrus.WithError(err).Errorln("Failed dispatch command")
				os.Exit(1)
			} else {
				logrus.WithFields(logrus.Fields{
					"server":  server,
					"command": command,
				}).Infof("Command dispatched")
			}
		},
	}

	c.Flags().StringP("server", "s", "global", "server")

	c.Flags().StringP("command", "c", "", "command")
	c.MarkFlagRequired("command")

	return c
}
