package group

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/synchthia/systera-api/systerapb"
	"github.com/synchthia/tailtool/systera"
)

func apply() *cobra.Command {
	c := &cobra.Command{
		Use:   "apply",
		Short: "apply group",
		Run: func(cmd *cobra.Command, args []string) {
			f, _ := cmd.Flags().GetString("file")

			systera.NewClient()
			defer systera.Shutdown()

			raw, err := os.ReadFile(f)
			if err != nil {
				panic(err)
			}

			var res *systerapb.FetchGroupsResponse

			json.Unmarshal(raw, &res)

			// Get
			getRes, getErr := systera.GetGroups()
			if getErr != nil {
				panic(getErr)
			}

			// 作成
			for _, group := range res.Groups {
				if GetGroup(getRes.Groups, group.GroupName) == nil {
					fmt.Println("Create: " + group.GroupName)
					if err := systera.CreateGroup(group); err != nil {
						panic(err)
					}

				}
			}

			for _, group := range res.Groups {
				// データベースに存在する場合
				if GetGroup(getRes.Groups, group.GroupName) != nil {
					// 定義ファイルに存在する場合は作成
					if GetGroup(res.Groups, group.GroupName) != nil {
						// データベースに存在するグループなら更新
						fmt.Println("Update: " + group.GroupName)
                        pp.Println(group)
						if err := systera.UpdateGroup(group); err != nil {
							panic(err)
						}

					} else {
						fmt.Println("Remove: " + group.GroupName)
						if err := systera.RemoveGroup(group.GroupName); err != nil {
							panic(err)
						}
					}
				}
			}
		},
	}
	c.Flags().StringP("file", "f", "", "file")
	c.MarkFlagRequired("file")

	return c
}

func GetGroup(groups []*systerapb.GroupEntry, name string) *systerapb.GroupEntry {
	for _, group := range groups {
		if group.GroupName == name {
			return group
		}
	}
	return nil
}
