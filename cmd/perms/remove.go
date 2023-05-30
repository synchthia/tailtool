package perms

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func remove() *cobra.Command {
	c := &cobra.Command{
		Use:   "remove",
		Short: "remove perms",
		Run: func(cmd *cobra.Command, args []string) {
            name, _ := cmd.Flags().GetString("name")
            target, _ := cmd.Flags().GetString("target")
            perms, _ := cmd.Flags().GetStringArray("perms")

			systera.NewClient()
			defer systera.Shutdown()

			err := systera.RemovePermission(name, target, perms)
			if err != nil {
				logrus.WithError(err).Errorln("Failed fetch perms profile")
				os.Exit(1)
				return
			}
            fmt.Printf("Perms: %s [%s] %v", name, target, perms)
		},
	}
	c.Flags().StringP("name", "n", "", "name")
	c.MarkFlagRequired("name")

	c.Flags().StringP("target", "t", "global", "target")

	c.Flags().StringArray("perms", []string{""}, "perms")
	c.MarkFlagRequired("perms")

	return c
}
