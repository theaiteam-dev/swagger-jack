package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminUsersListCmd_workspaceId string
	adminUsersListCmd_userId string
)

var adminUsersListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminUsersListCmd_workspaceId
		queryParams["userId"] = adminUsersListCmd_userId
		resp, err := c.Do("GET", "/api/admin/users/subscriptions", pathParams, queryParams, nil)
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
	adminUsersCmd.AddCommand(adminUsersListCmd)
	adminUsersListCmd.Flags().StringVar(&adminUsersListCmd_workspaceId, "workspaceId", "", "")
	adminUsersListCmd.Flags().StringVar(&adminUsersListCmd_userId, "userId", "", "")
	adminUsersListCmd.MarkFlagRequired("workspaceId")
	adminUsersListCmd.MarkFlagRequired("userId")
}
