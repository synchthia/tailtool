package player

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func setGroup() *cobra.Command {
	c := &cobra.Command{
		Use:   "set-group",
		Short: "Assign group",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			groups, _ := cmd.Flags().GetStringArray("groups")

			systera.NewClient()
			defer systera.Shutdown()

			fetchRes, err := systera.FetchPlayerProfileByName(name)
			if err != nil {
				logrus.WithError(err).Errorln("Failed fetch player profile")
				os.Exit(1)
				return
			}

			if err := systera.SetGroup(fetchRes.Uuid, groups); err != nil {
				logrus.WithError(err).Errorln("Failed set player group")
				os.Exit(1)
				return
			}
		},
	}

	c.Flags().StringP("name", "n", "", "name")
	c.Flags().StringArrayP("groups", "g", nil, "groups")

	c.MarkFlagRequired("name")
	c.MarkFlagRequired("groups")

	return c
}
