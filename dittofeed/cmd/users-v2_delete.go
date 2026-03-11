package cmd

import (
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	usersV2DeleteCmd_workspaceId string
)

var usersV2DeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = usersV2DeleteCmd_workspaceId
		usersV2DeleteCmd_userIds_vals, _ := cmd.Flags().GetStringArray("userIds")
		queryParams["userIds"] = strings.Join(usersV2DeleteCmd_userIds_vals, ",")
		resp, err := c.Do("DELETE", "/api/users/v2", pathParams, queryParams, nil)
		if err != nil {
			return err
		}
		jsonMode, _ := cmd.Root().PersistentFlags().GetBool("json")
		noColor, _ := cmd.Root().PersistentFlags().GetBool("no-color")
		if jsonMode {
			fmt.Printf("%s\n", string(resp))
		} else {
			if err := output.PrintTable(resp, noColor); err != nil {
				fmt.Println(string(resp))
			}
		}
		return nil
	},
}

func init() {
	usersV2Cmd.AddCommand(usersV2DeleteCmd)
	usersV2DeleteCmd.Flags().StringVar(&usersV2DeleteCmd_workspaceId, "workspaceId", "", "")
	usersV2DeleteCmd.Flags().StringArray("userIds", nil, "")
	usersV2DeleteCmd.MarkFlagRequired("workspaceId")
	usersV2DeleteCmd.MarkFlagRequired("userIds")
}
