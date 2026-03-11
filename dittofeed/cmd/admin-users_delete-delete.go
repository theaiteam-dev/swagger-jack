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
	adminUsersDeleteDeleteCmd_workspaceId string
)

var adminUsersDeleteDeleteCmd = &cobra.Command{
	Use: "delete-delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminUsersDeleteDeleteCmd_workspaceId
		adminUsersDeleteDeleteCmd_userIds_vals, _ := cmd.Flags().GetStringArray("userIds")
		queryParams["userIds"] = strings.Join(adminUsersDeleteDeleteCmd_userIds_vals, ",")
		resp, err := c.Do("DELETE", "/api/admin/users/v2", pathParams, queryParams, nil)
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
	adminUsersCmd.AddCommand(adminUsersDeleteDeleteCmd)
	adminUsersDeleteDeleteCmd.Flags().StringVar(&adminUsersDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	adminUsersDeleteDeleteCmd.Flags().StringArray("userIds", nil, "")
	adminUsersDeleteDeleteCmd.MarkFlagRequired("workspaceId")
	adminUsersDeleteDeleteCmd.MarkFlagRequired("userIds")
}
