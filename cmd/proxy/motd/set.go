package motd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/nebula"
	"github.com/synchthia/tailtool/utils"
)

func set() *cobra.Command {
	c := &cobra.Command{
		Use:   "set",
		Short: "Set motd",
		Run: func(cmd *cobra.Command, args []string) {
			motd, _ := cmd.Flags().GetString("motd")

			nebula.NewClient()
			defer nebula.Shutdown()
			if err := nebula.SetMotd(motd); err != nil {
				logrus.WithError(err).Errorln("Failed update motd")
				os.Exit(1)
			}

			coloredMotd := utils.ColorText(strings.ReplaceAll(motd, "\\n", "\n"))
			logrus.Infof("Motd updated to:")
			fmt.Println(coloredMotd)
		},
	}

	c.Flags().StringP("motd", "m", "", "motd")
	c.MarkFlagRequired("motd")

	return c
}
