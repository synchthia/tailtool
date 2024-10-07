package group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synchthia/tailtool/systera"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func list() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "list group",
		Run: func(cmd *cobra.Command, args []string) {
			systera.NewClient()
			defer systera.Shutdown()

			res, err := systera.GetGroups()
			if err != nil {
				logrus.WithError(err).Errorln("Failed fetch group profile")
				os.Exit(1)
				return
			}

			b := new(bytes.Buffer)
			e := json.NewEncoder(b)
			e.SetEscapeHTML(false)
			e.Encode(res)

			r2, err := PrettyString(b.String())
			fmt.Println(r2)
		},
	}

	return c
}
